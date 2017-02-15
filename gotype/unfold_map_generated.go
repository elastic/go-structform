// This file has been generated from 'unfold_map.yml', do not edit
package gotype

import (
	"sync"

	"github.com/urso/go-structform"
)

type unfolderMapBool struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]bool
}

var unfolderMapBoolPool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapBool{}
	},
}

func newUnfolderMapBool(ctx *unfoldCtx, to *map[string]bool) *unfolderMapBool {
	u := unfolderMapBoolPool.Get().(*unfolderMapBool)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapBool{ctx: ctx, to: to}
}

func (u *unfolderMapBool) free() {
	*u = unfolderMapBool{}
	unfolderMapBoolPool.Put(u)
}

func (u *unfolderMapBool) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapBool) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapBool) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapBool) put(v bool) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]bool{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapString struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]string
}

var unfolderMapStringPool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapString{}
	},
}

func newUnfolderMapString(ctx *unfoldCtx, to *map[string]string) *unfolderMapString {
	u := unfolderMapStringPool.Get().(*unfolderMapString)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapString{ctx: ctx, to: to}
}

func (u *unfolderMapString) free() {
	*u = unfolderMapString{}
	unfolderMapStringPool.Put(u)
}

func (u *unfolderMapString) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapString) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapString) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapString) put(v string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]string{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]uint
}

var unfolderMapUintPool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapUint{}
	},
}

func newUnfolderMapUint(ctx *unfoldCtx, to *map[string]uint) *unfolderMapUint {
	u := unfolderMapUintPool.Get().(*unfolderMapUint)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapUint{ctx: ctx, to: to}
}

func (u *unfolderMapUint) free() {
	*u = unfolderMapUint{}
	unfolderMapUintPool.Put(u)
}

func (u *unfolderMapUint) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapUint) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint) put(v uint) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]uint{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]uint8
}

var unfolderMapUint8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapUint8{}
	},
}

func newUnfolderMapUint8(ctx *unfoldCtx, to *map[string]uint8) *unfolderMapUint8 {
	u := unfolderMapUint8Pool.Get().(*unfolderMapUint8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapUint8{ctx: ctx, to: to}
}

func (u *unfolderMapUint8) free() {
	*u = unfolderMapUint8{}
	unfolderMapUint8Pool.Put(u)
}

func (u *unfolderMapUint8) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint8) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapUint8) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint8) put(v uint8) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]uint8{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]uint16
}

var unfolderMapUint16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapUint16{}
	},
}

func newUnfolderMapUint16(ctx *unfoldCtx, to *map[string]uint16) *unfolderMapUint16 {
	u := unfolderMapUint16Pool.Get().(*unfolderMapUint16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapUint16{ctx: ctx, to: to}
}

func (u *unfolderMapUint16) free() {
	*u = unfolderMapUint16{}
	unfolderMapUint16Pool.Put(u)
}

func (u *unfolderMapUint16) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint16) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapUint16) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint16) put(v uint16) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]uint16{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]uint32
}

var unfolderMapUint32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapUint32{}
	},
}

func newUnfolderMapUint32(ctx *unfoldCtx, to *map[string]uint32) *unfolderMapUint32 {
	u := unfolderMapUint32Pool.Get().(*unfolderMapUint32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapUint32{ctx: ctx, to: to}
}

func (u *unfolderMapUint32) free() {
	*u = unfolderMapUint32{}
	unfolderMapUint32Pool.Put(u)
}

func (u *unfolderMapUint32) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint32) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapUint32) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint32) put(v uint32) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]uint32{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapUint64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]uint64
}

var unfolderMapUint64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapUint64{}
	},
}

func newUnfolderMapUint64(ctx *unfoldCtx, to *map[string]uint64) *unfolderMapUint64 {
	u := unfolderMapUint64Pool.Get().(*unfolderMapUint64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapUint64{ctx: ctx, to: to}
}

func (u *unfolderMapUint64) free() {
	*u = unfolderMapUint64{}
	unfolderMapUint64Pool.Put(u)
}

func (u *unfolderMapUint64) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapUint64) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapUint64) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapUint64) put(v uint64) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]uint64{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]int
}

var unfolderMapIntPool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapInt{}
	},
}

func newUnfolderMapInt(ctx *unfoldCtx, to *map[string]int) *unfolderMapInt {
	u := unfolderMapIntPool.Get().(*unfolderMapInt)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapInt{ctx: ctx, to: to}
}

func (u *unfolderMapInt) free() {
	*u = unfolderMapInt{}
	unfolderMapIntPool.Put(u)
}

func (u *unfolderMapInt) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapInt) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt) put(v int) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]int{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt8 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]int8
}

var unfolderMapInt8Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapInt8{}
	},
}

func newUnfolderMapInt8(ctx *unfoldCtx, to *map[string]int8) *unfolderMapInt8 {
	u := unfolderMapInt8Pool.Get().(*unfolderMapInt8)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapInt8{ctx: ctx, to: to}
}

func (u *unfolderMapInt8) free() {
	*u = unfolderMapInt8{}
	unfolderMapInt8Pool.Put(u)
}

func (u *unfolderMapInt8) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt8) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapInt8) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt8) put(v int8) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]int8{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt16 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]int16
}

var unfolderMapInt16Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapInt16{}
	},
}

func newUnfolderMapInt16(ctx *unfoldCtx, to *map[string]int16) *unfolderMapInt16 {
	u := unfolderMapInt16Pool.Get().(*unfolderMapInt16)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapInt16{ctx: ctx, to: to}
}

func (u *unfolderMapInt16) free() {
	*u = unfolderMapInt16{}
	unfolderMapInt16Pool.Put(u)
}

func (u *unfolderMapInt16) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt16) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapInt16) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt16) put(v int16) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]int16{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]int32
}

var unfolderMapInt32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapInt32{}
	},
}

func newUnfolderMapInt32(ctx *unfoldCtx, to *map[string]int32) *unfolderMapInt32 {
	u := unfolderMapInt32Pool.Get().(*unfolderMapInt32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapInt32{ctx: ctx, to: to}
}

func (u *unfolderMapInt32) free() {
	*u = unfolderMapInt32{}
	unfolderMapInt32Pool.Put(u)
}

func (u *unfolderMapInt32) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt32) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapInt32) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt32) put(v int32) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]int32{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapInt64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]int64
}

var unfolderMapInt64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapInt64{}
	},
}

func newUnfolderMapInt64(ctx *unfoldCtx, to *map[string]int64) *unfolderMapInt64 {
	u := unfolderMapInt64Pool.Get().(*unfolderMapInt64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapInt64{ctx: ctx, to: to}
}

func (u *unfolderMapInt64) free() {
	*u = unfolderMapInt64{}
	unfolderMapInt64Pool.Put(u)
}

func (u *unfolderMapInt64) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapInt64) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapInt64) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapInt64) put(v int64) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]int64{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapFloat32 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]float32
}

var unfolderMapFloat32Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapFloat32{}
	},
}

func newUnfolderMapFloat32(ctx *unfoldCtx, to *map[string]float32) *unfolderMapFloat32 {
	u := unfolderMapFloat32Pool.Get().(*unfolderMapFloat32)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapFloat32{ctx: ctx, to: to}
}

func (u *unfolderMapFloat32) free() {
	*u = unfolderMapFloat32{}
	unfolderMapFloat32Pool.Put(u)
}

func (u *unfolderMapFloat32) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapFloat32) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapFloat32) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapFloat32) put(v float32) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]float32{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapFloat64 struct {
	reUnfoldEmpty
	ctx *unfoldCtx
	to  *map[string]float64
}

var unfolderMapFloat64Pool = sync.Pool{
	New: func() interface{} {
		return &unfolderMapFloat64{}
	},
}

func newUnfolderMapFloat64(ctx *unfoldCtx, to *map[string]float64) *unfolderMapFloat64 {
	u := unfolderMapFloat64Pool.Get().(*unfolderMapFloat64)
	u.ctx = ctx
	u.to = to
	return u
	// return &unfolderMapFloat64{ctx: ctx, to: to}
}

func (u *unfolderMapFloat64) free() {
	*u = unfolderMapFloat64{}
	unfolderMapFloat64Pool.Put(u)
}

func (u *unfolderMapFloat64) OnObjectStart(l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitStart {
		return errUnexpectedObjectStart
	}

	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapFloat64) OnObjectFinished() error {
	dtl := &u.ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	u.ctx.unfolder.pop()
	u.free()
	return nil
}

func (u *unfolderMapFloat64) OnKey(key string) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	u.ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapFloat64) put(v float64) error {
	dtl := &u.ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	if *u.to == nil {
		*u.to = map[string]float64{}
	}
	(*u.to)[u.ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapBool) OnNil() error        { return u.put(false) }
func (u *unfolderMapBool) OnBool(v bool) error { return u.put(v) }

func (u *unfolderMapString) OnNil() error            { return u.put("") }
func (u *unfolderMapString) OnString(v string) error { return u.put(v) }

func (u *unfolderMapUint) OnNil() error              { return u.put(0) }
func (u *unfolderMapUint) OnInt8(v int8) error       { return u.put(uint(v)) }
func (u *unfolderMapUint) OnInt16(v int16) error     { return u.put(uint(v)) }
func (u *unfolderMapUint) OnInt32(v int32) error     { return u.put(uint(v)) }
func (u *unfolderMapUint) OnInt64(v int64) error     { return u.put(uint(v)) }
func (u *unfolderMapUint) OnInt(v int) error         { return u.put(uint(v)) }
func (u *unfolderMapUint) OnByte(v byte) error       { return u.put(uint(v)) }
func (u *unfolderMapUint) OnUint8(v uint8) error     { return u.put(uint(v)) }
func (u *unfolderMapUint) OnUint16(v uint16) error   { return u.put(uint(v)) }
func (u *unfolderMapUint) OnUint32(v uint32) error   { return u.put(uint(v)) }
func (u *unfolderMapUint) OnUint64(v uint64) error   { return u.put(uint(v)) }
func (u *unfolderMapUint) OnUint(v uint) error       { return u.put(uint(v)) }
func (u *unfolderMapUint) OnFloat32(v float32) error { return u.put(uint(v)) }
func (u *unfolderMapUint) OnFloat64(v float64) error { return u.put(uint(v)) }

func (u *unfolderMapUint8) OnNil() error              { return u.put(0) }
func (u *unfolderMapUint8) OnInt8(v int8) error       { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnInt16(v int16) error     { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnInt32(v int32) error     { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnInt64(v int64) error     { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnInt(v int) error         { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnByte(v byte) error       { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnUint8(v uint8) error     { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnUint16(v uint16) error   { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnUint32(v uint32) error   { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnUint64(v uint64) error   { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnUint(v uint) error       { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnFloat32(v float32) error { return u.put(uint8(v)) }
func (u *unfolderMapUint8) OnFloat64(v float64) error { return u.put(uint8(v)) }

func (u *unfolderMapUint16) OnNil() error              { return u.put(0) }
func (u *unfolderMapUint16) OnInt8(v int8) error       { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnInt16(v int16) error     { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnInt32(v int32) error     { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnInt64(v int64) error     { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnInt(v int) error         { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnByte(v byte) error       { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnUint8(v uint8) error     { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnUint16(v uint16) error   { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnUint32(v uint32) error   { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnUint64(v uint64) error   { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnUint(v uint) error       { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnFloat32(v float32) error { return u.put(uint16(v)) }
func (u *unfolderMapUint16) OnFloat64(v float64) error { return u.put(uint16(v)) }

func (u *unfolderMapUint32) OnNil() error              { return u.put(0) }
func (u *unfolderMapUint32) OnInt8(v int8) error       { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnInt16(v int16) error     { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnInt32(v int32) error     { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnInt64(v int64) error     { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnInt(v int) error         { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnByte(v byte) error       { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnUint8(v uint8) error     { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnUint16(v uint16) error   { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnUint32(v uint32) error   { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnUint64(v uint64) error   { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnUint(v uint) error       { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnFloat32(v float32) error { return u.put(uint32(v)) }
func (u *unfolderMapUint32) OnFloat64(v float64) error { return u.put(uint32(v)) }

func (u *unfolderMapUint64) OnNil() error              { return u.put(0) }
func (u *unfolderMapUint64) OnInt8(v int8) error       { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnInt16(v int16) error     { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnInt32(v int32) error     { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnInt64(v int64) error     { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnInt(v int) error         { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnByte(v byte) error       { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnUint8(v uint8) error     { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnUint16(v uint16) error   { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnUint32(v uint32) error   { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnUint64(v uint64) error   { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnUint(v uint) error       { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnFloat32(v float32) error { return u.put(uint64(v)) }
func (u *unfolderMapUint64) OnFloat64(v float64) error { return u.put(uint64(v)) }

func (u *unfolderMapInt) OnNil() error              { return u.put(0) }
func (u *unfolderMapInt) OnInt8(v int8) error       { return u.put(int(v)) }
func (u *unfolderMapInt) OnInt16(v int16) error     { return u.put(int(v)) }
func (u *unfolderMapInt) OnInt32(v int32) error     { return u.put(int(v)) }
func (u *unfolderMapInt) OnInt64(v int64) error     { return u.put(int(v)) }
func (u *unfolderMapInt) OnInt(v int) error         { return u.put(int(v)) }
func (u *unfolderMapInt) OnByte(v byte) error       { return u.put(int(v)) }
func (u *unfolderMapInt) OnUint8(v uint8) error     { return u.put(int(v)) }
func (u *unfolderMapInt) OnUint16(v uint16) error   { return u.put(int(v)) }
func (u *unfolderMapInt) OnUint32(v uint32) error   { return u.put(int(v)) }
func (u *unfolderMapInt) OnUint64(v uint64) error   { return u.put(int(v)) }
func (u *unfolderMapInt) OnUint(v uint) error       { return u.put(int(v)) }
func (u *unfolderMapInt) OnFloat32(v float32) error { return u.put(int(v)) }
func (u *unfolderMapInt) OnFloat64(v float64) error { return u.put(int(v)) }

func (u *unfolderMapInt8) OnNil() error              { return u.put(0) }
func (u *unfolderMapInt8) OnInt8(v int8) error       { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnInt16(v int16) error     { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnInt32(v int32) error     { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnInt64(v int64) error     { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnInt(v int) error         { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnByte(v byte) error       { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnUint8(v uint8) error     { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnUint16(v uint16) error   { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnUint32(v uint32) error   { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnUint64(v uint64) error   { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnUint(v uint) error       { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnFloat32(v float32) error { return u.put(int8(v)) }
func (u *unfolderMapInt8) OnFloat64(v float64) error { return u.put(int8(v)) }

func (u *unfolderMapInt16) OnNil() error              { return u.put(0) }
func (u *unfolderMapInt16) OnInt8(v int8) error       { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnInt16(v int16) error     { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnInt32(v int32) error     { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnInt64(v int64) error     { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnInt(v int) error         { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnByte(v byte) error       { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnUint8(v uint8) error     { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnUint16(v uint16) error   { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnUint32(v uint32) error   { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnUint64(v uint64) error   { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnUint(v uint) error       { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnFloat32(v float32) error { return u.put(int16(v)) }
func (u *unfolderMapInt16) OnFloat64(v float64) error { return u.put(int16(v)) }

func (u *unfolderMapInt32) OnNil() error              { return u.put(0) }
func (u *unfolderMapInt32) OnInt8(v int8) error       { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnInt16(v int16) error     { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnInt32(v int32) error     { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnInt64(v int64) error     { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnInt(v int) error         { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnByte(v byte) error       { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnUint8(v uint8) error     { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnUint16(v uint16) error   { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnUint32(v uint32) error   { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnUint64(v uint64) error   { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnUint(v uint) error       { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnFloat32(v float32) error { return u.put(int32(v)) }
func (u *unfolderMapInt32) OnFloat64(v float64) error { return u.put(int32(v)) }

func (u *unfolderMapInt64) OnNil() error              { return u.put(0) }
func (u *unfolderMapInt64) OnInt8(v int8) error       { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnInt16(v int16) error     { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnInt32(v int32) error     { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnInt64(v int64) error     { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnInt(v int) error         { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnByte(v byte) error       { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnUint8(v uint8) error     { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnUint16(v uint16) error   { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnUint32(v uint32) error   { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnUint64(v uint64) error   { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnUint(v uint) error       { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnFloat32(v float32) error { return u.put(int64(v)) }
func (u *unfolderMapInt64) OnFloat64(v float64) error { return u.put(int64(v)) }

func (u *unfolderMapFloat32) OnNil() error              { return u.put(0) }
func (u *unfolderMapFloat32) OnInt8(v int8) error       { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnInt16(v int16) error     { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnInt32(v int32) error     { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnInt64(v int64) error     { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnInt(v int) error         { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnByte(v byte) error       { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnUint8(v uint8) error     { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnUint16(v uint16) error   { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnUint32(v uint32) error   { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnUint64(v uint64) error   { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnUint(v uint) error       { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnFloat32(v float32) error { return u.put(float32(v)) }
func (u *unfolderMapFloat32) OnFloat64(v float64) error { return u.put(float32(v)) }

func (u *unfolderMapFloat64) OnNil() error              { return u.put(0) }
func (u *unfolderMapFloat64) OnInt8(v int8) error       { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnInt16(v int16) error     { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnInt32(v int32) error     { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnInt64(v int64) error     { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnInt(v int) error         { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnByte(v byte) error       { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnUint8(v uint8) error     { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnUint16(v uint16) error   { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnUint32(v uint32) error   { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnUint64(v uint64) error   { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnUint(v uint) error       { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnFloat32(v float32) error { return u.put(float64(v)) }
func (u *unfolderMapFloat64) OnFloat64(v float64) error { return u.put(float64(v)) }
