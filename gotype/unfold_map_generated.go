// This file has been generated from 'unfold_map.yml', do not edit
package gotype

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/urso/go-structform"
)

type unfolderMapIfc struct {
	reUnfoldEmpty
}

var _singletonUnfolderMapIfc = &unfolderMapIfc{}

func newUnfolderMapIfc() *unfolderMapIfc {
	return _singletonUnfolderMapIfc
}

func (u *unfolderMapIfc) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return unfoldIfcStartSubMap(ctx, l, baseType)

	}

	return nil
}

func (u *unfolderMapIfc) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

	dtl.pop()
	ctx.ptr.pop()
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderMapIfc) OnKey(ctx *unfoldCtx, key string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitKey {
		return errUnexpectedObjectKey
	}

	ctx.key.push(key)
	dtl.current = unfoldWaitElem
	return nil
}

func (u *unfolderMapIfc) ptr(ctx *unfoldCtx) *map[string]interface{} {
	return (*map[string]interface{})(ctx.ptr.current)
}

func (u *unfolderMapIfc) put(ctx *unfoldCtx, v interface{}) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
	if *to == nil {
		*to = map[string]interface{}{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

type unfolderMapBool struct {
	reUnfoldEmpty
}

var _singletonUnfolderMapBool = &unfolderMapBool{}

func newUnfolderMapBool() *unfolderMapBool {
	return _singletonUnfolderMapBool
}

func (u *unfolderMapBool) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapBool) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapBool) ptr(ctx *unfoldCtx) *map[string]bool {
	return (*map[string]bool)(ctx.ptr.current)
}

func (u *unfolderMapBool) put(ctx *unfoldCtx, v bool) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapString = &unfolderMapString{}

func newUnfolderMapString() *unfolderMapString {
	return _singletonUnfolderMapString
}

func (u *unfolderMapString) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapString) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapString) ptr(ctx *unfoldCtx) *map[string]string {
	return (*map[string]string)(ctx.ptr.current)
}

func (u *unfolderMapString) put(ctx *unfoldCtx, v string) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapUint = &unfolderMapUint{}

func newUnfolderMapUint() *unfolderMapUint {
	return _singletonUnfolderMapUint
}

func (u *unfolderMapUint) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapUint) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapUint) ptr(ctx *unfoldCtx) *map[string]uint {
	return (*map[string]uint)(ctx.ptr.current)
}

func (u *unfolderMapUint) put(ctx *unfoldCtx, v uint) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapUint8 = &unfolderMapUint8{}

func newUnfolderMapUint8() *unfolderMapUint8 {
	return _singletonUnfolderMapUint8
}

func (u *unfolderMapUint8) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapUint8) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapUint8) ptr(ctx *unfoldCtx) *map[string]uint8 {
	return (*map[string]uint8)(ctx.ptr.current)
}

func (u *unfolderMapUint8) put(ctx *unfoldCtx, v uint8) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapUint16 = &unfolderMapUint16{}

func newUnfolderMapUint16() *unfolderMapUint16 {
	return _singletonUnfolderMapUint16
}

func (u *unfolderMapUint16) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapUint16) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapUint16) ptr(ctx *unfoldCtx) *map[string]uint16 {
	return (*map[string]uint16)(ctx.ptr.current)
}

func (u *unfolderMapUint16) put(ctx *unfoldCtx, v uint16) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapUint32 = &unfolderMapUint32{}

func newUnfolderMapUint32() *unfolderMapUint32 {
	return _singletonUnfolderMapUint32
}

func (u *unfolderMapUint32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapUint32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapUint32) ptr(ctx *unfoldCtx) *map[string]uint32 {
	return (*map[string]uint32)(ctx.ptr.current)
}

func (u *unfolderMapUint32) put(ctx *unfoldCtx, v uint32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapUint64 = &unfolderMapUint64{}

func newUnfolderMapUint64() *unfolderMapUint64 {
	return _singletonUnfolderMapUint64
}

func (u *unfolderMapUint64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapUint64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapUint64) ptr(ctx *unfoldCtx) *map[string]uint64 {
	return (*map[string]uint64)(ctx.ptr.current)
}

func (u *unfolderMapUint64) put(ctx *unfoldCtx, v uint64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapInt = &unfolderMapInt{}

func newUnfolderMapInt() *unfolderMapInt {
	return _singletonUnfolderMapInt
}

func (u *unfolderMapInt) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapInt) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapInt) ptr(ctx *unfoldCtx) *map[string]int {
	return (*map[string]int)(ctx.ptr.current)
}

func (u *unfolderMapInt) put(ctx *unfoldCtx, v int) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapInt8 = &unfolderMapInt8{}

func newUnfolderMapInt8() *unfolderMapInt8 {
	return _singletonUnfolderMapInt8
}

func (u *unfolderMapInt8) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapInt8) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapInt8) ptr(ctx *unfoldCtx) *map[string]int8 {
	return (*map[string]int8)(ctx.ptr.current)
}

func (u *unfolderMapInt8) put(ctx *unfoldCtx, v int8) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapInt16 = &unfolderMapInt16{}

func newUnfolderMapInt16() *unfolderMapInt16 {
	return _singletonUnfolderMapInt16
}

func (u *unfolderMapInt16) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapInt16) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapInt16) ptr(ctx *unfoldCtx) *map[string]int16 {
	return (*map[string]int16)(ctx.ptr.current)
}

func (u *unfolderMapInt16) put(ctx *unfoldCtx, v int16) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapInt32 = &unfolderMapInt32{}

func newUnfolderMapInt32() *unfolderMapInt32 {
	return _singletonUnfolderMapInt32
}

func (u *unfolderMapInt32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapInt32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapInt32) ptr(ctx *unfoldCtx) *map[string]int32 {
	return (*map[string]int32)(ctx.ptr.current)
}

func (u *unfolderMapInt32) put(ctx *unfoldCtx, v int32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapInt64 = &unfolderMapInt64{}

func newUnfolderMapInt64() *unfolderMapInt64 {
	return _singletonUnfolderMapInt64
}

func (u *unfolderMapInt64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapInt64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapInt64) ptr(ctx *unfoldCtx) *map[string]int64 {
	return (*map[string]int64)(ctx.ptr.current)
}

func (u *unfolderMapInt64) put(ctx *unfoldCtx, v int64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapFloat32 = &unfolderMapFloat32{}

func newUnfolderMapFloat32() *unfolderMapFloat32 {
	return _singletonUnfolderMapFloat32
}

func (u *unfolderMapFloat32) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapFloat32) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapFloat32) ptr(ctx *unfoldCtx) *map[string]float32 {
	return (*map[string]float32)(ctx.ptr.current)
}

func (u *unfolderMapFloat32) put(ctx *unfoldCtx, v float32) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
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

var _singletonUnfolderMapFloat64 = &unfolderMapFloat64{}

func newUnfolderMapFloat64() *unfolderMapFloat64 {
	return _singletonUnfolderMapFloat64
}

func (u *unfolderMapFloat64) OnObjectStart(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	// TODO: validate baseType

	dtl := &ctx.detail
	switch dtl.current {
	case unfoldWaitStart:
		dtl.current = unfoldWaitKey
	case unfoldWaitKey:
		return errExpectedObjectKey
	default:

		return errUnsupported

	}

	return nil
}

func (u *unfolderMapFloat64) OnObjectFinished(ctx *unfoldCtx) error {
	dtl := &ctx.detail
	if dtl.current != unfoldWaitKey {
		return errExpectedObjectKey
	}

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

func (u *unfolderMapFloat64) ptr(ctx *unfoldCtx) *map[string]float64 {
	return (*map[string]float64)(ctx.ptr.current)
}

func (u *unfolderMapFloat64) put(ctx *unfoldCtx, v float64) error {
	dtl := &ctx.detail

	if dtl.current != unfoldWaitElem {
		return errExpectedObjectKey
	}

	to := u.ptr(ctx)
	if *to == nil {
		*to = map[string]float64{}
	}
	(*to)[ctx.key.pop()] = v
	dtl.current = unfoldWaitKey
	return nil
}

func (u *unfolderMapIfc) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, nil)
}

func (u *unfolderMapIfc) OnBool(ctx *unfoldCtx, v bool) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnString(ctx *unfoldCtx, v string) error { return u.put(ctx, v) }

func (u *unfolderMapIfc) OnInt8(ctx *unfoldCtx, v int8) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnInt16(ctx *unfoldCtx, v int16) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnInt32(ctx *unfoldCtx, v int32) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnInt64(ctx *unfoldCtx, v int64) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnInt(ctx *unfoldCtx, v int) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnByte(ctx *unfoldCtx, v byte) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnUint8(ctx *unfoldCtx, v uint8) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnUint16(ctx *unfoldCtx, v uint16) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnUint32(ctx *unfoldCtx, v uint32) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnUint64(ctx *unfoldCtx, v uint64) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnUint(ctx *unfoldCtx, v uint) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnFloat32(ctx *unfoldCtx, v float32) error {
	return u.put(ctx, (interface{})(v))
}
func (u *unfolderMapIfc) OnFloat64(ctx *unfoldCtx, v float64) error {
	return u.put(ctx, (interface{})(v))
}

func (u *unfolderMapBool) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, false)
}

func (u *unfolderMapBool) OnBool(ctx *unfoldCtx, v bool) error { return u.put(ctx, v) }

func (u *unfolderMapString) OnNil(ctx *unfoldCtx) error {
	return u.put(ctx, "")
}

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

func unfoldIfcStartSubMap(ctx *unfoldCtx, l int, baseType structform.BaseType) error {
	tmp := makeMapPtr(ctx, l, baseType)
	if !ctx.trySetGotypeTarget(tmp) {
		// use reflection
		v := reflect.ValueOf(tmp)
		ctx.ptr.push(unsafe.Pointer(v.Pointer()))
		ctx.value.push(v.Elem())
		ctx.state.push(unfoldMapState)
		ctx.detail.push(unfoldWaitStart)
		ctx.unfolder.push(newUnfolderRefl())
	} else {
		// duplicate pointer, so we can assign after child-object is finished
		ctx.ptr.push(ctx.ptr.current)
	}

	ctx.baseType.push(baseType)
	return ctx.unfolder.current.OnObjectStart(ctx, l, baseType)
}

func unfoldIfcFinishSubMap(ctx *unfoldCtx) (interface{}, error) {
	child := ctx.ptr.pop()
	bt := ctx.baseType.pop()
	switch bt {

	case structform.AnyType:
		return *(*map[string]interface{})(child), nil

	case structform.BoolType:
		return *(*map[string]bool)(child), nil

	case structform.ByteType:
		return *(*map[string]byte)(child), nil

	case structform.Float32Type:
		return *(*map[string]float32)(child), nil

	case structform.Int16Type:
		return *(*map[string]int16)(child), nil

	case structform.Int32Type:
		return *(*map[string]int32)(child), nil

	case structform.Int64Type:
		return *(*map[string]int64)(child), nil

	case structform.Int8Type:
		return *(*map[string]int8)(child), nil

	case structform.IntType:
		return *(*map[string]int)(child), nil

	case structform.StringType:
		return *(*map[string]string)(child), nil

	case structform.Uint16Type:
		return *(*map[string]uint16)(child), nil

	case structform.Uint32Type:
		return *(*map[string]uint32)(child), nil

	case structform.Uint64Type:
		return *(*map[string]uint64)(child), nil

	case structform.Uint8Type:
		return *(*map[string]uint8)(child), nil

	case structform.UintType:
		return *(*map[string]uint)(child), nil

	case structform.ZeroType:
		return *(*map[string]interface{})(child), nil

	default:
		fmt.Println("base type: ", bt)
		return nil, errTODO()
	}
}

func makeMapPtr(ctx *unfoldCtx, l int, bt structform.BaseType) interface{} {
	switch bt {

	case structform.AnyType:
		return new(map[string]interface{})

	case structform.BoolType:
		return new(map[string]bool)

	case structform.ByteType:
		return new(map[string]byte)

	case structform.Float32Type:
		return new(map[string]float32)

	case structform.Int16Type:
		return new(map[string]int16)

	case structform.Int32Type:
		return new(map[string]int32)

	case structform.Int64Type:
		return new(map[string]int64)

	case structform.Int8Type:
		return new(map[string]int8)

	case structform.IntType:
		return new(map[string]int)

	case structform.StringType:
		return new(map[string]string)

	case structform.Uint16Type:
		return new(map[string]uint16)

	case structform.Uint32Type:
		return new(map[string]uint32)

	case structform.Uint64Type:
		return new(map[string]uint64)

	case structform.Uint8Type:
		return new(map[string]uint8)

	case structform.UintType:
		return new(map[string]uint)

	case structform.ZeroType:
		return new(map[string]interface{})

	default:
		panic("invalid type code")
		return nil
	}
}
