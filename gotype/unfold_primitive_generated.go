// This file has been generated from 'unfold_primitive.yml', do not edit
package gotype

type unfolderIfc struct {
	reUnfoldEmpty
}

var _singletonUnfolderIfc = &unfolderIfc{}

func newUnfolderIfc() *unfolderIfc {
	return _singletonUnfolderIfc
}

func (u *unfolderIfc) ptr(ctx *unfoldCtx) *interface{} {
	return (*interface{})(ctx.ptr.current)
}

func (u *unfolderIfc) assign(ctx *unfoldCtx, v interface{}) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderBool struct {
	reUnfoldEmpty
}

var _singletonUnfolderBool = &unfolderBool{}

func newUnfolderBool() *unfolderBool {
	return _singletonUnfolderBool
}

func (u *unfolderBool) ptr(ctx *unfoldCtx) *bool {
	return (*bool)(ctx.ptr.current)
}

func (u *unfolderBool) assign(ctx *unfoldCtx, v bool) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderString struct {
	reUnfoldEmpty
}

var _singletonUnfolderString = &unfolderString{}

func newUnfolderString() *unfolderString {
	return _singletonUnfolderString
}

func (u *unfolderString) ptr(ctx *unfoldCtx) *string {
	return (*string)(ctx.ptr.current)
}

func (u *unfolderString) assign(ctx *unfoldCtx, v string) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint struct {
	reUnfoldEmpty
}

var _singletonUnfolderUint = &unfolderUint{}

func newUnfolderUint() *unfolderUint {
	return _singletonUnfolderUint
}

func (u *unfolderUint) ptr(ctx *unfoldCtx) *uint {
	return (*uint)(ctx.ptr.current)
}

func (u *unfolderUint) assign(ctx *unfoldCtx, v uint) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint8 struct {
	reUnfoldEmpty
}

var _singletonUnfolderUint8 = &unfolderUint8{}

func newUnfolderUint8() *unfolderUint8 {
	return _singletonUnfolderUint8
}

func (u *unfolderUint8) ptr(ctx *unfoldCtx) *uint8 {
	return (*uint8)(ctx.ptr.current)
}

func (u *unfolderUint8) assign(ctx *unfoldCtx, v uint8) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint16 struct {
	reUnfoldEmpty
}

var _singletonUnfolderUint16 = &unfolderUint16{}

func newUnfolderUint16() *unfolderUint16 {
	return _singletonUnfolderUint16
}

func (u *unfolderUint16) ptr(ctx *unfoldCtx) *uint16 {
	return (*uint16)(ctx.ptr.current)
}

func (u *unfolderUint16) assign(ctx *unfoldCtx, v uint16) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint32 struct {
	reUnfoldEmpty
}

var _singletonUnfolderUint32 = &unfolderUint32{}

func newUnfolderUint32() *unfolderUint32 {
	return _singletonUnfolderUint32
}

func (u *unfolderUint32) ptr(ctx *unfoldCtx) *uint32 {
	return (*uint32)(ctx.ptr.current)
}

func (u *unfolderUint32) assign(ctx *unfoldCtx, v uint32) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint64 struct {
	reUnfoldEmpty
}

var _singletonUnfolderUint64 = &unfolderUint64{}

func newUnfolderUint64() *unfolderUint64 {
	return _singletonUnfolderUint64
}

func (u *unfolderUint64) ptr(ctx *unfoldCtx) *uint64 {
	return (*uint64)(ctx.ptr.current)
}

func (u *unfolderUint64) assign(ctx *unfoldCtx, v uint64) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt struct {
	reUnfoldEmpty
}

var _singletonUnfolderInt = &unfolderInt{}

func newUnfolderInt() *unfolderInt {
	return _singletonUnfolderInt
}

func (u *unfolderInt) ptr(ctx *unfoldCtx) *int {
	return (*int)(ctx.ptr.current)
}

func (u *unfolderInt) assign(ctx *unfoldCtx, v int) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt8 struct {
	reUnfoldEmpty
}

var _singletonUnfolderInt8 = &unfolderInt8{}

func newUnfolderInt8() *unfolderInt8 {
	return _singletonUnfolderInt8
}

func (u *unfolderInt8) ptr(ctx *unfoldCtx) *int8 {
	return (*int8)(ctx.ptr.current)
}

func (u *unfolderInt8) assign(ctx *unfoldCtx, v int8) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt16 struct {
	reUnfoldEmpty
}

var _singletonUnfolderInt16 = &unfolderInt16{}

func newUnfolderInt16() *unfolderInt16 {
	return _singletonUnfolderInt16
}

func (u *unfolderInt16) ptr(ctx *unfoldCtx) *int16 {
	return (*int16)(ctx.ptr.current)
}

func (u *unfolderInt16) assign(ctx *unfoldCtx, v int16) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt32 struct {
	reUnfoldEmpty
}

var _singletonUnfolderInt32 = &unfolderInt32{}

func newUnfolderInt32() *unfolderInt32 {
	return _singletonUnfolderInt32
}

func (u *unfolderInt32) ptr(ctx *unfoldCtx) *int32 {
	return (*int32)(ctx.ptr.current)
}

func (u *unfolderInt32) assign(ctx *unfoldCtx, v int32) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt64 struct {
	reUnfoldEmpty
}

var _singletonUnfolderInt64 = &unfolderInt64{}

func newUnfolderInt64() *unfolderInt64 {
	return _singletonUnfolderInt64
}

func (u *unfolderInt64) ptr(ctx *unfoldCtx) *int64 {
	return (*int64)(ctx.ptr.current)
}

func (u *unfolderInt64) assign(ctx *unfoldCtx, v int64) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderFloat32 struct {
	reUnfoldEmpty
}

var _singletonUnfolderFloat32 = &unfolderFloat32{}

func newUnfolderFloat32() *unfolderFloat32 {
	return _singletonUnfolderFloat32
}

func (u *unfolderFloat32) ptr(ctx *unfoldCtx) *float32 {
	return (*float32)(ctx.ptr.current)
}

func (u *unfolderFloat32) assign(ctx *unfoldCtx, v float32) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderFloat64 struct {
	reUnfoldEmpty
}

var _singletonUnfolderFloat64 = &unfolderFloat64{}

func newUnfolderFloat64() *unfolderFloat64 {
	return _singletonUnfolderFloat64
}

func (u *unfolderFloat64) ptr(ctx *unfoldCtx) *float64 {
	return (*float64)(ctx.ptr.current)
}

func (u *unfolderFloat64) assign(ctx *unfoldCtx, v float64) error {
	*u.ptr(ctx) = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderIfc) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, nil)
}

func (u *unfolderIfc) OnBool(ctx *unfoldCtx, v bool) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnString(ctx *unfoldCtx, v string) error { return u.assign(ctx, v) }

func (u *unfolderIfc) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, (interface{})(v))
}
func (u *unfolderIfc) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, (interface{})(v))
}

func (u *unfolderBool) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, false)
}

func (u *unfolderBool) OnBool(ctx *unfoldCtx, v bool) error { return u.assign(ctx, v) }

func (u *unfolderString) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, "")
}

func (u *unfolderString) OnString(ctx *unfoldCtx, v string) error { return u.assign(ctx, v) }

func (u *unfolderUint) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderUint) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, uint(v))
}
func (u *unfolderUint) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, uint(v))
}

func (u *unfolderUint8) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderUint8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, uint8(v))
}
func (u *unfolderUint8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, uint8(v))
}

func (u *unfolderUint16) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderUint16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, uint16(v))
}
func (u *unfolderUint16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, uint16(v))
}

func (u *unfolderUint32) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderUint32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, uint32(v))
}
func (u *unfolderUint32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, uint32(v))
}

func (u *unfolderUint64) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderUint64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, uint64(v))
}
func (u *unfolderUint64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, uint64(v))
}

func (u *unfolderInt) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderInt) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, int(v))
}
func (u *unfolderInt) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, int(v))
}

func (u *unfolderInt8) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderInt8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, int8(v))
}
func (u *unfolderInt8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, int8(v))
}

func (u *unfolderInt16) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderInt16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, int16(v))
}
func (u *unfolderInt16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, int16(v))
}

func (u *unfolderInt32) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderInt32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, int32(v))
}
func (u *unfolderInt32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, int32(v))
}

func (u *unfolderInt64) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderInt64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, int64(v))
}
func (u *unfolderInt64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, int64(v))
}

func (u *unfolderFloat32) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderFloat32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, float32(v))
}
func (u *unfolderFloat32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, float32(v))
}

func (u *unfolderFloat64) OnNil(ctx *unfoldCtx) error {
	return u.assign(ctx, 0)
}

func (u *unfolderFloat64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnInt(ctx *unfoldCtx, v int) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.assign(ctx, float64(v))
}
func (u *unfolderFloat64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.assign(ctx, float64(v))
}
