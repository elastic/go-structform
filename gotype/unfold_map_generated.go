// This file has been generated from 'unfold_map.yml', do not edit
package gotype

import "github.com/urso/go-structform"

type unfolderMapBool struct {
	reUnfoldEmpty
}

/*
var unfolderMapBoolPool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapBool{}
  },
}
*/

var _singletonUnfolderMapBool = &unfolderMapBool{}

func newUnfolderMapBool() *unfolderMapBool {
	return _singletonUnfolderMapBool
	// u := unfolderMapBoolPool.Get().(*unfolderMapBool)
	// u.to = to
	// return u
	// return &unfolderMapBool{to: to}
}

func (u *unfolderMapBool) free() {
	/*
	  *u = unfolderMapBool{}
	  unfolderMapBoolPool.Put(u)
	*/
}

func (u *unfolderMapBool) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapBool) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapBool) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapBool) put(ctx *unfoldCtx, v bool) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]bool)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]bool{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapString struct {
	reUnfoldEmpty
}

/*
var unfolderMapStringPool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapString{}
  },
}
*/

var _singletonUnfolderMapString = &unfolderMapString{}

func newUnfolderMapString() *unfolderMapString {
	return _singletonUnfolderMapString
	// u := unfolderMapStringPool.Get().(*unfolderMapString)
	// u.to = to
	// return u
	// return &unfolderMapString{to: to}
}

func (u *unfolderMapString) free() {
	/*
	  *u = unfolderMapString{}
	  unfolderMapStringPool.Put(u)
	*/
}

func (u *unfolderMapString) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapString) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapString) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapString) put(ctx *unfoldCtx, v string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]string)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]string{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint struct {
	reUnfoldEmpty
}

/*
var unfolderMapUintPool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapUint{}
  },
}
*/

var _singletonUnfolderMapUint = &unfolderMapUint{}

func newUnfolderMapUint() *unfolderMapUint {
	return _singletonUnfolderMapUint
	// u := unfolderMapUintPool.Get().(*unfolderMapUint)
	// u.to = to
	// return u
	// return &unfolderMapUint{to: to}
}

func (u *unfolderMapUint) free() {
	/*
	  *u = unfolderMapUint{}
	  unfolderMapUintPool.Put(u)
	*/
}

func (u *unfolderMapUint) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapUint) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint) put(ctx *unfoldCtx, v uint) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]uint)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]uint{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint8 struct {
	reUnfoldEmpty
}

/*
var unfolderMapUint8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapUint8{}
  },
}
*/

var _singletonUnfolderMapUint8 = &unfolderMapUint8{}

func newUnfolderMapUint8() *unfolderMapUint8 {
	return _singletonUnfolderMapUint8
	// u := unfolderMapUint8Pool.Get().(*unfolderMapUint8)
	// u.to = to
	// return u
	// return &unfolderMapUint8{to: to}
}

func (u *unfolderMapUint8) free() {
	/*
	  *u = unfolderMapUint8{}
	  unfolderMapUint8Pool.Put(u)
	*/
}

func (u *unfolderMapUint8) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint8) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapUint8) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint8) put(ctx *unfoldCtx, v uint8) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]uint8)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]uint8{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint16 struct {
	reUnfoldEmpty
}

/*
var unfolderMapUint16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapUint16{}
  },
}
*/

var _singletonUnfolderMapUint16 = &unfolderMapUint16{}

func newUnfolderMapUint16() *unfolderMapUint16 {
	return _singletonUnfolderMapUint16
	// u := unfolderMapUint16Pool.Get().(*unfolderMapUint16)
	// u.to = to
	// return u
	// return &unfolderMapUint16{to: to}
}

func (u *unfolderMapUint16) free() {
	/*
	  *u = unfolderMapUint16{}
	  unfolderMapUint16Pool.Put(u)
	*/
}

func (u *unfolderMapUint16) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint16) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapUint16) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint16) put(ctx *unfoldCtx, v uint16) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]uint16)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]uint16{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint32 struct {
	reUnfoldEmpty
}

/*
var unfolderMapUint32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapUint32{}
  },
}
*/

var _singletonUnfolderMapUint32 = &unfolderMapUint32{}

func newUnfolderMapUint32() *unfolderMapUint32 {
	return _singletonUnfolderMapUint32
	// u := unfolderMapUint32Pool.Get().(*unfolderMapUint32)
	// u.to = to
	// return u
	// return &unfolderMapUint32{to: to}
}

func (u *unfolderMapUint32) free() {
	/*
	  *u = unfolderMapUint32{}
	  unfolderMapUint32Pool.Put(u)
	*/
}

func (u *unfolderMapUint32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapUint32) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint32) put(ctx *unfoldCtx, v uint32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]uint32)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]uint32{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint64 struct {
	reUnfoldEmpty
}

/*
var unfolderMapUint64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapUint64{}
  },
}
*/

var _singletonUnfolderMapUint64 = &unfolderMapUint64{}

func newUnfolderMapUint64() *unfolderMapUint64 {
	return _singletonUnfolderMapUint64
	// u := unfolderMapUint64Pool.Get().(*unfolderMapUint64)
	// u.to = to
	// return u
	// return &unfolderMapUint64{to: to}
}

func (u *unfolderMapUint64) free() {
	/*
	  *u = unfolderMapUint64{}
	  unfolderMapUint64Pool.Put(u)
	*/
}

func (u *unfolderMapUint64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapUint64) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint64) put(ctx *unfoldCtx, v uint64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]uint64)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]uint64{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt struct {
	reUnfoldEmpty
}

/*
var unfolderMapIntPool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapInt{}
  },
}
*/

var _singletonUnfolderMapInt = &unfolderMapInt{}

func newUnfolderMapInt() *unfolderMapInt {
	return _singletonUnfolderMapInt
	// u := unfolderMapIntPool.Get().(*unfolderMapInt)
	// u.to = to
	// return u
	// return &unfolderMapInt{to: to}
}

func (u *unfolderMapInt) free() {
	/*
	  *u = unfolderMapInt{}
	  unfolderMapIntPool.Put(u)
	*/
}

func (u *unfolderMapInt) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapInt) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt) put(ctx *unfoldCtx, v int) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]int)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]int{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt8 struct {
	reUnfoldEmpty
}

/*
var unfolderMapInt8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapInt8{}
  },
}
*/

var _singletonUnfolderMapInt8 = &unfolderMapInt8{}

func newUnfolderMapInt8() *unfolderMapInt8 {
	return _singletonUnfolderMapInt8
	// u := unfolderMapInt8Pool.Get().(*unfolderMapInt8)
	// u.to = to
	// return u
	// return &unfolderMapInt8{to: to}
}

func (u *unfolderMapInt8) free() {
	/*
	  *u = unfolderMapInt8{}
	  unfolderMapInt8Pool.Put(u)
	*/
}

func (u *unfolderMapInt8) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt8) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapInt8) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt8) put(ctx *unfoldCtx, v int8) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]int8)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]int8{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt16 struct {
	reUnfoldEmpty
}

/*
var unfolderMapInt16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapInt16{}
  },
}
*/

var _singletonUnfolderMapInt16 = &unfolderMapInt16{}

func newUnfolderMapInt16() *unfolderMapInt16 {
	return _singletonUnfolderMapInt16
	// u := unfolderMapInt16Pool.Get().(*unfolderMapInt16)
	// u.to = to
	// return u
	// return &unfolderMapInt16{to: to}
}

func (u *unfolderMapInt16) free() {
	/*
	  *u = unfolderMapInt16{}
	  unfolderMapInt16Pool.Put(u)
	*/
}

func (u *unfolderMapInt16) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt16) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapInt16) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt16) put(ctx *unfoldCtx, v int16) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]int16)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]int16{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt32 struct {
	reUnfoldEmpty
}

/*
var unfolderMapInt32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapInt32{}
  },
}
*/

var _singletonUnfolderMapInt32 = &unfolderMapInt32{}

func newUnfolderMapInt32() *unfolderMapInt32 {
	return _singletonUnfolderMapInt32
	// u := unfolderMapInt32Pool.Get().(*unfolderMapInt32)
	// u.to = to
	// return u
	// return &unfolderMapInt32{to: to}
}

func (u *unfolderMapInt32) free() {
	/*
	  *u = unfolderMapInt32{}
	  unfolderMapInt32Pool.Put(u)
	*/
}

func (u *unfolderMapInt32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapInt32) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt32) put(ctx *unfoldCtx, v int32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]int32)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]int32{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt64 struct {
	reUnfoldEmpty
}

/*
var unfolderMapInt64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapInt64{}
  },
}
*/

var _singletonUnfolderMapInt64 = &unfolderMapInt64{}

func newUnfolderMapInt64() *unfolderMapInt64 {
	return _singletonUnfolderMapInt64
	// u := unfolderMapInt64Pool.Get().(*unfolderMapInt64)
	// u.to = to
	// return u
	// return &unfolderMapInt64{to: to}
}

func (u *unfolderMapInt64) free() {
	/*
	  *u = unfolderMapInt64{}
	  unfolderMapInt64Pool.Put(u)
	*/
}

func (u *unfolderMapInt64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapInt64) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt64) put(ctx *unfoldCtx, v int64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]int64)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]int64{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapFloat32 struct {
	reUnfoldEmpty
}

/*
var unfolderMapFloat32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapFloat32{}
  },
}
*/

var _singletonUnfolderMapFloat32 = &unfolderMapFloat32{}

func newUnfolderMapFloat32() *unfolderMapFloat32 {
	return _singletonUnfolderMapFloat32
	// u := unfolderMapFloat32Pool.Get().(*unfolderMapFloat32)
	// u.to = to
	// return u
	// return &unfolderMapFloat32{to: to}
}

func (u *unfolderMapFloat32) free() {
	/*
	  *u = unfolderMapFloat32{}
	  unfolderMapFloat32Pool.Put(u)
	*/
}

func (u *unfolderMapFloat32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapFloat32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapFloat32) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapFloat32) put(ctx *unfoldCtx, v float32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]float32)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]float32{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapFloat64 struct {
	reUnfoldEmpty
}

/*
var unfolderMapFloat64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderMapFloat64{}
  },
}
*/

var _singletonUnfolderMapFloat64 = &unfolderMapFloat64{}

func newUnfolderMapFloat64() *unfolderMapFloat64 {
	return _singletonUnfolderMapFloat64
	// u := unfolderMapFloat64Pool.Get().(*unfolderMapFloat64)
	// u.to = to
	// return u
	// return &unfolderMapFloat64{to: to}
}

func (u *unfolderMapFloat64) free() {
	/*
	  *u = unfolderMapFloat64{}
	  unfolderMapFloat64Pool.Put(u)
	*/
}

func (u *unfolderMapFloat64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapFloat64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	u.free()

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapFloat64) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapFloat64) put(ctx *unfoldCtx, v float64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := (*map[string]float64)(ctx.ptr.current)

	if *to == nil {
		*to = map[string]float64{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapBool) OnNil(ctx *unfoldCtx) error          { return u.put(ctx, false) }
func (u *unfolderMapBool) OnBool(ctx *unfoldCtx, v bool) error { return u.put(ctx, v) }

func (u *unfolderMapString) OnNil(ctx *unfoldCtx) error              { return u.put(ctx, "") }
func (u *unfolderMapString) OnString(ctx *unfoldCtx, v string) error { return u.put(ctx, v) }

func (u *unfolderMapUint) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapUint) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, uint(v))
}
func (u *unfolderMapUint) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, uint(v))
}

func (u *unfolderMapUint8) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapUint8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, uint8(v))
}
func (u *unfolderMapUint8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, uint8(v))
}

func (u *unfolderMapUint16) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapUint16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, uint16(v))
}
func (u *unfolderMapUint16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, uint16(v))
}

func (u *unfolderMapUint32) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapUint32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, uint32(v))
}
func (u *unfolderMapUint32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, uint32(v))
}

func (u *unfolderMapUint64) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapUint64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, uint64(v))
}
func (u *unfolderMapUint64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, uint64(v))
}

func (u *unfolderMapInt) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapInt) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, int(v))
}
func (u *unfolderMapInt) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, int(v))
}

func (u *unfolderMapInt8) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapInt8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, int8(v))
}
func (u *unfolderMapInt8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, int8(v))
}

func (u *unfolderMapInt16) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapInt16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, int16(v))
}
func (u *unfolderMapInt16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, int16(v))
}

func (u *unfolderMapInt32) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapInt32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, int32(v))
}
func (u *unfolderMapInt32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, int32(v))
}

func (u *unfolderMapInt64) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapInt64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, int64(v))
}
func (u *unfolderMapInt64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, int64(v))
}

func (u *unfolderMapFloat32) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapFloat32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, float32(v))
}
func (u *unfolderMapFloat32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, float32(v))
}

func (u *unfolderMapFloat64) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, 0)
}
func (u *unfolderMapFloat64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, float64(v))
}
func (u *unfolderMapFloat64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, float64(v))
}
