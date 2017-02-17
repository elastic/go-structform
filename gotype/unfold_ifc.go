package gotype

import structform "github.com/urso/go-structform"

func (u *unfolderIfc) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	return unfoldIfcStartSubArray(ctx, l, baseType)
}

func (u *unfolderIfc) OnChildArrayDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubArray(ctx)
	if err == nil {
		err = u.assign(ctx, v)
	}
	return err
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
