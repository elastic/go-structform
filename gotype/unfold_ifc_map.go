package gotype

import structform "github.com/urso/go-structform"

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

func (u *unfolderMapIfc) OnChildObjectDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubMap(ctx)
	if err == nil {
		err = u.put(ctx, v)
	}
	return err
}
