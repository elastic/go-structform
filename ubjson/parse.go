package ubjson

import (
	"encoding/binary"
	"errors"
	"io"
	"math"

	structform "github.com/urso/go-structform"
)

type Parser struct {
	visitor structform.Visitor

	// last fail state
	err error

	// parser state machine
	state      stateStack
	valueState stateStack

	length lengthStack

	buffer  []byte
	buffer0 [64]byte

	// internal parser state
	marker    byte
	valueType structform.BaseType
}

type stateStack struct {
	stack   []state // state stack for nested arrays/objects
	stack0  [32]state
	current state
}

type lengthStack struct {
	stack   []int64
	stack0  [32]int64
	current int64
}

type stateType uint8
type stateStep uint8

type state struct {
	stateType
	stateStep
}

const (
	stFail stateType = iota
	stNext
	stFixed       // values of fixed size
	stHighPrec    // high precision number
	stString      // string
	stArray       // array
	stArrayDyn    // dynamic array
	stArrayCount  // array with element count
	stArrayTyped  // typed array with element count
	stObject      // object
	stObjectDyn   // dynamic object
	stObjectCount // object with known # of fields
	stObjectTyped // object with all values of same type
)

const (
	stStart stateStep = iota

	// stValue sub-states
	stNil
	stNoop
	stTrue
	stFalse
	stInt8
	stUInt8
	stInt16
	stInt32
	stInt64
	stFloat32
	stFloat64
	stChar

	// variable size primitive value types
	stWithLen

	// array/object states
	stWithType0
	stWithType1
	stCont
	stFieldName
	stFieldNameLen
	stNextField
)

var (
	errUnknownMarker = errors.New("unknown ubjson marker")
	errIncomplete    = errors.New("Incomplete UBJSON input")
	errNegativeLen   = errors.New("negative length encountered")
	errInvalidState  = errors.New("invalid state")
	errMissingArrEnd = errors.New("missing ']'")
	errMissingObjEnd = errors.New("missing '}'")
	errMissingCount  = errors.New("missing count marker")
)

func ParseReader(in io.Reader, vs structform.Visitor) (int64, error) {
	return NewParser(vs).ParseReader(in)
}

func Parse(b []byte, vs structform.Visitor) error {
	return NewParser(vs).Parse(b)
}

func ParseString(str string, vs structform.Visitor) error {
	return NewParser(vs).ParseString(str)
}

func NewParser(vs structform.Visitor) *Parser {
	p := &Parser{
		visitor: vs,
	}
	p.buffer = p.buffer0[:0]
	p.length.stack = p.length.stack0[:0]
	p.state.current = state{stNext, stStart}
	p.state.stack = p.state.stack0[:0]
	p.valueState.stack = p.valueState.stack0[:0]
	return p
}

func (p *Parser) Parse(b []byte) error {
	p.err = p.feed(b)
	if p.err != nil {
		p.err = p.finalize()
	}
	return p.err
}

func (p *Parser) ParseReader(in io.Reader) (int64, error) {
	n, err := io.Copy(p, in)
	if err == nil {
		err = p.finalize()
	}
	return n, err
}

func (p *Parser) ParseString(s string) error {
	return p.Parse(str2Bytes(s))
}

func (p *Parser) finalize() error {
	st := &p.state.current
	incomplete := len(p.state.stack) > 0 ||
		st.stateStep != stStart ||
		st.stateType != stNext

	if incomplete {
		return errIncomplete
	}
	return nil
}

func (p *Parser) Write(b []byte) (int, error) {
	p.err = p.feed(b)
	if p.err != nil {
		p.state.current = state{stFail, stStart}
		return 0, p.err
	}
	return len(b), nil
}

func (p *Parser) feed(b []byte) error {
	for len(b) > 0 {
		var err error

		switch p.state.current.stateType {
		case stFail:
			return p.err
		case stNext:
			b, err = p.stepStart(b)
		case stFixed:
			b, err = p.stepFixedValue(b)
		case stHighPrec:
			b, err = p.stepString(b)
		case stString:
			b, err = p.stepString(b)

		case stArray:
			b, err = p.stepArrayInit(b)
		case stArrayDyn:
			b, err = p.stepArrayDyn(b)
		case stArrayCount:
			b, err = p.stepArrayCount(b)
		case stArrayTyped:
			b, err = p.stepArrayTyped(b)

		case stObject:
			b, err = p.stepObjectInit(b)
		case stObjectDyn:
			b, err = p.stepObjectDyn(b)
		case stObjectCount:
			b, err = p.stepObjectCount(b)
		case stObjectTyped:
			b, err = p.stepObjectTyped(b)

		default:
			err = errInvalidState
		}
		if err != nil {
			p.err = err
			return err
		}
	}

	return nil
}

func (p *Parser) stepFixedValue(b []byte) ([]byte, error) {
	var (
		tmp []byte
		err error
	)

	switch p.state.current.stateStep {
	case stNil:
		p.popState()
		err = p.visitor.OnNil()
	case stNoop:
		p.popState()
	case stTrue:
		p.popState()
		err = p.visitor.OnBool(true)
	case stFalse:
		p.popState()
		err = p.visitor.OnBool(false)
	case stInt8:
		p.popState()
		b, err = b[1:], p.visitor.OnInt8(int8(b[0]))
	case stUInt8:
		p.popState()
		b, err = b[1:], p.visitor.OnUint8(b[0])
	case stChar:
		b, tmp = p.collect(b, 1)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnByte(tmp[0])
		}
	case stInt16:
		b, tmp = p.collect(b, 2)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnInt16(readInt16(tmp))
		}
	case stInt32:
		b, tmp = p.collect(b, 4)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnInt32(readInt32(tmp))
		}
	case stInt64:
		b, tmp = p.collect(b, 8)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnInt64(readInt64(tmp))
		}
	case stFloat32:
		b, tmp = p.collect(b, 4)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnFloat32(readFloat32(tmp))
		}
	case stFloat64:
		b, tmp = p.collect(b, 8)
		if tmp != nil {
			p.popState()
			err = p.visitor.OnFloat64(readFloat64(tmp))
		}
	}

	return b, err
}

func (p *Parser) stepString(b []byte) ([]byte, error) {
	var (
		err error
		st  = &p.state.current
	)

	switch st.stateStep {
	case stStart:
		b, err = p.stepLen(b, st.withStep(stWithLen))
	case stWithLen:
		L := p.length.current
		var tmp []byte
		if b, tmp = p.collect(b, int(L)); tmp != nil {
			p.popLenState()
			err = p.visitor.OnString(bytes2Str(tmp))
		}
	}

	return b, err
}

func (p *Parser) stepArrayInit(b []byte) ([]byte, error) {
	var (
		err error
		st  = &p.state.current
	)

	switch b[0] {
	case countMarker:
		b, st.stateType = b[1:], stArrayCount
	case typeMarker:
		b, st.stateType = b[1:], stArrayTyped
	default:
		st.stateType = stArrayDyn
		err = p.visitor.OnArrayStart(-1, structform.AnyType)
	}

	return b, err
}

func (p *Parser) stepArrayDyn(b []byte) ([]byte, error) {
	if b[0] == arrEndMarker {
		p.popState()
		return b[1:], p.visitor.OnArrayFinished()
	}

	if st := &p.state.current; st.stateStep == stStart {
		st.stateStep = stCont // ensure continuation state is pushed to stack
		return p.stepStart(b)
	}
	return p.stepStart(b)
}

func (p *Parser) stepArrayCount(b []byte) ([]byte, error) {
	var (
		st   = &p.state.current
		step = st.stateStep
	)

	// parse array header
	switch step {
	case stStart:
		return p.stepLen(b, st.withStep(stWithLen))
	case stWithLen:
		err := p.visitor.OnArrayStart(int(p.length.current), structform.AnyType)
		if err != nil {
			return b, err
		}
	}

	end, b, err := p.stepArrayCountedContent(b)
	if end || err != nil {
		p.popState()
		return b, err
	}
	return p.stepStart(b) // read next value
}

func (p *Parser) stepArrayTyped(b []byte) ([]byte, error) {
	step := p.state.current.stateStep

	// parse typed array header
	switch step {
	case stStart, stWithType0, stWithType1:
		return p.stepTypeLenHeader(b, stWithLen)

	case stWithLen:
		err := p.visitor.OnArrayStart(int(p.length.current), p.valueType)
		if err != nil {
			return b, err
		}
	}

	end, b, err := p.stepArrayCountedContent(b)
	if end {
		p.popState()
		p.valueState.pop()
	} else if err == nil {
		p.pushState(p.valueState.current)
	}

	return b, err
}

func (p *Parser) stepArrayCountedContent(b []byte) (bool, []byte, error) {
	end, b, err := p.checkArrEnd(b)
	if err != nil {
		return end, b, err
	}

	if end {
		p.length.pop()
		return end, b, p.visitor.OnArrayFinished()
	}

	if st := &p.state.current; st.stateStep == stWithLen {
		st.stateStep = stCont
	}

	p.length.current--
	return end, b, err
}

func (p *Parser) stepTypeLenHeader(b []byte, cont stateStep) ([]byte, error) {
	st := p.state.current
	step := st.stateStep

	switch step {
	case stStart:
		return p.stepType(b, st.withStep(stWithType0))

	case stWithType0:
		if b[0] != countMarker {
			return b, errMissingCount
		}
		p.state.current = st.withStep(stWithType1)
		return b[1:], nil

	case stWithType1:
		return p.stepLen(b, st.withStep(cont))

	default:
		return b, nil
	}
}

func (p *Parser) stepObjectInit(b []byte) ([]byte, error) {
	var (
		st  = &p.state.current
		err error
	)

	switch b[0] {
	case countMarker:
		b, st.stateType = b[1:], stObjectCount
	case typeMarker:
		b, st.stateType = b[1:], stObjectTyped
	default:
		st.stateType, err = stObjectDyn, p.visitor.OnObjectStart(-1, structform.AnyType)
	}

	return b, err
}

func (p *Parser) stepObjectDyn(b []byte) ([]byte, error) {
	var (
		err  error
		st   = &p.state.current
		step = st.stateStep
	)

	if step == stStart || step == stNextField {
		if b[0] == objEndMarker {
			p.popState()
			return b[1:], p.visitor.OnObjectFinished()
		}
	}

	switch step {
	case stStart:
		b, err = p.stepLen(b, st.withStep(stFieldNameLen))
	case stFieldNameLen:
		L := p.length.current
		var tmp []byte
		if b, tmp = p.collect(b, int(L)); tmp != nil {
			p.popLen()
			err = p.visitor.OnKey(bytes2Str(tmp))
		}
		st.stateStep = stCont
	case stCont:
		st.stateStep = stStart
		b, err = p.stepStart(b)
	}

	return b, err
}

func (p *Parser) stepObjectCount(b []byte) ([]byte, error) {
	st := &p.state.current
	step := st.stateStep

	if step == stStart {
		return p.stepLen(b, st.withStep(stWithLen))
	}

	end, b, err := p.stepObjectCountedContent(b, false)
	if end {
		p.popLenState()
	}
	return b, err
}

func (p *Parser) stepObjectTyped(b []byte) ([]byte, error) {
	st := &p.state.current
	step := st.stateStep

	switch step {
	case stStart, stWithType0, stWithType1:
		return p.stepTypeLenHeader(b, stWithLen)
	}

	end, b, err := p.stepObjectCountedContent(b, true)
	if end {
		p.popLenState()
		p.valueState.pop()
	}
	return b, err
}

func (p *Parser) stepObjectCountedContent(b []byte, typed bool) (bool, []byte, error) {
	var (
		err  error
		st   = &p.state.current
		step = st.stateStep
		end  = false
	)

	switch step {
	case stWithLen:
		L := p.length.current
		err := p.visitor.OnObjectStart(int(L), structform.AnyType)
		if err != nil {
			return end, b, err
		}

		if L == 0 {
			end, b, err = p.checkObjEnd(b)
			break
		}

		st.stateStep = stFieldName
		fallthrough

	case stFieldName:
		b, err = p.stepLen(b, st.withStep(stFieldNameLen))

	case stFieldNameLen:
		L := p.length.current
		var tmp []byte
		if b, tmp = p.collect(b, int(L)); tmp != nil {
			p.popLen()
			err = p.visitor.OnKey(bytes2Str(tmp))
		}
		st.stateStep = stCont

	case stCont:
		st.stateStep = stNextField

		// handle object field value
		if typed {
			p.pushState(p.valueState.current)
		} else {
			b, err = p.stepStart(b)
		}

	case stNextField:
		p.length.current--
		end, b, err = p.checkObjEnd(b)
		if err == nil && !end {
			st.stateStep = stFieldName
			b, err = p.stepLen(b, st.withStep(stFieldNameLen))
		}
	}

	if end {
		err = p.visitor.OnObjectFinished()
	}
	return end, b, err
}

func (p *Parser) checkArrEnd(b []byte) (bool, []byte, error) {
	return p.checkTEnd(b, arrEndMarker, errMissingArrEnd)
}

func (p *Parser) checkObjEnd(b []byte) (bool, []byte, error) {
	return p.checkTEnd(b, objEndMarker, errMissingObjEnd)
}

func (p *Parser) checkTEnd(b []byte, marker byte, err error) (bool, []byte, error) {
	if p.length.current > 0 {
		return false, b, nil
	}
	if b[0] != marker {
		return false, nil, err
	}
	return true, b[1:], nil
}

func (p *Parser) stepType(b []byte, cont state) ([]byte, error) {
	marker := b[0]
	b = b[1:]
	p.state.current = cont

	// TODO: analyze marker
	state, err := markerToStartState(marker)
	if err != nil {
		return nil, err
	}
	p.valueState.push(state)
	p.valueType = markerToBaseType(marker)

	return b, nil
}

func (p *Parser) stepLen(b []byte, cont state) ([]byte, error) {
	if p.marker == noMarker {
		p.marker = b[0]
		b = b[1:]
		if len(b) == 0 {
			return nil, nil
		}
	}

	var tmp []byte
	complete := false
	L := int64(-1)

	switch p.marker {
	case int8Marker:
		complete, L, b = true, int64(int8(b[0])), b[1:]
	case uint8Marker:
		complete, L, b = true, int64(b[0]), b[1:]
	case int16Marker:
		if b, tmp = p.collect(b, 2); tmp != nil {
			complete, L = true, int64(readInt16(tmp))
		}
	case int32Marker:
		if b, tmp = p.collect(b, 4); tmp != nil {
			complete, L = true, int64(readInt32(tmp))
		}
	case int64Marker:
		if b, tmp = p.collect(b, 8); tmp != nil {
			complete, L = true, readInt64(tmp)
		}
	}

	if !complete {
		return b, nil
	}

	if L < 0 {
		return nil, errNegativeLen
	}

	p.marker = noMarker
	p.state.current = cont
	p.pushLen(L)
	return b, nil
}

func (p *Parser) collect(b []byte, count int) ([]byte, []byte) {
	if len(p.buffer) > 0 {
		delta := len(p.buffer) - count
		if delta > 0 {
			N := delta
			complete := true
			if N > len(b) {
				complete = false
				N = len(b)
			}

			p.buffer = append(p.buffer, b[:N]...)
			if !complete {
				return nil, nil
			}
		}

		if len(p.buffer) >= count {
			tmp := p.buffer[:count]
			if len(p.buffer) == count {
				p.buffer = p.buffer0[:0]
			} else {
				p.buffer = p.buffer[count:]
			}
			return b, tmp
		}
	}

	if len(b) >= count {
		return b[count:], b[:count]
	}

	p.buffer = append(p.buffer, b...)
	return nil, nil
}

func (p *Parser) stepStart(b []byte) ([]byte, error) {
	state, err := markerToStartState(b[0])
	if err != nil {
		return nil, err
	}

	switch state.stateStep {
	case stNil:
		return b[1:], p.visitor.OnNil()
	case stNoop:
		return b[1:], nil
	case stTrue:
		return b[1:], p.visitor.OnBool(true)
	case stFalse:
		return b[1:], p.visitor.OnBool(false)
	default:
		return p.advanceMarker(state, b)
	}
}

func (p *Parser) advanceMarker(s state, b []byte) ([]byte, error) {
	p.pushState(s)
	return b[1:], nil
}

func (p *Parser) pushState(next state) { p.state.push(next) }
func (p *Parser) popState()            { p.state.pop() }
func (p *Parser) pushLen(l int64)      { p.length.push(l) }
func (p *Parser) popLen()              { p.length.pop() }

func (p *Parser) popLenState() {
	p.popState()
	p.popLen()
}

func (s *lengthStack) push(l int64) {
	s.stack = append(s.stack, s.current)
	s.current = l
}

func (s *lengthStack) pop() {
	if len(s.stack) == 0 {
		s.current = -1
	} else {
		last := len(s.stack) - 1
		s.current = s.stack[last]
		s.stack = s.stack[:last]
	}
}

func (s *stateStack) push(next state) {
	if s.current.stateType != stFail {
		s.stack = append(s.stack, s.current)
	}
	s.current = next
}

func (s *stateStack) pop() {
	if len(s.stack) == 0 {
		s.current = state{stFail, stStart}
	} else {
		last := len(s.stack) - 1
		s.current = s.stack[last]
		s.stack = s.stack[:last]
	}
}

func readInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(b))
}

func readInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

func readInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

func readFloat32(b []byte) float32 {
	bits := binary.BigEndian.Uint32(b)
	return math.Float32frombits(bits)
}

func readFloat64(b []byte) float64 {
	bits := binary.BigEndian.Uint64(b)
	return math.Float64frombits(bits)
}

func markerToStartState(marker byte) (state, error) {
	switch marker {
	case nullMarker:
		return state{stFixed, stNil}, nil
	case noopMarker:
		return state{stFixed, stNoop}, nil
	case trueMarker:
		return state{stFixed, stTrue}, nil
	case falseMarker:
		return state{stFixed, stFalse}, nil
	case int8Marker:
		return state{stFixed, stInt8}, nil
	case uint8Marker:
		return state{stFixed, stUInt8}, nil
	case int16Marker:
		return state{stFixed, stInt16}, nil
	case int32Marker:
		return state{stFixed, stInt32}, nil
	case int64Marker:
		return state{stFixed, stInt64}, nil
	case float32Marker:
		return state{stFixed, stFloat32}, nil
	case float64Marker:
		return state{stFixed, stFloat64}, nil
	case highPrecMarker:
		return state{stHighPrec, stStart}, nil
	case charMarker:
		return state{stFixed, stChar}, nil
	case stringMarker:
		return state{stString, stStart}, nil
	case objStartMarker:
		return state{stObject, stStart}, nil
	case arrStartMarker:
		return state{stArray, stStart}, nil
	default:
		return state{stFail, stStart}, errUnknownMarker
	}
}

func markerToBaseType(marker byte) structform.BaseType {
	switch marker {
	case falseMarker, trueMarker:
		return structform.BoolType
	case charMarker:
		return structform.ByteType
	case int8Marker:
		return structform.Int8Type
	case uint8Marker:
		return structform.Uint8Type
	case int16Marker:
		return structform.Int16Type
	case int32Marker:
		return structform.Int32Type
	case int64Marker:
		return structform.Int64Type
	case float32Marker:
		return structform.Float32Type
	case float64Marker:
		return structform.Float64Type
	case highPrecMarker, stringMarker:
		return structform.StringType
	default:
		return structform.AnyType
	}
}

func (st state) withStep(s stateStep) state {
	st.stateStep = s
	return st
}
