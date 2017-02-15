// This file has been generated from 'unfold_primitive.yml', do not edit
package gotype

import "sync"

type unfolderBool struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *bool
}

var unfolderBoolPool = sync.Pool{
	New: func() interface{} {
		return &unfolderBool{}
	},
}

func newUnfolderBool(ctx *unfoldCtx, to *bool) *unfolderBool {
	u := unfolderBoolPool.Get().(*unfolderBool)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderBool{ctx: ctx, to: to}
}

func (u *unfolderBool) free() {
	*u = unfolderBool{}
	unfolderBoolPool.Put(u)
}

func (u *unfolderBool) assign(v bool) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderString struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *string
}

var unfolderStringPool = sync.Pool{
	New: func() interface{} {
		return &unfolderString{}
	},
}

func newUnfolderString(ctx *unfoldCtx, to *string) *unfolderString {
	u := unfolderStringPool.Get().(*unfolderString)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderString{ctx: ctx, to: to}
}

func (u *unfolderString) free() {
	*u = unfolderString{}
	unfolderStringPool.Put(u)
}

func (u *unfolderString) assign(v string) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderUint struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *uint
}

var unfolderUintPool = sync.Pool{
	New: func() interface{} {
		return &unfolderUint{}
	},
}

func newUnfolderUint(ctx *unfoldCtx, to *uint) *unfolderUint {
	u := unfolderUintPool.Get().(*unfolderUint)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderUint{ctx: ctx, to: to}
}

func (u *unfolderUint) free() {
	*u = unfolderUint{}
	unfolderUintPool.Put(u)
}

func (u *unfolderUint) assign(v uint) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderUint8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *uint8
}

var unfolderUint8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderUint8{}
	},
}

func newUnfolderUint8(ctx *unfoldCtx, to *uint8) *unfolderUint8 {
	u := unfolderUint8Pool.Get().(*unfolderUint8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderUint8{ctx: ctx, to: to}
}

func (u *unfolderUint8) free() {
	*u = unfolderUint8{}
	unfolderUint8Pool.Put(u)
}

func (u *unfolderUint8) assign(v uint8) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderUint16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *uint16
}

var unfolderUint16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderUint16{}
	},
}

func newUnfolderUint16(ctx *unfoldCtx, to *uint16) *unfolderUint16 {
	u := unfolderUint16Pool.Get().(*unfolderUint16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderUint16{ctx: ctx, to: to}
}

func (u *unfolderUint16) free() {
	*u = unfolderUint16{}
	unfolderUint16Pool.Put(u)
}

func (u *unfolderUint16) assign(v uint16) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderUint32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *uint32
}

var unfolderUint32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderUint32{}
	},
}

func newUnfolderUint32(ctx *unfoldCtx, to *uint32) *unfolderUint32 {
	u := unfolderUint32Pool.Get().(*unfolderUint32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderUint32{ctx: ctx, to: to}
}

func (u *unfolderUint32) free() {
	*u = unfolderUint32{}
	unfolderUint32Pool.Put(u)
}

func (u *unfolderUint32) assign(v uint32) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderUint64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *uint64
}

var unfolderUint64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderUint64{}
	},
}

func newUnfolderUint64(ctx *unfoldCtx, to *uint64) *unfolderUint64 {
	u := unfolderUint64Pool.Get().(*unfolderUint64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderUint64{ctx: ctx, to: to}
}

func (u *unfolderUint64) free() {
	*u = unfolderUint64{}
	unfolderUint64Pool.Put(u)
}

func (u *unfolderUint64) assign(v uint64) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderInt struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *int
}

var unfolderIntPool = sync.Pool{
	New: func() interface{} {
		return &unfolderInt{}
	},
}

func newUnfolderInt(ctx *unfoldCtx, to *int) *unfolderInt {
	u := unfolderIntPool.Get().(*unfolderInt)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderInt{ctx: ctx, to: to}
}

func (u *unfolderInt) free() {
	*u = unfolderInt{}
	unfolderIntPool.Put(u)
}

func (u *unfolderInt) assign(v int) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderInt8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *int8
}

var unfolderInt8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderInt8{}
	},
}

func newUnfolderInt8(ctx *unfoldCtx, to *int8) *unfolderInt8 {
	u := unfolderInt8Pool.Get().(*unfolderInt8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderInt8{ctx: ctx, to: to}
}

func (u *unfolderInt8) free() {
	*u = unfolderInt8{}
	unfolderInt8Pool.Put(u)
}

func (u *unfolderInt8) assign(v int8) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderInt16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *int16
}

var unfolderInt16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderInt16{}
	},
}

func newUnfolderInt16(ctx *unfoldCtx, to *int16) *unfolderInt16 {
	u := unfolderInt16Pool.Get().(*unfolderInt16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderInt16{ctx: ctx, to: to}
}

func (u *unfolderInt16) free() {
	*u = unfolderInt16{}
	unfolderInt16Pool.Put(u)
}

func (u *unfolderInt16) assign(v int16) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderInt32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *int32
}

var unfolderInt32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderInt32{}
	},
}

func newUnfolderInt32(ctx *unfoldCtx, to *int32) *unfolderInt32 {
	u := unfolderInt32Pool.Get().(*unfolderInt32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderInt32{ctx: ctx, to: to}
}

func (u *unfolderInt32) free() {
	*u = unfolderInt32{}
	unfolderInt32Pool.Put(u)
}

func (u *unfolderInt32) assign(v int32) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderInt64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *int64
}

var unfolderInt64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderInt64{}
	},
}

func newUnfolderInt64(ctx *unfoldCtx, to *int64) *unfolderInt64 {
	u := unfolderInt64Pool.Get().(*unfolderInt64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderInt64{ctx: ctx, to: to}
}

func (u *unfolderInt64) free() {
	*u = unfolderInt64{}
	unfolderInt64Pool.Put(u)
}

func (u *unfolderInt64) assign(v int64) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderFloat32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *float32
}

var unfolderFloat32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderFloat32{}
	},
}

func newUnfolderFloat32(ctx *unfoldCtx, to *float32) *unfolderFloat32 {
	u := unfolderFloat32Pool.Get().(*unfolderFloat32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderFloat32{ctx: ctx, to: to}
}

func (u *unfolderFloat32) free() {
	*u = unfolderFloat32{}
	unfolderFloat32Pool.Put(u)
}

func (u *unfolderFloat32) assign(v float32) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

type unfolderFloat64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *float64
}

var unfolderFloat64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderFloat64{}
	},
}

func newUnfolderFloat64(ctx *unfoldCtx, to *float64) *unfolderFloat64 {
	u := unfolderFloat64Pool.Get().(*unfolderFloat64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderFloat64{ctx: ctx, to: to}
}

func (u *unfolderFloat64) free() {
	*u = unfolderFloat64{}
	unfolderFloat64Pool.Put(u)
}

func (u *unfolderFloat64) assign(v float64) error {
	*u.to = v
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderBool) OnNil() error        { return u.assign(false) }
func (u *unfolderBool) OnBool(v bool) error { return u.assign(v) }

func (u *unfolderString) OnNil() error            { return u.assign("") }
func (u *unfolderString) OnString(v string) error { return u.assign(v) }

func (u *unfolderUint) OnNil() error              { return u.assign(0) }
func (u *unfolderUint) OnInt8(v int8) error       { return u.assign(uint(v)) }
func (u *unfolderUint) OnInt16(v int16) error     { return u.assign(uint(v)) }
func (u *unfolderUint) OnInt32(v int32) error     { return u.assign(uint(v)) }
func (u *unfolderUint) OnInt64(v int64) error     { return u.assign(uint(v)) }
func (u *unfolderUint) OnInt(v int) error         { return u.assign(uint(v)) }
func (u *unfolderUint) OnByte(v byte) error       { return u.assign(uint(v)) }
func (u *unfolderUint) OnUint8(v uint8) error     { return u.assign(uint(v)) }
func (u *unfolderUint) OnUint16(v uint16) error   { return u.assign(uint(v)) }
func (u *unfolderUint) OnUint32(v uint32) error   { return u.assign(uint(v)) }
func (u *unfolderUint) OnUint64(v uint64) error   { return u.assign(uint(v)) }
func (u *unfolderUint) OnUint(v uint) error       { return u.assign(uint(v)) }
func (u *unfolderUint) OnFloat32(v float32) error { return u.assign(uint(v)) }
func (u *unfolderUint) OnFloat64(v float64) error { return u.assign(uint(v)) }

func (u *unfolderUint8) OnNil() error              { return u.assign(0) }
func (u *unfolderUint8) OnInt8(v int8) error       { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnInt16(v int16) error     { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnInt32(v int32) error     { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnInt64(v int64) error     { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnInt(v int) error         { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnByte(v byte) error       { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnUint8(v uint8) error     { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnUint16(v uint16) error   { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnUint32(v uint32) error   { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnUint64(v uint64) error   { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnUint(v uint) error       { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnFloat32(v float32) error { return u.assign(uint8(v)) }
func (u *unfolderUint8) OnFloat64(v float64) error { return u.assign(uint8(v)) }

func (u *unfolderUint16) OnNil() error              { return u.assign(0) }
func (u *unfolderUint16) OnInt8(v int8) error       { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnInt16(v int16) error     { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnInt32(v int32) error     { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnInt64(v int64) error     { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnInt(v int) error         { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnByte(v byte) error       { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnUint8(v uint8) error     { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnUint16(v uint16) error   { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnUint32(v uint32) error   { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnUint64(v uint64) error   { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnUint(v uint) error       { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnFloat32(v float32) error { return u.assign(uint16(v)) }
func (u *unfolderUint16) OnFloat64(v float64) error { return u.assign(uint16(v)) }

func (u *unfolderUint32) OnNil() error              { return u.assign(0) }
func (u *unfolderUint32) OnInt8(v int8) error       { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnInt16(v int16) error     { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnInt32(v int32) error     { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnInt64(v int64) error     { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnInt(v int) error         { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnByte(v byte) error       { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnUint8(v uint8) error     { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnUint16(v uint16) error   { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnUint32(v uint32) error   { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnUint64(v uint64) error   { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnUint(v uint) error       { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnFloat32(v float32) error { return u.assign(uint32(v)) }
func (u *unfolderUint32) OnFloat64(v float64) error { return u.assign(uint32(v)) }

func (u *unfolderUint64) OnNil() error              { return u.assign(0) }
func (u *unfolderUint64) OnInt8(v int8) error       { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnInt16(v int16) error     { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnInt32(v int32) error     { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnInt64(v int64) error     { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnInt(v int) error         { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnByte(v byte) error       { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnUint8(v uint8) error     { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnUint16(v uint16) error   { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnUint32(v uint32) error   { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnUint64(v uint64) error   { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnUint(v uint) error       { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnFloat32(v float32) error { return u.assign(uint64(v)) }
func (u *unfolderUint64) OnFloat64(v float64) error { return u.assign(uint64(v)) }

func (u *unfolderInt) OnNil() error              { return u.assign(0) }
func (u *unfolderInt) OnInt8(v int8) error       { return u.assign(int(v)) }
func (u *unfolderInt) OnInt16(v int16) error     { return u.assign(int(v)) }
func (u *unfolderInt) OnInt32(v int32) error     { return u.assign(int(v)) }
func (u *unfolderInt) OnInt64(v int64) error     { return u.assign(int(v)) }
func (u *unfolderInt) OnInt(v int) error         { return u.assign(int(v)) }
func (u *unfolderInt) OnByte(v byte) error       { return u.assign(int(v)) }
func (u *unfolderInt) OnUint8(v uint8) error     { return u.assign(int(v)) }
func (u *unfolderInt) OnUint16(v uint16) error   { return u.assign(int(v)) }
func (u *unfolderInt) OnUint32(v uint32) error   { return u.assign(int(v)) }
func (u *unfolderInt) OnUint64(v uint64) error   { return u.assign(int(v)) }
func (u *unfolderInt) OnUint(v uint) error       { return u.assign(int(v)) }
func (u *unfolderInt) OnFloat32(v float32) error { return u.assign(int(v)) }
func (u *unfolderInt) OnFloat64(v float64) error { return u.assign(int(v)) }

func (u *unfolderInt8) OnNil() error              { return u.assign(0) }
func (u *unfolderInt8) OnInt8(v int8) error       { return u.assign(int8(v)) }
func (u *unfolderInt8) OnInt16(v int16) error     { return u.assign(int8(v)) }
func (u *unfolderInt8) OnInt32(v int32) error     { return u.assign(int8(v)) }
func (u *unfolderInt8) OnInt64(v int64) error     { return u.assign(int8(v)) }
func (u *unfolderInt8) OnInt(v int) error         { return u.assign(int8(v)) }
func (u *unfolderInt8) OnByte(v byte) error       { return u.assign(int8(v)) }
func (u *unfolderInt8) OnUint8(v uint8) error     { return u.assign(int8(v)) }
func (u *unfolderInt8) OnUint16(v uint16) error   { return u.assign(int8(v)) }
func (u *unfolderInt8) OnUint32(v uint32) error   { return u.assign(int8(v)) }
func (u *unfolderInt8) OnUint64(v uint64) error   { return u.assign(int8(v)) }
func (u *unfolderInt8) OnUint(v uint) error       { return u.assign(int8(v)) }
func (u *unfolderInt8) OnFloat32(v float32) error { return u.assign(int8(v)) }
func (u *unfolderInt8) OnFloat64(v float64) error { return u.assign(int8(v)) }

func (u *unfolderInt16) OnNil() error              { return u.assign(0) }
func (u *unfolderInt16) OnInt8(v int8) error       { return u.assign(int16(v)) }
func (u *unfolderInt16) OnInt16(v int16) error     { return u.assign(int16(v)) }
func (u *unfolderInt16) OnInt32(v int32) error     { return u.assign(int16(v)) }
func (u *unfolderInt16) OnInt64(v int64) error     { return u.assign(int16(v)) }
func (u *unfolderInt16) OnInt(v int) error         { return u.assign(int16(v)) }
func (u *unfolderInt16) OnByte(v byte) error       { return u.assign(int16(v)) }
func (u *unfolderInt16) OnUint8(v uint8) error     { return u.assign(int16(v)) }
func (u *unfolderInt16) OnUint16(v uint16) error   { return u.assign(int16(v)) }
func (u *unfolderInt16) OnUint32(v uint32) error   { return u.assign(int16(v)) }
func (u *unfolderInt16) OnUint64(v uint64) error   { return u.assign(int16(v)) }
func (u *unfolderInt16) OnUint(v uint) error       { return u.assign(int16(v)) }
func (u *unfolderInt16) OnFloat32(v float32) error { return u.assign(int16(v)) }
func (u *unfolderInt16) OnFloat64(v float64) error { return u.assign(int16(v)) }

func (u *unfolderInt32) OnNil() error              { return u.assign(0) }
func (u *unfolderInt32) OnInt8(v int8) error       { return u.assign(int32(v)) }
func (u *unfolderInt32) OnInt16(v int16) error     { return u.assign(int32(v)) }
func (u *unfolderInt32) OnInt32(v int32) error     { return u.assign(int32(v)) }
func (u *unfolderInt32) OnInt64(v int64) error     { return u.assign(int32(v)) }
func (u *unfolderInt32) OnInt(v int) error         { return u.assign(int32(v)) }
func (u *unfolderInt32) OnByte(v byte) error       { return u.assign(int32(v)) }
func (u *unfolderInt32) OnUint8(v uint8) error     { return u.assign(int32(v)) }
func (u *unfolderInt32) OnUint16(v uint16) error   { return u.assign(int32(v)) }
func (u *unfolderInt32) OnUint32(v uint32) error   { return u.assign(int32(v)) }
func (u *unfolderInt32) OnUint64(v uint64) error   { return u.assign(int32(v)) }
func (u *unfolderInt32) OnUint(v uint) error       { return u.assign(int32(v)) }
func (u *unfolderInt32) OnFloat32(v float32) error { return u.assign(int32(v)) }
func (u *unfolderInt32) OnFloat64(v float64) error { return u.assign(int32(v)) }

func (u *unfolderInt64) OnNil() error              { return u.assign(0) }
func (u *unfolderInt64) OnInt8(v int8) error       { return u.assign(int64(v)) }
func (u *unfolderInt64) OnInt16(v int16) error     { return u.assign(int64(v)) }
func (u *unfolderInt64) OnInt32(v int32) error     { return u.assign(int64(v)) }
func (u *unfolderInt64) OnInt64(v int64) error     { return u.assign(int64(v)) }
func (u *unfolderInt64) OnInt(v int) error         { return u.assign(int64(v)) }
func (u *unfolderInt64) OnByte(v byte) error       { return u.assign(int64(v)) }
func (u *unfolderInt64) OnUint8(v uint8) error     { return u.assign(int64(v)) }
func (u *unfolderInt64) OnUint16(v uint16) error   { return u.assign(int64(v)) }
func (u *unfolderInt64) OnUint32(v uint32) error   { return u.assign(int64(v)) }
func (u *unfolderInt64) OnUint64(v uint64) error   { return u.assign(int64(v)) }
func (u *unfolderInt64) OnUint(v uint) error       { return u.assign(int64(v)) }
func (u *unfolderInt64) OnFloat32(v float32) error { return u.assign(int64(v)) }
func (u *unfolderInt64) OnFloat64(v float64) error { return u.assign(int64(v)) }

func (u *unfolderFloat32) OnNil() error              { return u.assign(0) }
func (u *unfolderFloat32) OnInt8(v int8) error       { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnInt16(v int16) error     { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnInt32(v int32) error     { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnInt64(v int64) error     { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnInt(v int) error         { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnByte(v byte) error       { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnUint8(v uint8) error     { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnUint16(v uint16) error   { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnUint32(v uint32) error   { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnUint64(v uint64) error   { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnUint(v uint) error       { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnFloat32(v float32) error { return u.assign(float32(v)) }
func (u *unfolderFloat32) OnFloat64(v float64) error { return u.assign(float32(v)) }

func (u *unfolderFloat64) OnNil() error              { return u.assign(0) }
func (u *unfolderFloat64) OnInt8(v int8) error       { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnInt16(v int16) error     { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnInt32(v int32) error     { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnInt64(v int64) error     { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnInt(v int) error         { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnByte(v byte) error       { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnUint8(v uint8) error     { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnUint16(v uint16) error   { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnUint32(v uint32) error   { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnUint64(v uint64) error   { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnUint(v uint) error       { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnFloat32(v float32) error { return u.assign(float64(v)) }
func (u *unfolderFloat64) OnFloat64(v float64) error { return u.assign(float64(v)) }
