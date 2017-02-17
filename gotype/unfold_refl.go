package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type unfolderRefl struct {
}

var (
	invalidValue = reflect.Value{}
	trueValue    = reflect.ValueOf(true)
	falseValue   = reflect.ValueOf(false)
)

var singletionUnfolderRefl = &unfolderRefl{}

func newUnfolderRefl() *unfolderRefl {
	return singletionUnfolderRefl
}

func (unfolderRefl) OnChildArrayDone(_ *unfoldCtx) error  { return nil }
func (unfolderRefl) OnChildObjectDone(_ *unfoldCtx) error { return nil }

func (unfolderRefl) OnObjectStart(ctx *unfoldCtx, len int, baseType structform.BaseType) error {
	println("OnObjectStart:")
	ctx.printStacks()
	defer func() {
		println("  -> OnObjectStart")
		ctx.printStacks()
	}()

	st, dtl, v := &ctx.state, &ctx.detail, &ctx.value

	if dtl.current == unfoldWaitKey {
		return errStartObjectWaitingForKey
	}

	switch st.current {
	case unfoldAssignState:
		st.push(unfoldMapState)
		dtl.push(unfoldWaitKey)
		ctx.unfolder.push(newUnfolderRefl())

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
	ctx.unfolder.push(newUnfolderRefl())

	if v.current.Type().Elem() == tInterface {
		v.push(makeMap(baseType))
	} else {
		v.push(makeGoMap(v.current.Type().Elem().Elem()))
	}
	return nil

}

func (u unfolderRefl) OnObjectFinished(ctx *unfoldCtx) error {
	println("OnObjectFinished:")
	ctx.printStacks()
	defer func() {
		println("  -> OnObjectFinished")
		ctx.printStacks()
	}()

	st, dtl, v := &ctx.state, &ctx.detail, &ctx.value

	// scheduled waiting for new sub-object or assignment to nested interface, but
	// received finished signal for parent object -> drop state waiting for start of
	// nested object
	waitAssigned := (dtl.current == unfoldWaitStart &&
		(st.current == unfoldMapState || st.current == unfoldStructState)) ||
		(st.current == unfoldAssignState && dtl.current == unfoldWaitElem)
	if waitAssigned {
		st.pop()
		ctx.unfolder.pop()
		dtl.pop()
		v.pop()
	}

	dtl.pop()
	st.pop()
	ctx.unfolder.pop()
	value := v.pop()
	return u.onValue(ctx, value)
}

func (u unfolderRefl) OnKey(ctx *unfoldCtx, s string) error {
	st, dtl := &ctx.state, &ctx.detail

	println("OnKey: ", s)

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	switch st.current {
	case unfoldMapState:
		return u.onMapKey(ctx, s)
	case unfoldStructState:
		return u.onStructKey(ctx, s)
	default:
		return errTODO()
	}
}

func (u unfolderRefl) onMapKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	// dtl, v := &ctx.detail, &ctx.value
	// elem := chaseValue(v.current.MapIndex(reflect.ValueOf(key)))
	// println("  map element: ", elem)
	// u.tryAssignElem(ctx, elem)

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u unfolderRefl) onStructKey(ctx *unfoldCtx, key string) error {
	return errTODO()
}

func (u unfolderRefl) OnArrayStart(ctx *unfoldCtx, len int, baseType structform.BaseType) error {
	println("OnArrayStart:")
	ctx.printStacks()
	defer func() {
		println("  -> OnArrayStart")
		ctx.printStacks()
	}()

	st, dtl, idx, v := &ctx.state, &ctx.detail, &ctx.idx, &ctx.value

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
		ctx.unfolder.push(newUnfolderRefl())
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
			u.tryArrayAssign(ctx)
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
	ctx.unfolder.push(newUnfolderRefl())
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
func (u unfolderRefl) tryArrayAssign(ctx *unfoldCtx) {
	idx, v := &ctx.idx, &ctx.value

	if v.current.Len() > idx.current {
		elem := chaseValue(v.current.Index(idx.current))
		printf("  value at index %v: %v", idx.current, elem)
		u.tryAssignElem(ctx, elem)
	}
}

func (u unfolderRefl) tryAssignElem(ctx *unfoldCtx, elem reflect.Value) {
	st, dtl, v := &ctx.state, &ctx.detail, &ctx.value

	println("  try to assign kind: ", elem.Kind())

	switch elem.Kind() {
	case reflect.Slice, reflect.Array:
		if !elem.IsNil() {
			if elem.Kind() == reflect.Slice && elem.CanAddr() {
				elem.SetLen(0)
			}

			st.push(unfoldArrayState)
			ctx.unfolder.push(newUnfolderRefl())
			dtl.push(unfoldWaitStart)
			v.push(elem)
		}

	case reflect.Map:
		println("try assign map")
		if !elem.IsNil() {
			st.push(unfoldMapState)
			ctx.unfolder.push(newUnfolderRefl())
			dtl.push(unfoldWaitStart)
			v.push(elem)
		}

	case reflect.Struct:
		errTODO()

	case reflect.Interface:
		if !elem.IsNil() && elem.CanSet() {
			st.push(unfoldAssignState)
			ctx.unfolder.push(newUnfolderRefl())
			dtl.push(unfoldWaitElem)
			v.push(elem)
		}
	}
}

func (u unfolderRefl) OnArrayFinished(ctx *unfoldCtx) error {

	st, dtl, idx, v := &ctx.state, &ctx.detail, &ctx.idx, &ctx.value

	// scheduled waiting for new sub-array or assignment to nested interface, but
	// received finished signal for parent array -> drop state waiting for start of
	// nested array
	waitAssigned := (st.current == unfoldArrayState && dtl.current == unfoldWaitStart) ||
		(st.current == unfoldAssignState && dtl.current == unfoldWaitElem)
	if waitAssigned {
		st.pop()
		ctx.unfolder.pop()
		dtl.pop()
		v.pop()
	}

	dtl.pop()
	idx.pop()
	st.pop()
	ctx.unfolder.pop()
	value := v.pop()
	return u.onValue(ctx, value)

	// return errTODO()
}

func (u unfolderRefl) OnBool(ctx *unfoldCtx, v bool) error {
	if v == true {
		return u.onValue(ctx, trueValue)
	}
	return u.onValue(ctx, falseValue)
}

func (u unfolderRefl) onValue(ctx *unfoldCtx, v reflect.Value) error {
	println("onValue: ", v)
	ctx.printStacks()

	defer func() {
		println("  ->onValue")
		ctx.printStacks()
	}()

	switch ctx.state.current {
	case unfoldInvalidState:
		return errInvalidState
	case unfoldStartState:
		return nil
	case unfoldArrayState:
		return u.intoArray(ctx, v)
	case unfoldMapState:
		return u.intoMap(ctx, v)
	case unfoldStructState:
		return errTODO()
	case unfoldAssignState:
		ctx.state.pop()
		return u.assignValue(ctx, v)
	}
	return nil
}

func (u unfolderRefl) assignValue(ctx *unfoldCtx, v reflect.Value) error {
	// try to merge v into u.value.current

	value := ctx.value.current
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

func (u unfolderRefl) intoMap(ctx *unfoldCtx, v reflect.Value) error {
	dtl, val, key := &ctx.detail, &ctx.value, &ctx.key
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

	if val.current.IsNil() {
		val.current.Set(reflect.MakeMap(val.current.Type()))
	}

	k := reflect.ValueOf(key.pop())
	val.current.SetMapIndex(k, v)
	dtl.current = unfoldWaitKey
	return nil
}

func (u unfolderRefl) intoStruct(ctx *unfoldCtx, v reflect.Value) error {
	return errTODO()
}

func (u unfolderRefl) intoArray(ctx *unfoldCtx, v reflect.Value) error {
	val, idx := &ctx.value, &ctx.idx
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

	u.tryArrayAssign(ctx)
	return nil
}

func (u unfolderRefl) OnNil(ctx *unfoldCtx) error {
	return u.onValue(ctx, invalidValue)
}

func (u unfolderRefl) OnString(ctx *unfoldCtx, v string) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnInt(ctx *unfoldCtx, v int) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnByte(ctx *unfoldCtx, v byte) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnUint(ctx *unfoldCtx, v uint) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.onValue(ctx, reflect.ValueOf(v))
}

func (u unfolderRefl) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.onValue(ctx, reflect.ValueOf(v))
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
	println("    unfolder: ", len(u.unfolder.stack)+1)
}
