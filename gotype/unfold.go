package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type Unfolder struct {
	unfoldCtx
}

type unfoldCtx struct {
	opts options

	// function stacks
	// sel unfoldTypeStack

	// reflection state stacks
	value  reflectValueStack
	key    keyStack
	idx    idxStack
	state  unfoldStateStack
	detail unfoldStateDetailStack

	unfolder unfolderStack
}

// unfold state used to handled next value
type unfoldState uint8

type unfoldStateDetail uint8

type unfoldType uint8

func println(v ...interface{}) {
	// fmt.Println(v...)
}

func printf(s string, v ...interface{}) {
	// fmt.Printf(s+"\n", v...)
}

//go:generate stringer -type=unfoldState
const (
	unfoldInvalidState unfoldState = iota
	unfoldStartState
	unfoldArrayState
	unfoldMapState
	unfoldStructState
	unfoldAssignState
)

//go:generate stringer -type=unfoldStateDetail
const (
	unfoldWaitStart unfoldStateDetail = iota
	unfoldWaitKey
	unfoldWaitElem
)

var (
	tInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	tString    = reflect.TypeOf("")
	tBool      = reflect.TypeOf(true)
	tInt       = reflect.TypeOf(int(0))
	tInt8      = reflect.TypeOf(int8(0))
	tInt16     = reflect.TypeOf(int16(0))
	tInt32     = reflect.TypeOf(int32(0))
	tInt64     = reflect.TypeOf(int64(0))
	tUint      = reflect.TypeOf(uint(0))
	tByte      = reflect.TypeOf(byte(0))
	tUint8     = reflect.TypeOf(uint8(0))
	tUint16    = reflect.TypeOf(uint16(0))
	tUint32    = reflect.TypeOf(uint32(0))
	tUint64    = reflect.TypeOf(uint64(0))
	tFloat32   = reflect.TypeOf(float32(0))
	tFloat64   = reflect.TypeOf(float64(0))
)

func NewUnfolder(to interface{}) (*Unfolder, error) {
	u := &Unfolder{}
	u.opts = options{
		tag: "struct",
	}

	if err := u.setTarget(to); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Unfolder) setTarget(to interface{}) error {
	if to == nil {
		return errNilInput
	}

	if u.trySetGotypeTarget(to) {
		return nil
	}

	vTo := reflect.ValueOf(to)
	k := vTo.Kind()
	isValid := to != nil && (k == reflect.Ptr || k == reflect.Map)
	if !isValid {
		return errRequiresPointer
	}

	vTo = chaseValuePointers(vTo)
	tTo := chaseTypePointers(vTo.Type())
	k = tTo.Kind()

	var state0 = unfoldInvalidState
	var detail0 = unfoldWaitStart

	switch k {
	case reflect.Map:
		state0 = unfoldMapState

		/* TODO:
		case reflect.Struct:
			state0 = unfoldStructState
		*/

	case reflect.Array, reflect.Slice:
		state0 = unfoldArrayState

	case reflect.Interface:
		if tTo == tInterface {
			state0 = unfoldAssignState
		}

	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		state0 = unfoldAssignState
	}

	// init processing stacks
	u.value.init(reflect.Value{})
	u.key.init()
	u.idx.init()
	u.state.init(unfoldStartState)
	u.state.push(state0)
	u.detail.init(unfoldWaitElem)
	u.detail.push(detail0)

	u.unfolder.init(newUnfolderRefl(&u.unfoldCtx))
	u.value.push(vTo)

	// u.unfoldCtx.setUnfoldGoTypes(to)

	return nil
}

func (u *unfoldCtx) OnObjectStart(len int, baseType structform.BaseType) error {
	return u.unfolder.current.OnObjectStart(len, baseType)
}

func (u *unfoldCtx) OnObjectFinished() error {
	return u.unfolder.current.OnObjectFinished()
}

func (u *unfoldCtx) OnKey(s string) error {
	return u.unfolder.current.OnKey(s)
}

func (u *unfoldCtx) OnArrayStart(len int, baseType structform.BaseType) error {
	return u.unfolder.current.OnArrayStart(len, baseType)
}

func (u *unfoldCtx) OnArrayFinished() error {
	return u.unfolder.current.OnArrayFinished()
}

func (u *unfoldCtx) OnNil() error {
	return u.unfolder.current.OnNil()
}

func (u *unfoldCtx) OnBool(b bool) error {
	return u.unfolder.current.OnBool(b)
}

func (u *unfoldCtx) OnString(s string) error {
	return u.unfolder.current.OnString(s)
}

func (u *unfoldCtx) OnInt8(i int8) error {
	return u.unfolder.current.OnInt8(i)
}

func (u *unfoldCtx) OnInt16(i int16) error {
	return u.unfolder.current.OnInt16(i)
}

func (u *unfoldCtx) OnInt32(i int32) error {
	return u.unfolder.current.OnInt32(i)
}

func (u *unfoldCtx) OnInt64(i int64) error {
	return u.unfolder.current.OnInt64(i)
}

func (u *unfoldCtx) OnInt(i int) error {
	return u.unfolder.current.OnInt(i)
}

func (u *unfoldCtx) OnByte(b byte) error {
	return u.unfolder.current.OnByte(b)
}

func (u *unfoldCtx) OnUint8(v uint8) error {
	return u.unfolder.current.OnUint8(v)
}

func (u *unfoldCtx) OnUint16(v uint16) error {
	return u.unfolder.current.OnUint16(v)
}

func (u *unfoldCtx) OnUint32(v uint32) error {
	return u.unfolder.current.OnUint32(v)
}

func (u *unfoldCtx) OnUint64(v uint64) error {
	return u.unfolder.current.OnUint64(v)
}

func (u *unfoldCtx) OnUint(v uint) error {
	return u.unfolder.current.OnUint(v)
}

func (u *unfoldCtx) OnFloat32(f float32) error {
	return u.unfolder.current.OnFloat32(f)
}

func (u *unfoldCtx) OnFloat64(f float64) error {
	return u.unfolder.current.OnFloat64(f)
}

func (u *unfoldCtx) setUnfoldGoTypes(v interface{}) {
	switch v.(type) {
	default:
	}

}
