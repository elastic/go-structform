package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	structform "github.com/urso/go-structform"
)

type parser struct {
	visitor structform.Visitor

	// last fail state
	err error

	// parser state machine
	states       []state // state stack for nested arrays/objects
	currentState state

	// preallocate stack memory for up to 32 nested arrays/objects
	statesBuf [32]state

	literalBuffer  []byte
	literalBuffer0 [64]byte
	isDouble       bool
	inEscape       bool
	required       int
}

var (
	errFailing             = errors.New("JSON parser failed")
	errIncomplete          = errors.New("Incomplete JSON input")
	errUnknownChar         = errors.New("unknown character")
	errQuoteMissing        = errors.New("missing closing quote")
	errExpectColon         = errors.New("expected ':' after map key")
	errUnexpectedDictClose = errors.New("unexpected '}'")
	errUnexpectedArrClose  = errors.New("unexpected ']'")
	errExpectedDigit       = errors.New("expected a digit")
	errExpectedObject      = errors.New("expected JSON object")
	errExpectedArray       = errors.New("expected JSON array")
	errExpectedFieldName   = errors.New("expected JSON object field name")
	errExpectedInteger     = errors.New("expected integer value")
	errExpectedNull        = errors.New("expected null value")
	errExpectedFalse       = errors.New("expected false value")
	errExpectedTrue        = errors.New("expected true value")
	errExpectedArrayField  = errors.New("expected ']' or ','")
)

var (
	nullSymbol  = []byte("null")
	trueSymbol  = []byte("true")
	falseSymbol = []byte("false")
)

var whitespace = " \t\r\n"

type state uint8

const (
	failedState state = iota
	startState

	arrState
	arrStateValue
	arrStateNext

	dictState
	dictFieldState
	dictNextFieldState
	dictFieldValue
	dictFieldValueSep
	dictFieldStateEnd

	nullState
	trueState
	falseState
	stringState
	numberState
)

func ParseReader(in io.Reader, vs structform.Visitor) (int64, error) {
	p := newParser(vs)
	i, err := io.Copy(p, in)
	if err == nil {
		err = p.finalize()
	}
	return i, err
}

func Parse(b []byte, vs structform.Visitor) error {
	p := newParser(vs)
	p.err = p.feed(b)
	if p.err == nil {
		p.err = p.finalize()
	}
	return p.err
}

func ParseString(str string, vs structform.Visitor) error {
	return Parse([]byte(str), vs)
}

func newParser(vs structform.Visitor) *parser {
	p := &parser{
		visitor:      vs,
		currentState: startState,
	}
	p.states = p.statesBuf[:0]
	p.literalBuffer = p.literalBuffer0[:0]
	return p
}

func (p *parser) Write(b []byte) (int, error) {
	p.err = p.feed(b)
	if p.err != nil {
		return 0, p.err
	}
	return len(b), nil
}

func (p *parser) feed(b []byte) error {
	for len(b) > 0 {
		var err error

		switch p.currentState {
		case failedState:
			return p.err
		case startState:
			b, err = p.stepStart(b)

		case dictState:
			b, err = p.stepDict(b, true)

		case dictNextFieldState:
			b, err = p.stepDict(b, false)

		case dictFieldState:
			b, err = p.stepDictKey(b)

		case dictFieldValueSep:
			if b = bytes.TrimLeft(b, whitespace); len(b) > 0 {
				if b[0] != ':' {
					err = errExpectColon
				}
				b = b[1:]
				p.currentState = dictFieldValue
			}

		case dictFieldValue:
			b, err = p.stepValue(b, dictFieldStateEnd)

		case dictFieldStateEnd:
			b, err = p.stepDictValueEnd(b)

		case arrState:
			b, err = p.stepArray(b, true)

		case arrStateValue:
			b, err = p.stepValue(b, arrStateNext)

		case arrStateNext:
			b, err = p.stepArrValueEnd(b)

		case nullState:
			b, err = p.stepNULL(b)

		case trueState:
			b, err = p.stepTRUE(b)

		case falseState:
			b, err = p.stepFALSE(b)

		case stringState:
			b, err = p.stepString(b)

		case numberState:
			b, err = p.stepNumber(b)

		default:
			return errFailing
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) finalize() error {
	if p.currentState == numberState {
		err := p.reportNumber(p.literalBuffer, p.isDouble)
		if err != nil {
			return err
		}
		p.popState()
	}

	if len(p.states) > 0 && p.currentState != startState {
		return errIncomplete
	}

	return nil
}

func (p *parser) pushState(next state) {
	if p.currentState != failedState {
		p.states = append(p.states, p.currentState)
	}
	p.currentState = next
}

func (p *parser) popState() {
	if len(p.states) == 0 {
		p.currentState = failedState
	} else {
		last := len(p.states) - 1
		p.currentState = p.states[last]
		p.states = p.states[:last]
	}
}

func (p *parser) stepStart(b []byte) ([]byte, error) {
	return p.stepValue(b, p.currentState)
}

func (p *parser) stepValue(b []byte, retState state) ([]byte, error) {
	b = bytes.TrimLeft(b, whitespace)
	if len(b) == 0 {
		return b, nil
	}

	p.currentState = retState
	c := b[0]
	switch c {
	case '{': // start dictionary
		p.pushState(dictState)
		return b[1:], p.visitor.OnObjectStart(-1, structform.AnyType)

	case '[': // start array
		p.pushState(arrState)
		return b[1:], p.visitor.OnArrayStart(-1, structform.AnyType)

	case 'n': // parse "null"
		p.pushState(nullState)
		p.required = 3
		return p.stepNULL(b[1:])

	case 'f': // parse "false"
		p.pushState(falseState)
		p.required = 4
		return p.stepFALSE(b[1:])

	case 't': // parse "true"
		p.pushState(trueState)
		p.required = 3
		return p.stepTRUE(b[1:])

	case '"': // parse string
		p.literalBuffer = p.literalBuffer[:0]
		p.pushState(stringState)
		p.inEscape = false
		return p.stepString(b[:])

	default:
		// parse number?
		isNumber := c == '-' || c == '+' || c == '.' || isDigit(c)
		if !isNumber {
			return b, errUnknownChar
		}

		p.literalBuffer = p.literalBuffer0[:0]
		p.pushState(numberState)
		p.isDouble = false
		return p.stepNumber(b)
	}
}

func (p *parser) stepDict(b []byte, allowEnd bool) ([]byte, error) {
	b = bytes.TrimLeft(b, whitespace)
	if len(b) == 0 {
		return b, nil
	}

	c := b[0]
	switch c {
	case '}':
		if !allowEnd {
			return nil, errUnexpectedDictClose
		}
		return p.endDict(b)

	case '"':
		p.currentState = dictFieldState
		return b, nil

	default:
		return nil, errExpectedFieldName
	}
}

func (p *parser) stepDictKey(b []byte) ([]byte, error) {
	str, done, b, err := p.doString(b)
	if done && err == nil {
		p.currentState = dictFieldValueSep
		err = p.visitor.OnKey(str)
	}
	return b, err
}

func (p *parser) stepDictValueEnd(b []byte) ([]byte, error) {
	b = bytes.TrimLeft(b, whitespace)
	if len(b) == 0 {
		return b, nil
	}

	c := b[0]
	switch c {
	case '}':
		return p.endDict(b)
	case ',':
		p.currentState = dictNextFieldState
		return b[1:], p.visitor.OnFieldNext()
	default:
		return nil, errUnknownChar
	}
}

func (p *parser) endDict(b []byte) ([]byte, error) {
	p.popState()
	return b[1:], p.visitor.OnObjectFinished()
}

func (p *parser) stepArray(b []byte, allowEnd bool) ([]byte, error) {
	b = bytes.TrimLeft(b, whitespace)
	if len(b) == 0 {
		return b, nil
	}

	c := b[0]
	switch c {
	case ']':
		if !allowEnd {
			return nil, errUnexpectedArrClose
		}
		return p.endArray(b)
	}

	p.currentState = arrStateValue
	return b, nil
}

func (p *parser) stepArrValueEnd(b []byte) ([]byte, error) {
	b = bytes.TrimLeft(b, whitespace)
	if len(b) == 0 {
		return b, nil
	}

	c := b[0]
	switch c {
	case ']':
		return p.endArray(b)
	case ',':
		p.currentState = arrStateValue
		return b[1:], p.visitor.OnElemNext()
	default:
		return nil, errUnknownChar
	}
}

func (p *parser) endArray(b []byte) ([]byte, error) {
	p.popState()
	return b[1:], p.visitor.OnArrayFinished()
}

func (p *parser) stepString(b []byte) ([]byte, error) {
	str, done, b, err := p.doString(b)
	if done && err == nil {
		p.popState()
		err = p.visitor.OnString(str)
	}
	return b, err
}

func (p *parser) doString(b []byte) (string, bool, []byte, error) {
	stop := -1
	done := false

	delta := 1
	buf := b
	atStart := len(p.literalBuffer) == 0
	if atStart {
		delta = 2
		buf = b[1:]
	}

	for i, c := range buf {
		if p.inEscape {
			p.inEscape = false
			continue
		}

		if c == '"' {
			done = true
			stop = i + delta
			break
		}
		if c == '\\' {
			p.inEscape = true
		}
	}

	if !done {
		p.literalBuffer = append(p.literalBuffer, b...)
		return "", false, nil, nil
	}

	rest := b[stop:]
	b = b[:stop]
	if len(p.literalBuffer) > 0 {
		b = append(p.literalBuffer, b...)
		p.literalBuffer = b[:0] // reset buffer
	}

	// use encoding/json to unescape and parse into go string
	var str string
	err := json.Unmarshal(b, &str)
	return str, done, rest, err
}

func (p *parser) stepNumber(b []byte) ([]byte, error) {
	// search for char in stop-set
	stop := -1
	done := false
	for i, c := range b {
		isStopChar := c == ' ' || c == '\t' || c == '\f' || c == '\n' || c == '\r' ||
			c == ',' ||
			c == ']' ||
			c == '}'
		if isStopChar {
			stop = i
			done = true
			break
		}

		p.isDouble = p.isDouble || c == '.' || c == 'e' || c == 'E'
	}

	if !done {
		p.literalBuffer = append(p.literalBuffer, b...)
		return nil, nil
	}

	rest := b[stop:]
	b = b[:stop]
	if len(p.literalBuffer) > 0 {
		b = append(p.literalBuffer, b...)
		p.literalBuffer = b[:0] // reset buffer
	}

	err := p.reportNumber(b, p.isDouble)
	p.popState()
	return rest, err
}

func (p *parser) reportNumber(b []byte, isDouble bool) error {
	// parse number
	var err error
	if isDouble {
		var f float64
		if f, err = strconv.ParseFloat(string(b), 64); err == nil {
			err = p.visitor.OnFloat64(f)
		}
	} else {
		var i int64
		if i, err = strconv.ParseInt(string(b), 10, 64); err == nil {
			err = p.visitor.OnInt64(i)
		}
	}

	return err
}

func (p *parser) stepNULL(b []byte) ([]byte, error) {
	b, done, err := p.stepKind(b, []byte("null"), errExpectedNull)
	if done {
		err = p.visitor.OnNil()
	}
	return b, err
}

func (p *parser) stepTRUE(b []byte) ([]byte, error) {
	b, done, err := p.stepKind(b, []byte("true"), errExpectedTrue)
	if done {
		err = p.visitor.OnBool(true)
	}
	return b, err
}

func (p *parser) stepFALSE(b []byte) ([]byte, error) {
	b, done, err := p.stepKind(b, []byte("false"), errExpectedFalse)
	if done {
		err = p.visitor.OnBool(false)
	}
	return b, err
}

func (p *parser) stepKind(b []byte, kind []byte, err error) ([]byte, bool, error) {
	n := p.required
	s := kind[len(kind)-n:]
	done := true
	if L := len(b); L < n {
		done = false
		p.required = n - L
		n = L
		s = s[:L]
	}

	if !bytes.HasPrefix(b, s) {
		return b, false, err
	}

	if done {
		p.popState()
	}
	return b[n:], done, nil
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
