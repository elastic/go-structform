package gotype

import "github.com/urso/go-structform"

func (u *unfolderArrIfc) OnChildArrayDone(ctx *unfoldCtx) error {
	v, err := unfoldIfcFinishSubArray(ctx)
	if err == nil {
		err = u.append(ctx, v)
	}
	return err
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
