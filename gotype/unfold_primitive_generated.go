// This file has been generated from 'unfold_primitive.yml', do not edit
package gotype

type unfolderBool struct {
	reUnfoldEmpty
}

/*
var unfolderBoolPool = sync.Pool{
  New: func() interface{} {
    return &unfolderBool{}
  },
}
*/

var _singletonUnfolderBool = &unfolderBool{}

func newUnfolderBool() *unfolderBool {
	return _singletonUnfolderBool
	// u := unfolderBoolPool.Get().(*unfolderBool)
	// u.to = to
	// return u
	// return &unfolderBool{to: to}
}

func (u *unfolderBool) free() {
	/*
	  *u = unfolderBool{}
	  unfolderBoolPool.Put(u)
	*/
}

func (u *unfolderBool) assign(ctx *unfoldCtx, v bool) error {
	// *u.to = v
	u.free()

	to := (*bool)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderString struct {
	reUnfoldEmpty
}

/*
var unfolderStringPool = sync.Pool{
  New: func() interface{} {
    return &unfolderString{}
  },
}
*/

var _singletonUnfolderString = &unfolderString{}

func newUnfolderString() *unfolderString {
	return _singletonUnfolderString
	// u := unfolderStringPool.Get().(*unfolderString)
	// u.to = to
	// return u
	// return &unfolderString{to: to}
}

func (u *unfolderString) free() {
	/*
	  *u = unfolderString{}
	  unfolderStringPool.Put(u)
	*/
}

func (u *unfolderString) assign(ctx *unfoldCtx, v string) error {
	// *u.to = v
	u.free()

	to := (*string)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint struct {
	reUnfoldEmpty
}

/*
var unfolderUintPool = sync.Pool{
  New: func() interface{} {
    return &unfolderUint{}
  },
}
*/

var _singletonUnfolderUint = &unfolderUint{}

func newUnfolderUint() *unfolderUint {
	return _singletonUnfolderUint
	// u := unfolderUintPool.Get().(*unfolderUint)
	// u.to = to
	// return u
	// return &unfolderUint{to: to}
}

func (u *unfolderUint) free() {
	/*
	  *u = unfolderUint{}
	  unfolderUintPool.Put(u)
	*/
}

func (u *unfolderUint) assign(ctx *unfoldCtx, v uint) error {
	// *u.to = v
	u.free()

	to := (*uint)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint8 struct {
	reUnfoldEmpty
}

/*
var unfolderUint8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderUint8{}
  },
}
*/

var _singletonUnfolderUint8 = &unfolderUint8{}

func newUnfolderUint8() *unfolderUint8 {
	return _singletonUnfolderUint8
	// u := unfolderUint8Pool.Get().(*unfolderUint8)
	// u.to = to
	// return u
	// return &unfolderUint8{to: to}
}

func (u *unfolderUint8) free() {
	/*
	  *u = unfolderUint8{}
	  unfolderUint8Pool.Put(u)
	*/
}

func (u *unfolderUint8) assign(ctx *unfoldCtx, v uint8) error {
	// *u.to = v
	u.free()

	to := (*uint8)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint16 struct {
	reUnfoldEmpty
}

/*
var unfolderUint16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderUint16{}
  },
}
*/

var _singletonUnfolderUint16 = &unfolderUint16{}

func newUnfolderUint16() *unfolderUint16 {
	return _singletonUnfolderUint16
	// u := unfolderUint16Pool.Get().(*unfolderUint16)
	// u.to = to
	// return u
	// return &unfolderUint16{to: to}
}

func (u *unfolderUint16) free() {
	/*
	  *u = unfolderUint16{}
	  unfolderUint16Pool.Put(u)
	*/
}

func (u *unfolderUint16) assign(ctx *unfoldCtx, v uint16) error {
	// *u.to = v
	u.free()

	to := (*uint16)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint32 struct {
	reUnfoldEmpty
}

/*
var unfolderUint32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderUint32{}
  },
}
*/

var _singletonUnfolderUint32 = &unfolderUint32{}

func newUnfolderUint32() *unfolderUint32 {
	return _singletonUnfolderUint32
	// u := unfolderUint32Pool.Get().(*unfolderUint32)
	// u.to = to
	// return u
	// return &unfolderUint32{to: to}
}

func (u *unfolderUint32) free() {
	/*
	  *u = unfolderUint32{}
	  unfolderUint32Pool.Put(u)
	*/
}

func (u *unfolderUint32) assign(ctx *unfoldCtx, v uint32) error {
	// *u.to = v
	u.free()

	to := (*uint32)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderUint64 struct {
	reUnfoldEmpty
}

/*
var unfolderUint64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderUint64{}
  },
}
*/

var _singletonUnfolderUint64 = &unfolderUint64{}

func newUnfolderUint64() *unfolderUint64 {
	return _singletonUnfolderUint64
	// u := unfolderUint64Pool.Get().(*unfolderUint64)
	// u.to = to
	// return u
	// return &unfolderUint64{to: to}
}

func (u *unfolderUint64) free() {
	/*
	  *u = unfolderUint64{}
	  unfolderUint64Pool.Put(u)
	*/
}

func (u *unfolderUint64) assign(ctx *unfoldCtx, v uint64) error {
	// *u.to = v
	u.free()

	to := (*uint64)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt struct {
	reUnfoldEmpty
}

/*
var unfolderIntPool = sync.Pool{
  New: func() interface{} {
    return &unfolderInt{}
  },
}
*/

var _singletonUnfolderInt = &unfolderInt{}

func newUnfolderInt() *unfolderInt {
	return _singletonUnfolderInt
	// u := unfolderIntPool.Get().(*unfolderInt)
	// u.to = to
	// return u
	// return &unfolderInt{to: to}
}

func (u *unfolderInt) free() {
	/*
	  *u = unfolderInt{}
	  unfolderIntPool.Put(u)
	*/
}

func (u *unfolderInt) assign(ctx *unfoldCtx, v int) error {
	// *u.to = v
	u.free()

	to := (*int)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt8 struct {
	reUnfoldEmpty
}

/*
var unfolderInt8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderInt8{}
  },
}
*/

var _singletonUnfolderInt8 = &unfolderInt8{}

func newUnfolderInt8() *unfolderInt8 {
	return _singletonUnfolderInt8
	// u := unfolderInt8Pool.Get().(*unfolderInt8)
	// u.to = to
	// return u
	// return &unfolderInt8{to: to}
}

func (u *unfolderInt8) free() {
	/*
	  *u = unfolderInt8{}
	  unfolderInt8Pool.Put(u)
	*/
}

func (u *unfolderInt8) assign(ctx *unfoldCtx, v int8) error {
	// *u.to = v
	u.free()

	to := (*int8)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt16 struct {
	reUnfoldEmpty
}

/*
var unfolderInt16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderInt16{}
  },
}
*/

var _singletonUnfolderInt16 = &unfolderInt16{}

func newUnfolderInt16() *unfolderInt16 {
	return _singletonUnfolderInt16
	// u := unfolderInt16Pool.Get().(*unfolderInt16)
	// u.to = to
	// return u
	// return &unfolderInt16{to: to}
}

func (u *unfolderInt16) free() {
	/*
	  *u = unfolderInt16{}
	  unfolderInt16Pool.Put(u)
	*/
}

func (u *unfolderInt16) assign(ctx *unfoldCtx, v int16) error {
	// *u.to = v
	u.free()

	to := (*int16)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt32 struct {
	reUnfoldEmpty
}

/*
var unfolderInt32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderInt32{}
  },
}
*/

var _singletonUnfolderInt32 = &unfolderInt32{}

func newUnfolderInt32() *unfolderInt32 {
	return _singletonUnfolderInt32
	// u := unfolderInt32Pool.Get().(*unfolderInt32)
	// u.to = to
	// return u
	// return &unfolderInt32{to: to}
}

func (u *unfolderInt32) free() {
	/*
	  *u = unfolderInt32{}
	  unfolderInt32Pool.Put(u)
	*/
}

func (u *unfolderInt32) assign(ctx *unfoldCtx, v int32) error {
	// *u.to = v
	u.free()

	to := (*int32)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderInt64 struct {
	reUnfoldEmpty
}

/*
var unfolderInt64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderInt64{}
  },
}
*/

var _singletonUnfolderInt64 = &unfolderInt64{}

func newUnfolderInt64() *unfolderInt64 {
	return _singletonUnfolderInt64
	// u := unfolderInt64Pool.Get().(*unfolderInt64)
	// u.to = to
	// return u
	// return &unfolderInt64{to: to}
}

func (u *unfolderInt64) free() {
	/*
	  *u = unfolderInt64{}
	  unfolderInt64Pool.Put(u)
	*/
}

func (u *unfolderInt64) assign(ctx *unfoldCtx, v int64) error {
	// *u.to = v
	u.free()

	to := (*int64)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderFloat32 struct {
	reUnfoldEmpty
}

/*
var unfolderFloat32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderFloat32{}
  },
}
*/

var _singletonUnfolderFloat32 = &unfolderFloat32{}

func newUnfolderFloat32() *unfolderFloat32 {
	return _singletonUnfolderFloat32
	// u := unfolderFloat32Pool.Get().(*unfolderFloat32)
	// u.to = to
	// return u
	// return &unfolderFloat32{to: to}
}

func (u *unfolderFloat32) free() {
	/*
	  *u = unfolderFloat32{}
	  unfolderFloat32Pool.Put(u)
	*/
}

func (u *unfolderFloat32) assign(ctx *unfoldCtx, v float32) error {
	// *u.to = v
	u.free()

	to := (*float32)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

type unfolderFloat64 struct {
	reUnfoldEmpty
}

/*
var unfolderFloat64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderFloat64{}
  },
}
*/

var _singletonUnfolderFloat64 = &unfolderFloat64{}

func newUnfolderFloat64() *unfolderFloat64 {
	return _singletonUnfolderFloat64
	// u := unfolderFloat64Pool.Get().(*unfolderFloat64)
	// u.to = to
	// return u
	// return &unfolderFloat64{to: to}
}

func (u *unfolderFloat64) free() {
	/*
	  *u = unfolderFloat64{}
	  unfolderFloat64Pool.Put(u)
	*/
}

func (u *unfolderFloat64) assign(ctx *unfoldCtx, v float64) error {
	// *u.to = v
	u.free()

	to := (*float64)(ctx.ptr.current)
	*to = v

	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderBool) OnNil(ctx *unfoldCtx) error          { return u.assign(ctx, false) }
func (u *unfolderBool) OnBool(ctx *unfoldCtx, v bool) error { return u.assign(ctx, v) }

func (u *unfolderString) OnNil(ctx *unfoldCtx) error              { return u.assign(ctx, "") }
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
