package gotype

import "github.com/urso/go-structform"

type unfolderArrIfc struct{}

var _singletonUnfolderArrIfc = &unfolderArrIfc{}

func newUnfolderArrIfc() *unfolderArrIfc {
	return _singletonUnfolderArrIfc
}

func (*unfolderArrIfc) ptr(ctx *unfoldCtx) *[]interface{} {
	return (*[]interface{})(ctx.ptr.current)
}

func (u *unfolderArrIfc) append(ctx *unfoldCtx, v interface{}) error {
	idx := &ctx.idx
	to := (*[]interface{})(ctx.ptr.current)

	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

func (u *unfolderArrIfc) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	dtl := &ctx.detail

	if dtl.current == unfoldWaitStart {
		// TODO: validate type

		to := u.ptr(ctx)
		if l < 0 {
			l = 0
		}

		if *to == nil && l > 0 {
			*to = make([]interface{}, l)
		} else if l < len(*to) {
			*to = (*to)[:l]
		}

		dtl.current = unfoldWaitElem
		ctx.idx.push(0)
		return nil
	}

	// v := *ptr
	// TODO: what if v != nil

	return unfoldIfcStartSubArray(ctx, l, baseType)
}

func (u *unfolderArrIfc) OnChildArrayDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubArray(ctx)
	if err == nil {
		err = u.append(ctx, v)
	}
	return err
}

func (u *unfolderArrIfc) OnArrayFinished(ctx *unfoldCtx) error {
	ctx.unfolder.pop()
	ctx.ptr.pop()
	if ctx.detail.pop() != unfoldWaitStart {
		ctx.idx.pop()
	}

	return nil
}

func (u *unfolderArrIfc) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	return unfoldIfcStartSubMap(ctx, l, baseType)
}

func (u *unfolderArrIfc) OnChildObjectDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubMap(ctx)
	if err == nil {
		err = u.append(ctx, v)
	}
	return err
}

func (u *unfolderArrIfc) OnObjectFinished(*unfoldCtx) error {
	return errUnsupported
}

func (u *unfolderArrIfc) OnKey(*unfoldCtx, string) error {
	return errUnsupported
}

func (u *unfolderArrIfc) OnNil(ctx *unfoldCtx) error { return u.append(ctx, nil) }

func (u *unfolderArrIfc) OnBool(ctx *unfoldCtx, v bool) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnString(ctx *unfoldCtx, v string) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnInt8(ctx *unfoldCtx, v int8) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnInt16(ctx *unfoldCtx, v int16) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnInt32(ctx *unfoldCtx, v int32) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnInt64(ctx *unfoldCtx, v int64) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnInt(ctx *unfoldCtx, v int) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnByte(ctx *unfoldCtx, v byte) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnUint8(ctx *unfoldCtx, v uint8) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnUint16(ctx *unfoldCtx, v uint16) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnUint32(ctx *unfoldCtx, v uint32) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnUint64(ctx *unfoldCtx, v uint64) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnUint(ctx *unfoldCtx, v uint) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnFloat32(ctx *unfoldCtx, v float32) error { return u.append(ctx, v) }

func (u *unfolderArrIfc) OnFloat64(ctx *unfoldCtx, v float64) error { return u.append(ctx, v) }
