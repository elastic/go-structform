// This file has been generated from 'unfol_arr.yml', do not edit
package gotype

import (
	"sync"

	"github.com/urso/go-structform"
)

type unfolderArrBool struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]bool
}

var unfolderArrBoolPool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrBool{}
	},
}

func newUnfolderArrBool(ctx *unfoldCtx, to *[]bool) *unfolderArrBool {
	u := unfolderArrBoolPool.Get().(*unfolderArrBool)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrBool{ctx: ctx, to: to}
}

func (u *unfolderArrBool) free() {
	*u = unfolderArrBool{}
	unfolderArrBoolPool.Put(u)
}

func (u *unfolderArrBool) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]bool, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrBool) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrBool) append(v bool) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrString struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]string
}

var unfolderArrStringPool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrString{}
	},
}

func newUnfolderArrString(ctx *unfoldCtx, to *[]string) *unfolderArrString {
	u := unfolderArrStringPool.Get().(*unfolderArrString)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrString{ctx: ctx, to: to}
}

func (u *unfolderArrString) free() {
	*u = unfolderArrString{}
	unfolderArrStringPool.Put(u)
}

func (u *unfolderArrString) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]string, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrString) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrString) append(v string) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrUint struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]uint
}

var unfolderArrUintPool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrUint{}
	},
}

func newUnfolderArrUint(ctx *unfoldCtx, to *[]uint) *unfolderArrUint {
	u := unfolderArrUintPool.Get().(*unfolderArrUint)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrUint{ctx: ctx, to: to}
}

func (u *unfolderArrUint) free() {
	*u = unfolderArrUint{}
	unfolderArrUintPool.Put(u)
}

func (u *unfolderArrUint) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]uint, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrUint) append(v uint) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrUint8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]uint8
}

var unfolderArrUint8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrUint8{}
	},
}

func newUnfolderArrUint8(ctx *unfoldCtx, to *[]uint8) *unfolderArrUint8 {
	u := unfolderArrUint8Pool.Get().(*unfolderArrUint8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrUint8{ctx: ctx, to: to}
}

func (u *unfolderArrUint8) free() {
	*u = unfolderArrUint8{}
	unfolderArrUint8Pool.Put(u)
}

func (u *unfolderArrUint8) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]uint8, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint8) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrUint8) append(v uint8) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrUint16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]uint16
}

var unfolderArrUint16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrUint16{}
	},
}

func newUnfolderArrUint16(ctx *unfoldCtx, to *[]uint16) *unfolderArrUint16 {
	u := unfolderArrUint16Pool.Get().(*unfolderArrUint16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrUint16{ctx: ctx, to: to}
}

func (u *unfolderArrUint16) free() {
	*u = unfolderArrUint16{}
	unfolderArrUint16Pool.Put(u)
}

func (u *unfolderArrUint16) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]uint16, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint16) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrUint16) append(v uint16) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrUint32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]uint32
}

var unfolderArrUint32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrUint32{}
	},
}

func newUnfolderArrUint32(ctx *unfoldCtx, to *[]uint32) *unfolderArrUint32 {
	u := unfolderArrUint32Pool.Get().(*unfolderArrUint32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrUint32{ctx: ctx, to: to}
}

func (u *unfolderArrUint32) free() {
	*u = unfolderArrUint32{}
	unfolderArrUint32Pool.Put(u)
}

func (u *unfolderArrUint32) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]uint32, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint32) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrUint32) append(v uint32) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrUint64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]uint64
}

var unfolderArrUint64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrUint64{}
	},
}

func newUnfolderArrUint64(ctx *unfoldCtx, to *[]uint64) *unfolderArrUint64 {
	u := unfolderArrUint64Pool.Get().(*unfolderArrUint64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrUint64{ctx: ctx, to: to}
}

func (u *unfolderArrUint64) free() {
	*u = unfolderArrUint64{}
	unfolderArrUint64Pool.Put(u)
}

func (u *unfolderArrUint64) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]uint64, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint64) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrUint64) append(v uint64) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrInt struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]int
}

var unfolderArrIntPool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrInt{}
	},
}

func newUnfolderArrInt(ctx *unfoldCtx, to *[]int) *unfolderArrInt {
	u := unfolderArrIntPool.Get().(*unfolderArrInt)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrInt{ctx: ctx, to: to}
}

func (u *unfolderArrInt) free() {
	*u = unfolderArrInt{}
	unfolderArrIntPool.Put(u)
}

func (u *unfolderArrInt) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]int, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrInt) append(v int) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrInt8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]int8
}

var unfolderArrInt8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrInt8{}
	},
}

func newUnfolderArrInt8(ctx *unfoldCtx, to *[]int8) *unfolderArrInt8 {
	u := unfolderArrInt8Pool.Get().(*unfolderArrInt8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrInt8{ctx: ctx, to: to}
}

func (u *unfolderArrInt8) free() {
	*u = unfolderArrInt8{}
	unfolderArrInt8Pool.Put(u)
}

func (u *unfolderArrInt8) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]int8, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt8) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrInt8) append(v int8) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrInt16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]int16
}

var unfolderArrInt16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrInt16{}
	},
}

func newUnfolderArrInt16(ctx *unfoldCtx, to *[]int16) *unfolderArrInt16 {
	u := unfolderArrInt16Pool.Get().(*unfolderArrInt16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrInt16{ctx: ctx, to: to}
}

func (u *unfolderArrInt16) free() {
	*u = unfolderArrInt16{}
	unfolderArrInt16Pool.Put(u)
}

func (u *unfolderArrInt16) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]int16, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt16) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrInt16) append(v int16) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrInt32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]int32
}

var unfolderArrInt32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrInt32{}
	},
}

func newUnfolderArrInt32(ctx *unfoldCtx, to *[]int32) *unfolderArrInt32 {
	u := unfolderArrInt32Pool.Get().(*unfolderArrInt32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrInt32{ctx: ctx, to: to}
}

func (u *unfolderArrInt32) free() {
	*u = unfolderArrInt32{}
	unfolderArrInt32Pool.Put(u)
}

func (u *unfolderArrInt32) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]int32, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt32) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrInt32) append(v int32) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrInt64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]int64
}

var unfolderArrInt64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrInt64{}
	},
}

func newUnfolderArrInt64(ctx *unfoldCtx, to *[]int64) *unfolderArrInt64 {
	u := unfolderArrInt64Pool.Get().(*unfolderArrInt64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrInt64{ctx: ctx, to: to}
}

func (u *unfolderArrInt64) free() {
	*u = unfolderArrInt64{}
	unfolderArrInt64Pool.Put(u)
}

func (u *unfolderArrInt64) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]int64, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt64) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrInt64) append(v int64) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrFloat32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]float32
}

var unfolderArrFloat32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrFloat32{}
	},
}

func newUnfolderArrFloat32(ctx *unfoldCtx, to *[]float32) *unfolderArrFloat32 {
	u := unfolderArrFloat32Pool.Get().(*unfolderArrFloat32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrFloat32{ctx: ctx, to: to}
}

func (u *unfolderArrFloat32) free() {
	*u = unfolderArrFloat32{}
	unfolderArrFloat32Pool.Put(u)
}

func (u *unfolderArrFloat32) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]float32, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrFloat32) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrFloat32) append(v float32) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

type unfolderArrFloat64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *[]float64
}

var unfolderArrFloat64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderArrFloat64{}
	},
}

func newUnfolderArrFloat64(ctx *unfoldCtx, to *[]float64) *unfolderArrFloat64 {
	u := unfolderArrFloat64Pool.Get().(*unfolderArrFloat64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderArrFloat64{ctx: ctx, to: to}
}

func (u *unfolderArrFloat64) free() {
	*u = unfolderArrFloat64{}
	unfolderArrFloat64Pool.Put(u)
}

func (u *unfolderArrFloat64) OnArrayStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	if l > 0 && *u.to == nil {
		*u.to = make([]float64, l)
	}

	dtl.current = unfoldWaitElem
	u.ctx.idx.push(0)
	return nil
}

func (u *unfolderArrFloat64) OnArrayFinished() error {
	u.ctx.idx.pop()
	u.ctx.detail.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderArrFloat64) append(v float64) error {
	idx := &u.ctx.idx
	if len(*u.to) <= idx.current {
		*u.to = append(*u.to, v)
	} else {
		(*u.to)[idx.current] = v
	}
	idx.current++
	return nil
}

func (u *unfolderArrBool) OnNil() error        { return u.append(false) }
func (u *unfolderArrBool) OnBool(v bool) error { return u.append(v) }

func (u *unfolderArrString) OnNil() error            { return u.append("") }
func (u *unfolderArrString) OnString(v string) error { return u.append(v) }

func (u *unfolderArrUint) OnNil() error              { return u.append(0) }
func (u *unfolderArrUint) OnInt8(v int8) error       { return u.append(uint(v)) }
func (u *unfolderArrUint) OnInt16(v int16) error     { return u.append(uint(v)) }
func (u *unfolderArrUint) OnInt32(v int32) error     { return u.append(uint(v)) }
func (u *unfolderArrUint) OnInt64(v int64) error     { return u.append(uint(v)) }
func (u *unfolderArrUint) OnInt(v int) error         { return u.append(uint(v)) }
func (u *unfolderArrUint) OnByte(v byte) error       { return u.append(uint(v)) }
func (u *unfolderArrUint) OnUint8(v uint8) error     { return u.append(uint(v)) }
func (u *unfolderArrUint) OnUint16(v uint16) error   { return u.append(uint(v)) }
func (u *unfolderArrUint) OnUint32(v uint32) error   { return u.append(uint(v)) }
func (u *unfolderArrUint) OnUint64(v uint64) error   { return u.append(uint(v)) }
func (u *unfolderArrUint) OnUint(v uint) error       { return u.append(uint(v)) }
func (u *unfolderArrUint) OnFloat32(v float32) error { return u.append(uint(v)) }
func (u *unfolderArrUint) OnFloat64(v float64) error { return u.append(uint(v)) }

func (u *unfolderArrUint8) OnNil() error              { return u.append(0) }
func (u *unfolderArrUint8) OnInt8(v int8) error       { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnInt16(v int16) error     { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnInt32(v int32) error     { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnInt64(v int64) error     { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnInt(v int) error         { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnByte(v byte) error       { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnUint8(v uint8) error     { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnUint16(v uint16) error   { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnUint32(v uint32) error   { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnUint64(v uint64) error   { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnUint(v uint) error       { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnFloat32(v float32) error { return u.append(uint8(v)) }
func (u *unfolderArrUint8) OnFloat64(v float64) error { return u.append(uint8(v)) }

func (u *unfolderArrUint16) OnNil() error              { return u.append(0) }
func (u *unfolderArrUint16) OnInt8(v int8) error       { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnInt16(v int16) error     { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnInt32(v int32) error     { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnInt64(v int64) error     { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnInt(v int) error         { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnByte(v byte) error       { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnUint8(v uint8) error     { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnUint16(v uint16) error   { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnUint32(v uint32) error   { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnUint64(v uint64) error   { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnUint(v uint) error       { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnFloat32(v float32) error { return u.append(uint16(v)) }
func (u *unfolderArrUint16) OnFloat64(v float64) error { return u.append(uint16(v)) }

func (u *unfolderArrUint32) OnNil() error              { return u.append(0) }
func (u *unfolderArrUint32) OnInt8(v int8) error       { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnInt16(v int16) error     { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnInt32(v int32) error     { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnInt64(v int64) error     { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnInt(v int) error         { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnByte(v byte) error       { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnUint8(v uint8) error     { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnUint16(v uint16) error   { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnUint32(v uint32) error   { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnUint64(v uint64) error   { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnUint(v uint) error       { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnFloat32(v float32) error { return u.append(uint32(v)) }
func (u *unfolderArrUint32) OnFloat64(v float64) error { return u.append(uint32(v)) }

func (u *unfolderArrUint64) OnNil() error              { return u.append(0) }
func (u *unfolderArrUint64) OnInt8(v int8) error       { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnInt16(v int16) error     { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnInt32(v int32) error     { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnInt64(v int64) error     { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnInt(v int) error         { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnByte(v byte) error       { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnUint8(v uint8) error     { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnUint16(v uint16) error   { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnUint32(v uint32) error   { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnUint64(v uint64) error   { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnUint(v uint) error       { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnFloat32(v float32) error { return u.append(uint64(v)) }
func (u *unfolderArrUint64) OnFloat64(v float64) error { return u.append(uint64(v)) }

func (u *unfolderArrInt) OnNil() error              { return u.append(0) }
func (u *unfolderArrInt) OnInt8(v int8) error       { return u.append(int(v)) }
func (u *unfolderArrInt) OnInt16(v int16) error     { return u.append(int(v)) }
func (u *unfolderArrInt) OnInt32(v int32) error     { return u.append(int(v)) }
func (u *unfolderArrInt) OnInt64(v int64) error     { return u.append(int(v)) }
func (u *unfolderArrInt) OnInt(v int) error         { return u.append(int(v)) }
func (u *unfolderArrInt) OnByte(v byte) error       { return u.append(int(v)) }
func (u *unfolderArrInt) OnUint8(v uint8) error     { return u.append(int(v)) }
func (u *unfolderArrInt) OnUint16(v uint16) error   { return u.append(int(v)) }
func (u *unfolderArrInt) OnUint32(v uint32) error   { return u.append(int(v)) }
func (u *unfolderArrInt) OnUint64(v uint64) error   { return u.append(int(v)) }
func (u *unfolderArrInt) OnUint(v uint) error       { return u.append(int(v)) }
func (u *unfolderArrInt) OnFloat32(v float32) error { return u.append(int(v)) }
func (u *unfolderArrInt) OnFloat64(v float64) error { return u.append(int(v)) }

func (u *unfolderArrInt8) OnNil() error              { return u.append(0) }
func (u *unfolderArrInt8) OnInt8(v int8) error       { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnInt16(v int16) error     { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnInt32(v int32) error     { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnInt64(v int64) error     { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnInt(v int) error         { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnByte(v byte) error       { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnUint8(v uint8) error     { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnUint16(v uint16) error   { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnUint32(v uint32) error   { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnUint64(v uint64) error   { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnUint(v uint) error       { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnFloat32(v float32) error { return u.append(int8(v)) }
func (u *unfolderArrInt8) OnFloat64(v float64) error { return u.append(int8(v)) }

func (u *unfolderArrInt16) OnNil() error              { return u.append(0) }
func (u *unfolderArrInt16) OnInt8(v int8) error       { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnInt16(v int16) error     { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnInt32(v int32) error     { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnInt64(v int64) error     { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnInt(v int) error         { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnByte(v byte) error       { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnUint8(v uint8) error     { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnUint16(v uint16) error   { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnUint32(v uint32) error   { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnUint64(v uint64) error   { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnUint(v uint) error       { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnFloat32(v float32) error { return u.append(int16(v)) }
func (u *unfolderArrInt16) OnFloat64(v float64) error { return u.append(int16(v)) }

func (u *unfolderArrInt32) OnNil() error              { return u.append(0) }
func (u *unfolderArrInt32) OnInt8(v int8) error       { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnInt16(v int16) error     { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnInt32(v int32) error     { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnInt64(v int64) error     { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnInt(v int) error         { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnByte(v byte) error       { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnUint8(v uint8) error     { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnUint16(v uint16) error   { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnUint32(v uint32) error   { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnUint64(v uint64) error   { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnUint(v uint) error       { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnFloat32(v float32) error { return u.append(int32(v)) }
func (u *unfolderArrInt32) OnFloat64(v float64) error { return u.append(int32(v)) }

func (u *unfolderArrInt64) OnNil() error              { return u.append(0) }
func (u *unfolderArrInt64) OnInt8(v int8) error       { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnInt16(v int16) error     { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnInt32(v int32) error     { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnInt64(v int64) error     { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnInt(v int) error         { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnByte(v byte) error       { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnUint8(v uint8) error     { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnUint16(v uint16) error   { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnUint32(v uint32) error   { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnUint64(v uint64) error   { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnUint(v uint) error       { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnFloat32(v float32) error { return u.append(int64(v)) }
func (u *unfolderArrInt64) OnFloat64(v float64) error { return u.append(int64(v)) }

func (u *unfolderArrFloat32) OnNil() error              { return u.append(0) }
func (u *unfolderArrFloat32) OnInt8(v int8) error       { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnInt16(v int16) error     { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnInt32(v int32) error     { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnInt64(v int64) error     { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnInt(v int) error         { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnByte(v byte) error       { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnUint8(v uint8) error     { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnUint16(v uint16) error   { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnUint32(v uint32) error   { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnUint64(v uint64) error   { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnUint(v uint) error       { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnFloat32(v float32) error { return u.append(float32(v)) }
func (u *unfolderArrFloat32) OnFloat64(v float64) error { return u.append(float32(v)) }

func (u *unfolderArrFloat64) OnNil() error              { return u.append(0) }
func (u *unfolderArrFloat64) OnInt8(v int8) error       { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnInt16(v int16) error     { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnInt32(v int32) error     { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnInt64(v int64) error     { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnInt(v int) error         { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnByte(v byte) error       { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnUint8(v uint8) error     { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnUint16(v uint16) error   { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnUint32(v uint32) error   { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnUint64(v uint64) error   { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnUint(v uint) error       { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnFloat32(v float32) error { return u.append(float64(v)) }
func (u *unfolderArrFloat64) OnFloat64(v float64) error { return u.append(float64(v)) }
