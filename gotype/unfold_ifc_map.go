package gotype

import structform "github.com/urso/go-structform"

type unfolderMapIfc struct{}

var _singletonUnfolderMapIfc = &unfolderMapIfc{}

func newUnfolderMapIfc() *unfolderMapIfc {
	return _singletonUnfolderMapIfc
}

func (*unfolderMapIfc) ptr(ctx *unfoldCtx) *map[string]interface{} {
	return (*map[string]interface{})(ctx.ptr.current)
}

func (u *unfolderMapIfc) put(ctx *unfoldCtx, v interface{}) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
	if *to == nil {
		*to = map[string]interface{}{}
	}
	(*to)[ctx.key.pop()] = v

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapIfc) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		return errExpectedObjectNotArray
	case unfoldWaitKey:
		return errExpectedObjectKey
	}

	return unfoldIfcStartSubArray(ctx, l, baseType)
}

func (u *unfolderMapIfc) OnChildArrayDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubArray(ctx)
	if err == nil {
		err = u.put(ctx, v)
	}
	return err
}

func (u *unfolderMapIfc) OnArrayFinished(c *unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderMapIfc) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	dtl := &ctx.detail

	if dtl.current == unfoldWaitStart {
		// TODO: validate type

		dtl.current = unfoldWaitKey
		return nil
	}

	return unfoldIfcStartSubMap(ctx, l, baseType)
}

func (u *unfolderMapIfc) OnChildObjectDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubMap(ctx)
	if err == nil {
		err = u.put(ctx, v)
	}
	return err
}

func (u *unfolderMapIfc) OnObjectFinished(ctx *unfoldCtx) error {
	ctx.unfolder.pop()
	ctx.ptr.pop()
	ctx.detail.pop()

	return nil
}

func (u *unfolderMapIfc) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapIfc) OnNil(ctx *unfoldCtx) error { return u.put(ctx, nil) }

func (u *unfolderMapIfc) OnBool(ctx *unfoldCtx, v bool) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnString(ctx *unfoldCtx, v string) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt8(ctx *unfoldCtx, v int8) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt16(ctx *unfoldCtx, v int16) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt32(ctx *unfoldCtx, v int32) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt64(ctx *unfoldCtx, v int64) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt(ctx *unfoldCtx, v int) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnByte(ctx *unfoldCtx, v byte) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnUint8(ctx *unfoldCtx, v uint8) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnUint16(ctx *unfoldCtx, v uint16) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnUint32(ctx *unfoldCtx, v uint32) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnUint64(ctx *unfoldCtx, v uint64) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnUint(ctx *unfoldCtx, v uint) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnFloat32(ctx *unfoldCtx, v float32) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnFloat64(ctx *unfoldCtx, v float64) error { return u.put(ctx, v) }
