package gotype

import (
	"fmt"
	"reflect"

	structform "github.com/urso/go-structform"
)

type Unfolder struct {
	depth int
	opts  options

	value  reflectValueStack
	key    keyStack
	idx    idxStack
	state  unfoldStateStack
	detail unfoldStateDetailStack
}

// unfold state used to handled next value
type unfoldState uint8

type unfoldStateDetail uint8

func println(v ...interface{}) {
	fmt.Println(v...)
}

func printf(s string, v ...interface{}) {
	fmt.Printf(s+"\n", v...)
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
	u := &Unfolder{
		opts: options{
			tag: "struct",
		},
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
	u.value.init(vTo)
	u.key.init()
	u.idx.init()
	u.state.init(unfoldStartState)
	u.state.push(state0)
	u.detail.init(unfoldWaitElem)
	u.detail.push(detail0)

	return nil
}

func (u *Unfolder) OnObjectStart(len int, baseType structform.BaseType) error {
	println("OnObjectStart:")
	printStacks(u)
	defer func() {
		println("  -> OnObjectStart")
		printStacks(u)
	}()

	st, dtl, v := &u.state, &u.detail, &u.value

	if dtl.current == unfoldWaitKey {
		return errStartObjectWaitingForKey
	}

	switch st.current {
	case unfoldAssignState:
		st.push(unfoldMapState)
		dtl.push(unfoldWaitKey)

		v.push(makeMap(baseType))
		return nil

	case unfoldMapState:
		if dtl.current == unfoldWaitStart {
			// we did wait for 'OnObjectStart' for already scheduled map/struct
			// let's advance state and initialize key stack
			dtl.current = unfoldWaitKey
			return nil
		}

		break

	case unfoldArrayState:
		if dtl.current == unfoldWaitStart {
			return errExpectedArrayNotObject
		}

		break

	default:
		println("state: ", st.current)
		return errTODO()
	}

	// start new object to be inserted into current array/map
	println("  current state: ", st.current)
	println("  current type: ", v.current.Type())
	println("    kind: ", v.current.Kind())
	println("    elem type: ", v.current.Type().Elem())

	st.push(unfoldMapState)
	dtl.push(unfoldWaitKey)

	if v.current.Type().Elem() == tInterface {
		v.push(makeMap(baseType))
	} else {
		v.push(makeGoMap(v.current.Type().Elem().Elem()))
	}
	return nil

}

func (u *Unfolder) OnObjectFinished() error {
	println("OnObjectFinished:")
	printStacks(u)
	defer func() {
		println("  -> OnObjectFinished")
		printStacks(u)
	}()

	st, dtl, v := &u.state, &u.detail, &u.value

	// scheduled waiting for new sub-object or assignment to nested interface, but
	// received finished signal for parent object -> drop state waiting for start of
	// nested object
	waitAssigned := (dtl.current == unfoldWaitStart &&
		(st.current == unfoldMapState || st.current == unfoldStructState)) ||
		(st.current == unfoldAssignState && dtl.current == unfoldWaitElem)
	if waitAssigned {
		st.pop()
		dtl.pop()
		v.pop()
	}

	dtl.pop()
	st.pop()
	value := v.pop()
	return u.onValue(value)
}

func (u *Unfolder) OnKey(s string) error {
	st, dtl := &u.state, &u.detail

	println("OnKey: ", s)

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	switch st.current {
	case unfoldMapState:
		return u.onMapKey(s)
	case unfoldStructState:
		return u.onStructKey(s)
	default:
		return errTODO()
	}
}

func (u *Unfolder) onMapKey(key string) error {
	dtl, v := &u.detail, &u.value

	u.key.push(key)
	dtl.current = unfoldWaitElem
	elem := chaseValue(v.current.MapIndex(reflect.ValueOf(key)))
	println("  map element: ", elem)
	u.tryAssignElem(elem)
	return nil
}

func (u *Unfolder) onStructKey(key string) error {
	return errTODO()
}

func (u *Unfolder) OnArrayStart(len int, baseType structform.BaseType) error {
	println("OnArrayStart:")
	printStacks(u)
	defer func() {
		println("  -> OnArrayStart")
		printStacks(u)
	}()

	st, dtl, idx, v := &u.state, &u.detail, &u.idx, &u.value

	if dtl.current == unfoldWaitKey {
		return errStartArrayWaitingForKey
	}

	switch st.current {
	case unfoldAssignState:
		// check type of new array:
		println("  current state: ", st.current)
		println("  current type: ", v.current.Type())
		println("    kind: ", v.current.Kind())

		st.push(unfoldArrayState)
		dtl.push(unfoldWaitElem)
		idx.push()

		v.push(makeArray(len, baseType))
		// v.push(makeSlice(len, v.current.Type().Elem().Elem()))
		return nil

	case unfoldArrayState:
		if dtl.current == unfoldWaitStart {
			// we did wait for 'OnArrayStart' for already scheduled array
			// let's advance state and initialize array index stack
			dtl.current = unfoldWaitElem
			idx.push()
			u.tryArrayAssign()
			return nil
		}
		break

	case unfoldMapState:
		if dtl.current == unfoldWaitStart {
			return errExpectedObjectNotArray
		}
		break

	default:
		return errTODO()
	}

	// start new array to be inserted into current map/array
	println("  current state: ", st.current)
	println("  current type: ", v.current.Type())
	println("    kind: ", v.current.Kind())
	println("    elem type: ", v.current.Type().Elem())

	st.push(unfoldArrayState)
	dtl.push(unfoldWaitElem)
	idx.push()

	if v.current.Type().Elem() == tInterface {
		v.push(makeArray(len, baseType))
	} else {
		v.push(makeSlice(len, v.current.Type().Elem().Elem()))
	}
	return nil

}

// tryArrayAssign checks if value on stack has some element to be overwritten.
// If so, it does push the value on the processing stack waiting for the
// assignment to happen (not, OnArrayFinished might be received on input stream and
// cleanup stacks in this case).
func (u *Unfolder) tryArrayAssign() {
	idx, v := &u.idx, &u.value

	if v.current.Len() > idx.current {
		elem := chaseValue(v.current.Index(idx.current))
		printf("  value at index %v: %v", idx.current, elem)
		u.tryAssignElem(elem)
	}
}

func (u *Unfolder) tryAssignElem(elem reflect.Value) {
	st, dtl, v := &u.state, &u.detail, &u.value

	println("  try to assign kind: ", elem.Kind())

	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		if !elem.IsNil() {
			st.push(unfoldArrayState)
			dtl.push(unfoldWaitStart)
			v.push(elem)
		}

	case reflect.Map:
		println("try assign map")
		if !elem.IsNil() {
			st.push(unfoldMapState)
			dtl.push(unfoldWaitStart)
			v.push(elem)
		}

	case reflect.Struct:
		errTODO()

	case reflect.Interface:
		if !elem.IsNil() && elem.CanSet() {
			st.push(unfoldAssignState)
			dtl.push(unfoldWaitElem)
			v.push(elem)
		}
	}
}

func (u *Unfolder) OnArrayFinished() error {

	st, dtl, idx, v := &u.state, &u.detail, &u.idx, &u.value

	// scheduled waiting for new sub-array or assignment to nested interface, but
	// received finished signal for parent array -> drop state waiting for start of
	// nested array
	waitAssigned := (st.current == unfoldArrayState && dtl.current == unfoldWaitStart) ||
		(st.current == unfoldAssignState && dtl.current == unfoldWaitElem)
	if waitAssigned {
		st.pop()
		dtl.pop()
		v.pop()
	}

	dtl.pop()
	idx.pop()
	st.pop()
	value := v.pop()
	return u.onValue(value)

	// return errTODO()
}

func makeArray(len int, baseType structform.BaseType) reflect.Value {
	return makeSlice(len, lookupGoType(baseType))
}

func makeSlice(len int, t reflect.Type) reflect.Value {
	st := reflect.SliceOf(t)
	if len < 0 {
		return reflect.New(st).Elem()
	}
	return reflect.MakeSlice(st, len, len)
}

func makeMap(baseType structform.BaseType) reflect.Value {
	return makeGoMap(lookupGoType(baseType))
}

func makeGoMap(t reflect.Type) reflect.Value {
	mt := reflect.MapOf(tString, t)
	println("  make map: ", mt)
	return reflect.MakeMap(mt)
}

func (u *Unfolder) onValue(v reflect.Value) error {
	println("onValue: ", v)
	printStacks(u)

	defer func() {
		println("  ->onValue")
		printStacks(u)
	}()

	switch u.state.current {
	case unfoldInvalidState:
		return errInvalidState
	case unfoldStartState:
		return nil
	case unfoldArrayState:
		return u.intoArray(v)
	case unfoldMapState:
		return u.intoMap(v)
	case unfoldStructState:
		return errTODO()
	case unfoldAssignState:
		u.state.pop()
		return u.assignValue(v)
	}
	return nil
}

func (u *Unfolder) assignValue(v reflect.Value) error {
	// try to merge v into u.value.current

	value := u.value.current
	if !v.IsValid() {
		value.Set(reflect.Zero(value.Type()))
		return nil
	}

	if !v.Type().ConvertibleTo(value.Type()) {
		return errIncompatibleTypes
	}

	value.Set(v.Convert(value.Type()))
	return nil
}

func (u *Unfolder) intoMap(v reflect.Value) error {
	dtl, val, key := &u.detail, &u.value, &u.key
	tElem := val.current.Type().Elem()

	switch {
	case v == invalidValue:
		v = reflect.Zero(tElem)

	case tElem.Kind() == reflect.Ptr:
		bt := chaseTypePointers(tElem)
		v = pointerize(tElem, bt, v.Convert(bt))

	default:
		v = v.Convert(tElem)
	}

	k := reflect.ValueOf(key.pop())
	val.current.SetMapIndex(k, v)
	dtl.current = unfoldWaitKey
	return nil
}

func (u *Unfolder) intoStruct(v reflect.Value) error {
	return errTODO()
}

func (u *Unfolder) intoArray(v reflect.Value) error {
	val, idx := &u.value, &u.idx
	tElem := val.current.Type().Elem()

	switch {
	case v == invalidValue:
		v = reflect.Zero(tElem)

	case tElem.Kind() == reflect.Ptr:
		bt := chaseTypePointers(tElem)
		v = pointerize(tElem, bt, v.Convert(bt))

	default:
		v = v.Convert(tElem)
	}

	if val.current.Len() <= idx.current {
		// use append
		println("  append value")
		if !val.current.CanSet() {
			tmp := makeSlice(val.current.Len()+1, val.current.Type().Elem())
			reflect.Copy(tmp, val.current)
			val.current = tmp

			elem := val.current.Index(idx.current)
			elem.Set(v)
		} else {
			val.current.Set(reflect.Append(val.current, v))
		}
	} else {
		println("  assign array value")
		elem := val.current.Index(idx.current)
		elem.Set(v)
	}
	idx.current++

	u.tryArrayAssign()
	return nil
}

func lookupGoType(b structform.BaseType) reflect.Type {
	switch b {
	case structform.ByteType:
		return tByte
	case structform.StringType:
		return tString
	case structform.BoolType:
		return tBool
	case structform.IntType:
		return tInt
	case structform.Int8Type:
		return tInt8
	case structform.Int16Type:
		return tInt16
	case structform.Int32Type:
		return tInt32
	case structform.Int64Type:
		return tInt64
	case structform.UintType:
		return tUint
	case structform.Uint8Type:
		return tUint8
	case structform.Uint16Type:
		return tUint16
	case structform.Uint32Type:
		return tUint32
	case structform.Uint64Type:
		return tUint64
	case structform.Float32Type:
		return tFloat32
	case structform.Float64Type:
		return tFloat64
	default:
		return tInterface
	}

}

func printStacks(u *Unfolder) {
	println("  stacks:")
	println("    state: ", u.state.stack, u.state.current)
	println("    detail: ", u.detail.stack, u.detail.current)
	println("    value: ", u.value.stack, u.value.current)
	println("    key: ", u.key.stack, u.key.current)
	println("    idx: ", u.idx.stack, u.idx.current)
}
