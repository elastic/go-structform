package gotype

import (
	"fmt"
	"reflect"
	"unsafe"

	structform "github.com/urso/go-structform"
)

type unfolderIfc struct{}

var _singletonUnfolderIfc = &unfolderIfc{}

func newUnfolderIfc() *unfolderIfc {
	return _singletonUnfolderIfc
}

func (*unfolderIfc) ptr(ctx *unfoldCtx) *interface{} {
	return (*interface{})(ctx.ptr.current)
}

func (u *unfolderIfc) assign(ctx *unfoldCtx, v interface{}) error {
	*u.ptr(ctx) = v
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderIfc) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// v := *ptr
	// TODO: what if v != nil

	return unfoldIfcStartSubArray(ctx, l, baseType)
}

func (u *unfolderIfc) OnChildArrayDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubArray(ctx)
	if err == nil {
		err = u.assign(ctx, v)
	}
	return err
}

func (u *unfolderIfc) OnArrayFinished(c *unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderIfc) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	return unfoldIfcStartSubMap(ctx, l, baseType)
}

func (u *unfolderIfc) OnChildObjectDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubMap(ctx)
	if err == nil {
		err = u.assign(ctx, v)
	}
	return err
}

func (u *unfolderIfc) OnObjectFinished(*unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderIfc) OnKey(*unfoldCtx, string) error {
	return errUnsupported
}

func (u *unfolderIfc) OnNil(ctx *unfoldCtx) error { return u.assign(ctx, nil) }

func (u *unfolderIfc) OnBool(ctx *unfoldCtx, v bool) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnString(ctx *unfoldCtx, v string) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt8(ctx *unfoldCtx, v int8) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt16(ctx *unfoldCtx, v int16) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt32(ctx *unfoldCtx, v int32) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt64(ctx *unfoldCtx, v int64) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt(ctx *unfoldCtx, v int) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnByte(ctx *unfoldCtx, v byte) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnUint8(ctx *unfoldCtx, v uint8) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnUint16(ctx *unfoldCtx, v uint16) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnUint32(ctx *unfoldCtx, v uint32) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnUint64(ctx *unfoldCtx, v uint64) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnUint(ctx *unfoldCtx, v uint) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnFloat32(ctx *unfoldCtx, v float32) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnFloat64(ctx *unfoldCtx, v float64) error { return u.assign(ctx, v) }

// TODO: re-use preallocated nil object in unfoldCtx, so we don't have to allocate
func makeArrayPtr(ctx *unfoldCtx, l int, t structform.BaseType) interface{} {
	switch t {
	case structform.AnyType, structform.ZeroType:
		if l <= 0 {
			return new([]interface{})
		}
		tmp := make([]interface{}, l)
		return &tmp

	case structform.StringType:
		if l <= 0 {
			return new([]string)
		}
		tmp := make([]string, l)
		return &tmp

	case structform.BoolType:
		if l <= 0 {
			return new([]bool)
		}
		tmp := make([]bool, l)
		return &tmp

	case structform.IntType:
		if l <= 0 {
			return new([]int)
		}
		tmp := make([]int, l)
		return &tmp

	case structform.Int8Type:
		if l <= 0 {
			return new([]int8)
		}
		tmp := make([]int8, l)
		return &tmp

	case structform.Int16Type:
		if l <= 0 {
			return new([]int16)
		}
		tmp := make([]int16, l)
		return &tmp

	case structform.Int32Type:
		if l <= 0 {
			return new([]int32)
		}
		tmp := make([]int32, l)
		return &tmp

	case structform.Int64Type:
		if l <= 0 {
			return new([]int64)
		}
		tmp := make([]int64, l)
		return &tmp

	case structform.UintType:
		if l <= 0 {
			return new([]uint)
		}
		tmp := make([]uint, l)
		return &tmp

	case structform.ByteType, structform.Uint8Type:
		if l <= 0 {
			return new([]uint8)
		}
		tmp := make([]uint8, l)
		return &tmp

	case structform.Uint16Type:
		if l <= 0 {
			return new([]uint16)
		}
		tmp := make([]uint16, l)
		return &tmp

	case structform.Uint32Type:
		if l <= 0 {
			return new([]uint32)
		}
		tmp := make([]uint32, l)
		return &tmp

	case structform.Uint64Type:
		if l <= 0 {
			return new([]uint32)
		}
		tmp := make([]uint64, l)
		return &tmp

	case structform.Float32Type:
		if l <= 0 {
			return new([]float32)
		}
		tmp := make([]float32, l)
		return &tmp

	case structform.Float64Type:
		if l <= 0 {
			return new([]float64)
		}
		tmp := make([]float64, l)
		return &tmp

	default:
		panic("invalid type code")
		return nil
	}
}

func makeMapPtr(ctx *unfoldCtx, l int, t structform.BaseType) interface{} {
	switch t {
	case structform.AnyType, structform.ZeroType:
		return new(map[string]interface{})

	case structform.StringType:
		return new(map[string]string)

	case structform.BoolType:
		return new(map[string]bool)

	case structform.IntType:
		return new(map[string]int)

	case structform.Int8Type:
		return new(map[string]int8)

	case structform.Int16Type:
		return new(map[string]int16)

	case structform.Int32Type:
		return new(map[string]int32)

	case structform.Int64Type:
		return new(map[string]int64)

	case structform.UintType:
		return new(map[string]uint)

	case structform.ByteType, structform.Uint8Type:
		return new(map[string]uint8)

	case structform.Uint16Type:
		return new(map[string]uint16)

	case structform.Uint32Type:
		return new(map[string]uint32)

	case structform.Uint64Type:
		return new(map[string]uint64)

	case structform.Float32Type:
		return new(map[string]float32)

	case structform.Float64Type:
		return new(map[string]float64)

	default:
		panic("invalid type code")
		return nil
	}
}

func unfoldIfcStartSubArray(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	tmp := makeArrayPtr(ctx, l, baseType)
	if !ctx.trySetGotypeTarget(tmp) {
		// use reflection
		v := reflect.ValueOf(tmp)
		ctx.ptr.push(unsafe.Pointer(v.Pointer()))
		ctx.value.push(v.Elem())
		ctx.state.push(unfoldArrayState)
		ctx.detail.push(unfoldWaitStart)
		ctx.unfolder.push(newUnfolderRefl())
		ctx.baseType.push(structform.ZeroType)
	} else {
		// duplicate pointer, so we can assign after child-array is finished
		ctx.ptr.push(ctx.ptr.current)
		ctx.baseType.push(baseType)
	}

	return ctx.unfolder.current.OnArrayStart(ctx, l, baseType)
}

func unfoldIfcFinishSubArray(ctx *unfoldCtx) (interface{}, error) {
	child := ctx.ptr.pop()

	bt := ctx.baseType.pop()
	switch bt {
	case structform.ByteType, structform.Uint8Type:
		return *(*[]uint8)(child), nil

	case structform.UintType:
		return *(*[]uint)(child), nil

	case structform.IntType:
		return *(*[]int)(child), nil

	case structform.AnyType, structform.ZeroType:
		return *(*[]interface{})(child), nil

	default:
		fmt.Println("base type: ", bt)
		return nil, errTODO()
	}
}

func unfoldIfcStartSubMap(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	tmp := makeMapPtr(ctx, l, baseType)
	if !ctx.trySetGotypeTarget(tmp) {
		// use reflection
		v := reflect.ValueOf(tmp)
		ctx.ptr.push(unsafe.Pointer(v.Pointer()))
		ctx.value.push(v.Elem())
		ctx.state.push(unfoldMapState)
		ctx.detail.push(unfoldWaitStart)
		ctx.unfolder.push(newUnfolderRefl())
	} else {
		// duplicate pointer, so we can assign after child-object is finished
		ctx.ptr.push(ctx.ptr.current)
	}

	ctx.baseType.push(baseType)
	return ctx.unfolder.current.OnObjectStart(ctx, l, baseType)
}

func unfoldIfcFinishSubMap(ctx *unfoldCtx) (interface{}, error) {
	child := ctx.ptr.pop()
	switch ctx.baseType.pop() {
	case structform.AnyType:
		return *(*map[string]interface{})(child), nil

	case structform.ByteType, structform.Uint8Type:
		return *(*map[string]uint8)(child), nil

	case structform.StringType:
		return *(*map[string]string)(child), nil

	case structform.IntType:
		return *(*map[string]int)(child), nil

	default:
		return nil, errTODO()
	}
}
