package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type unfolderRefl struct {
	*unfoldCtx
}

var (
	invalidValue = reflect.Value{}
	trueValue    = reflect.ValueOf(true)
	falseValue   = reflect.ValueOf(false)
)

func newUnfolderRefl(ctx *unfoldCtx) *unfolderRefl {
	return &unfolderRefl{ctx}
}

func (u unfolderRefl) OnObjectStart(len int, baseType structform.BaseType) error {
	println("OnObjectStart:")
	u.printStacks()
	defer func() {
		println("  -> OnObjectStart")
		u.printStacks()
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

func (u unfolderRefl) OnObjectFinished() error {
	println("OnObjectFinished:")
	u.printStacks()
	defer func() {
		println("  -> OnObjectFinished")
		u.printStacks()
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

func (u unfolderRefl) OnKey(s string) error {
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

func (u unfolderRefl) onMapKey(key string) error {
	dtl, v := &u.detail, &u.value

	u.key.push(key)
	dtl.current = unfoldWaitElem
	elem := chaseValue(v.current.MapIndex(reflect.ValueOf(key)))
	println("  map element: ", elem)
	u.tryAssignElem(elem)
	return nil
}

func (u unfolderRefl) onStructKey(key string) error {
	return errTODO()
}

func (u unfolderRefl) OnArrayStart(len int, baseType structform.BaseType) error {
	println("OnArrayStart:")
	u.printStacks()
	defer func() {
		println("  -> OnArrayStart")
		u.printStacks()
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
		idx.push(0)

		v.push(makeArray(len, baseType))
		// v.push(makeSlice(len, v.current.Type().Elem().Elem()))
		return nil

	case unfoldArrayState:
		if dtl.current == unfoldWaitStart {
			// we did wait for 'OnArrayStart' for already scheduled array
			// let's advance state and initialize array index stack
			dtl.current = unfoldWaitElem
			idx.push(0)
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
	idx.push(0)

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
func (u unfolderRefl) tryArrayAssign() {
	idx, v := &u.idx, &u.value

	if v.current.Len() > idx.current {
		elem := chaseValue(v.current.Index(idx.current))
		printf("  value at index %v: %v", idx.current, elem)
		u.tryAssignElem(elem)
	}
}

func (u unfolderRefl) tryAssignElem(elem reflect.Value) {
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

func (u unfolderRefl) OnArrayFinished() error {

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

func (u unfolderRefl) OnBool(v bool) error {
	if v == true {
		return u.onValue(trueValue)
	}
	return u.onValue(falseValue)
}

func (u unfolderRefl) OnNil() error              { return u.onValue(invalidValue) }
func (u unfolderRefl) OnString(v string) error   { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnInt8(v int8) error       { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnInt16(v int16) error     { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnInt32(v int32) error     { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnInt64(v int64) error     { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnInt(v int) error         { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnByte(v byte) error       { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnUint8(v uint8) error     { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnUint16(v uint16) error   { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnUint32(v uint32) error   { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnUint64(v uint64) error   { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnUint(v uint) error       { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnFloat32(v float32) error { return u.onValue(reflect.ValueOf(v)) }
func (u unfolderRefl) OnFloat64(v float64) error { return u.onValue(reflect.ValueOf(v)) }

func (u unfolderRefl) onValue(v reflect.Value) error {
	println("onValue: ", v)
	u.printStacks()

	defer func() {
		println("  ->onValue")
		u.printStacks()
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

func (u unfolderRefl) assignValue(v reflect.Value) error {
	// try to merge v into u.value.current

	value := u.value.current
	if !v.IsValid() {
		value.Set(reflect.Zero(value.Type()))
		return nil
	}

	if !v.Type().ConvertibleTo(value.Type()) {
		return errIncompatibleTypes
	}

	if !v.Type().AssignableTo(value.Type()) {
		v = v.Convert(value.Type())
	}
	value.Set(v)
	return nil
}

func (u unfolderRefl) intoMap(v reflect.Value) error {
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

func (u unfolderRefl) intoStruct(v reflect.Value) error {
	return errTODO()
}

func (u unfolderRefl) intoArray(v reflect.Value) error {
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

func makeMap(baseType structform.BaseType) reflect.Value {
	return makeGoMap(lookupGoType(baseType))
}

func makeGoMap(t reflect.Type) reflect.Value {
	mt := reflect.MapOf(tString, t)
	println("  make map: ", mt)
	return reflect.MakeMap(mt)
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

func (u *unfoldCtx) printStacks() {
	println("  stacks:")
	println("    state: ", u.state.stack, u.state.current)
	println("    detail: ", u.detail.stack, u.detail.current)
	println("    value: ", u.value.stack, u.value.current)
	println("    key: ", u.key.stack, u.key.current)
	println("    idx: ", u.idx.stack, u.idx.current)
}
