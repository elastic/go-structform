// This file has been generated from 'unfold_arr.yml', do not edit
package gotype

import "github.com/urso/go-structform"

type unfolderArrBool struct {
	reUnfoldEmpty
}

/*
var unfolderArrBoolPool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrBool{}
  },
}
*/

var _singletonUnfolderArrBool = &unfolderArrBool{}

func newUnfolderArrBool() *unfolderArrBool {
	return _singletonUnfolderArrBool
	// u := unfolderArrBoolPool.Get().(*unfolderArrBool)
	// u.to = to
	// return u
	// return &unfolderArrBool{to: to}
}

func (u *unfolderArrBool) free() {
	/*
	  *u = unfolderArrBool{}
	  unfolderArrBoolPool.Put(u)
	*/
}

func (u *unfolderArrBool) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]bool)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]bool, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrBool) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrBool) ptr(ctx *unfoldCtx) *[]bool {
	return (*[]bool)(ctx.ptr.current)
}

func (u *unfolderArrBool) append(ctx *unfoldCtx, v bool) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrString struct {
	reUnfoldEmpty
}

/*
var unfolderArrStringPool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrString{}
  },
}
*/

var _singletonUnfolderArrString = &unfolderArrString{}

func newUnfolderArrString() *unfolderArrString {
	return _singletonUnfolderArrString
	// u := unfolderArrStringPool.Get().(*unfolderArrString)
	// u.to = to
	// return u
	// return &unfolderArrString{to: to}
}

func (u *unfolderArrString) free() {
	/*
	  *u = unfolderArrString{}
	  unfolderArrStringPool.Put(u)
	*/
}

func (u *unfolderArrString) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]string)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]string, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrString) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrString) ptr(ctx *unfoldCtx) *[]string {
	return (*[]string)(ctx.ptr.current)
}

func (u *unfolderArrString) append(ctx *unfoldCtx, v string) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrUint struct {
	reUnfoldEmpty
}

/*
var unfolderArrUintPool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrUint{}
  },
}
*/

var _singletonUnfolderArrUint = &unfolderArrUint{}

func newUnfolderArrUint() *unfolderArrUint {
	return _singletonUnfolderArrUint
	// u := unfolderArrUintPool.Get().(*unfolderArrUint)
	// u.to = to
	// return u
	// return &unfolderArrUint{to: to}
}

func (u *unfolderArrUint) free() {
	/*
	  *u = unfolderArrUint{}
	  unfolderArrUintPool.Put(u)
	*/
}

func (u *unfolderArrUint) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]uint)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]uint, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrUint) ptr(ctx *unfoldCtx) *[]uint {
	return (*[]uint)(ctx.ptr.current)
}

func (u *unfolderArrUint) append(ctx *unfoldCtx, v uint) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrUint8 struct {
	reUnfoldEmpty
}

/*
var unfolderArrUint8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrUint8{}
  },
}
*/

var _singletonUnfolderArrUint8 = &unfolderArrUint8{}

func newUnfolderArrUint8() *unfolderArrUint8 {
	return _singletonUnfolderArrUint8
	// u := unfolderArrUint8Pool.Get().(*unfolderArrUint8)
	// u.to = to
	// return u
	// return &unfolderArrUint8{to: to}
}

func (u *unfolderArrUint8) free() {
	/*
	  *u = unfolderArrUint8{}
	  unfolderArrUint8Pool.Put(u)
	*/
}

func (u *unfolderArrUint8) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]uint8)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]uint8, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint8) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrUint8) ptr(ctx *unfoldCtx) *[]uint8 {
	return (*[]uint8)(ctx.ptr.current)
}

func (u *unfolderArrUint8) append(ctx *unfoldCtx, v uint8) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrUint16 struct {
	reUnfoldEmpty
}

/*
var unfolderArrUint16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrUint16{}
  },
}
*/

var _singletonUnfolderArrUint16 = &unfolderArrUint16{}

func newUnfolderArrUint16() *unfolderArrUint16 {
	return _singletonUnfolderArrUint16
	// u := unfolderArrUint16Pool.Get().(*unfolderArrUint16)
	// u.to = to
	// return u
	// return &unfolderArrUint16{to: to}
}

func (u *unfolderArrUint16) free() {
	/*
	  *u = unfolderArrUint16{}
	  unfolderArrUint16Pool.Put(u)
	*/
}

func (u *unfolderArrUint16) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]uint16)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]uint16, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint16) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrUint16) ptr(ctx *unfoldCtx) *[]uint16 {
	return (*[]uint16)(ctx.ptr.current)
}

func (u *unfolderArrUint16) append(ctx *unfoldCtx, v uint16) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrUint32 struct {
	reUnfoldEmpty
}

/*
var unfolderArrUint32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrUint32{}
  },
}
*/

var _singletonUnfolderArrUint32 = &unfolderArrUint32{}

func newUnfolderArrUint32() *unfolderArrUint32 {
	return _singletonUnfolderArrUint32
	// u := unfolderArrUint32Pool.Get().(*unfolderArrUint32)
	// u.to = to
	// return u
	// return &unfolderArrUint32{to: to}
}

func (u *unfolderArrUint32) free() {
	/*
	  *u = unfolderArrUint32{}
	  unfolderArrUint32Pool.Put(u)
	*/
}

func (u *unfolderArrUint32) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]uint32)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]uint32, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint32) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrUint32) ptr(ctx *unfoldCtx) *[]uint32 {
	return (*[]uint32)(ctx.ptr.current)
}

func (u *unfolderArrUint32) append(ctx *unfoldCtx, v uint32) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrUint64 struct {
	reUnfoldEmpty
}

/*
var unfolderArrUint64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrUint64{}
  },
}
*/

var _singletonUnfolderArrUint64 = &unfolderArrUint64{}

func newUnfolderArrUint64() *unfolderArrUint64 {
	return _singletonUnfolderArrUint64
	// u := unfolderArrUint64Pool.Get().(*unfolderArrUint64)
	// u.to = to
	// return u
	// return &unfolderArrUint64{to: to}
}

func (u *unfolderArrUint64) free() {
	/*
	  *u = unfolderArrUint64{}
	  unfolderArrUint64Pool.Put(u)
	*/
}

func (u *unfolderArrUint64) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]uint64)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]uint64, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrUint64) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrUint64) ptr(ctx *unfoldCtx) *[]uint64 {
	return (*[]uint64)(ctx.ptr.current)
}

func (u *unfolderArrUint64) append(ctx *unfoldCtx, v uint64) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrInt struct {
	reUnfoldEmpty
}

/*
var unfolderArrIntPool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrInt{}
  },
}
*/

var _singletonUnfolderArrInt = &unfolderArrInt{}

func newUnfolderArrInt() *unfolderArrInt {
	return _singletonUnfolderArrInt
	// u := unfolderArrIntPool.Get().(*unfolderArrInt)
	// u.to = to
	// return u
	// return &unfolderArrInt{to: to}
}

func (u *unfolderArrInt) free() {
	/*
	  *u = unfolderArrInt{}
	  unfolderArrIntPool.Put(u)
	*/
}

func (u *unfolderArrInt) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]int)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]int, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrInt) ptr(ctx *unfoldCtx) *[]int {
	return (*[]int)(ctx.ptr.current)
}

func (u *unfolderArrInt) append(ctx *unfoldCtx, v int) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrInt8 struct {
	reUnfoldEmpty
}

/*
var unfolderArrInt8Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrInt8{}
  },
}
*/

var _singletonUnfolderArrInt8 = &unfolderArrInt8{}

func newUnfolderArrInt8() *unfolderArrInt8 {
	return _singletonUnfolderArrInt8
	// u := unfolderArrInt8Pool.Get().(*unfolderArrInt8)
	// u.to = to
	// return u
	// return &unfolderArrInt8{to: to}
}

func (u *unfolderArrInt8) free() {
	/*
	  *u = unfolderArrInt8{}
	  unfolderArrInt8Pool.Put(u)
	*/
}

func (u *unfolderArrInt8) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]int8)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]int8, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt8) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrInt8) ptr(ctx *unfoldCtx) *[]int8 {
	return (*[]int8)(ctx.ptr.current)
}

func (u *unfolderArrInt8) append(ctx *unfoldCtx, v int8) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrInt16 struct {
	reUnfoldEmpty
}

/*
var unfolderArrInt16Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrInt16{}
  },
}
*/

var _singletonUnfolderArrInt16 = &unfolderArrInt16{}

func newUnfolderArrInt16() *unfolderArrInt16 {
	return _singletonUnfolderArrInt16
	// u := unfolderArrInt16Pool.Get().(*unfolderArrInt16)
	// u.to = to
	// return u
	// return &unfolderArrInt16{to: to}
}

func (u *unfolderArrInt16) free() {
	/*
	  *u = unfolderArrInt16{}
	  unfolderArrInt16Pool.Put(u)
	*/
}

func (u *unfolderArrInt16) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]int16)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]int16, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt16) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrInt16) ptr(ctx *unfoldCtx) *[]int16 {
	return (*[]int16)(ctx.ptr.current)
}

func (u *unfolderArrInt16) append(ctx *unfoldCtx, v int16) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrInt32 struct {
	reUnfoldEmpty
}

/*
var unfolderArrInt32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrInt32{}
  },
}
*/

var _singletonUnfolderArrInt32 = &unfolderArrInt32{}

func newUnfolderArrInt32() *unfolderArrInt32 {
	return _singletonUnfolderArrInt32
	// u := unfolderArrInt32Pool.Get().(*unfolderArrInt32)
	// u.to = to
	// return u
	// return &unfolderArrInt32{to: to}
}

func (u *unfolderArrInt32) free() {
	/*
	  *u = unfolderArrInt32{}
	  unfolderArrInt32Pool.Put(u)
	*/
}

func (u *unfolderArrInt32) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]int32)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]int32, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt32) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrInt32) ptr(ctx *unfoldCtx) *[]int32 {
	return (*[]int32)(ctx.ptr.current)
}

func (u *unfolderArrInt32) append(ctx *unfoldCtx, v int32) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrInt64 struct {
	reUnfoldEmpty
}

/*
var unfolderArrInt64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrInt64{}
  },
}
*/

var _singletonUnfolderArrInt64 = &unfolderArrInt64{}

func newUnfolderArrInt64() *unfolderArrInt64 {
	return _singletonUnfolderArrInt64
	// u := unfolderArrInt64Pool.Get().(*unfolderArrInt64)
	// u.to = to
	// return u
	// return &unfolderArrInt64{to: to}
}

func (u *unfolderArrInt64) free() {
	/*
	  *u = unfolderArrInt64{}
	  unfolderArrInt64Pool.Put(u)
	*/
}

func (u *unfolderArrInt64) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]int64)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]int64, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrInt64) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrInt64) ptr(ctx *unfoldCtx) *[]int64 {
	return (*[]int64)(ctx.ptr.current)
}

func (u *unfolderArrInt64) append(ctx *unfoldCtx, v int64) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrFloat32 struct {
	reUnfoldEmpty
}

/*
var unfolderArrFloat32Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrFloat32{}
  },
}
*/

var _singletonUnfolderArrFloat32 = &unfolderArrFloat32{}

func newUnfolderArrFloat32() *unfolderArrFloat32 {
	return _singletonUnfolderArrFloat32
	// u := unfolderArrFloat32Pool.Get().(*unfolderArrFloat32)
	// u.to = to
	// return u
	// return &unfolderArrFloat32{to: to}
}

func (u *unfolderArrFloat32) free() {
	/*
	  *u = unfolderArrFloat32{}
	  unfolderArrFloat32Pool.Put(u)
	*/
}

func (u *unfolderArrFloat32) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]float32)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]float32, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrFloat32) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrFloat32) ptr(ctx *unfoldCtx) *[]float32 {
	return (*[]float32)(ctx.ptr.current)
}

func (u *unfolderArrFloat32) append(ctx *unfoldCtx, v float32) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

type unfolderArrFloat64 struct {
	reUnfoldEmpty
}

/*
var unfolderArrFloat64Pool = sync.Pool{
  New: func() interface{} {
    return &unfolderArrFloat64{}
  },
}
*/

var _singletonUnfolderArrFloat64 = &unfolderArrFloat64{}

func newUnfolderArrFloat64() *unfolderArrFloat64 {
	return _singletonUnfolderArrFloat64
	// u := unfolderArrFloat64Pool.Get().(*unfolderArrFloat64)
	// u.to = to
	// return u
	// return &unfolderArrFloat64{to: to}
}

func (u *unfolderArrFloat64) free() {
	/*
	  *u = unfolderArrFloat64{}
	  unfolderArrFloat64Pool.Put(u)
	*/
}

func (u *unfolderArrFloat64) OnArrayStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedArrayStart
	}

	to := (*[]float64)(ctx.ptr.current)
	if l < 0 {
		l = 0
	}

	if *to == nil && l > 0 {
		*to = make([]float64, l)
	} else if l < len(*to) {
		*to = (*to)[:l]
	}

	dtl.current = unfoldWaitElem
	ctx.idx.push(0)
	return nil
}

func (u *unfolderArrFloat64) OnArrayFinished(ctx *unfoldCtx) error {
	u.free()

	ctx.idx.pop()
	ctx.detail.pop()
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderArrFloat64) ptr(ctx *unfoldCtx) *[]float64 {
	return (*[]float64)(ctx.ptr.current)
}

func (u *unfolderArrFloat64) append(ctx *unfoldCtx, v float64) error {
	idx := &ctx.idx
	to := u.ptr(ctx)
	if len(*to) <= idx.current {
		*to = append(*to, v)
	} else {
		(*to)[idx.current] = v
	}

	idx.current++
	return nil
}

func (u *unfolderArrBool) OnNil(ctx *unfoldCtx) error          { return u.append(ctx, false) }
func (u *unfolderArrBool) OnBool(ctx *unfoldCtx, v bool) error { return u.append(ctx, v) }

func (u *unfolderArrString) OnNil(ctx *unfoldCtx) error              { return u.append(ctx, "") }
func (u *unfolderArrString) OnString(ctx *unfoldCtx, v string) error { return u.append(ctx, v) }

func (u *unfolderArrUint) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrUint) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, uint(v))
}
func (u *unfolderArrUint) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, uint(v))
}

func (u *unfolderArrUint8) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrUint8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, uint8(v))
}
func (u *unfolderArrUint8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, uint8(v))
}

func (u *unfolderArrUint16) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrUint16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, uint16(v))
}
func (u *unfolderArrUint16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, uint16(v))
}

func (u *unfolderArrUint32) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrUint32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, uint32(v))
}
func (u *unfolderArrUint32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, uint32(v))
}

func (u *unfolderArrUint64) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrUint64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, uint64(v))
}
func (u *unfolderArrUint64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, uint64(v))
}

func (u *unfolderArrInt) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrInt) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, int(v))
}
func (u *unfolderArrInt) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, int(v))
}

func (u *unfolderArrInt8) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrInt8) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, int8(v))
}
func (u *unfolderArrInt8) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, int8(v))
}

func (u *unfolderArrInt16) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrInt16) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, int16(v))
}
func (u *unfolderArrInt16) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, int16(v))
}

func (u *unfolderArrInt32) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrInt32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, int32(v))
}
func (u *unfolderArrInt32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, int32(v))
}

func (u *unfolderArrInt64) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrInt64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, int64(v))
}
func (u *unfolderArrInt64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, int64(v))
}

func (u *unfolderArrFloat32) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrFloat32) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, float32(v))
}
func (u *unfolderArrFloat32) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, float32(v))
}

func (u *unfolderArrFloat64) OnNil(ctx *unfoldCtx) error {
	return u.append(ctx, 0)
}
func (u *unfolderArrFloat64) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnInt(ctx *unfoldCtx, v int) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnByte(ctx *unfoldCtx, v byte) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnUint(ctx *unfoldCtx, v uint) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.append(ctx, float64(v))
}
func (u *unfolderArrFloat64) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.append(ctx, float64(v))
}
