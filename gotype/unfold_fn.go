package gotype

import (
	"errors"

	"github.com/urso/go-structform"
)

// type reUnfoldFn func(u *Unfolder, sig unfoldSignal) error
type reUnfoldEmpty struct{}

type reUnfoldStruct struct {
	ctx *unfoldCtx
}

type unfolder interface {
	// primitives
	OnNil(*unfoldCtx) error
	OnBool(*unfoldCtx, bool) error
	OnString(*unfoldCtx, string) error
	OnInt8(*unfoldCtx, int8) error
	OnInt16(*unfoldCtx, int16) error
	OnInt32(*unfoldCtx, int32) error
	OnInt64(*unfoldCtx, int64) error
	OnInt(*unfoldCtx, int) error
	OnByte(*unfoldCtx, byte) error
	OnUint8(*unfoldCtx, uint8) error
	OnUint16(*unfoldCtx, uint16) error
	OnUint32(*unfoldCtx, uint32) error
	OnUint64(*unfoldCtx, uint64) error
	OnUint(*unfoldCtx, uint) error
	OnFloat32(*unfoldCtx, float32) error
	OnFloat64(*unfoldCtx, float64) error

	// array types
	OnArrayStart(*unfoldCtx, int, structform.BaseType) error
	OnArrayFinished(*unfoldCtx) error

	// object types
	OnObjectStart(*unfoldCtx, int, structform.BaseType) error
	OnObjectFinished(*unfoldCtx) error
	OnKey(*unfoldCtx, string) error
}

type unfoldSignal uint8

var errNotSUpported = errors.New("not supported")

//go:generate stringer -type=unfoldSignal
const (
	sigObjectStart unfoldSignal = iota
	sigObjectFinished
	sigObjectKey
	sigArrayStart
	sigArrayFinished
	sigNil
	sigBool
	sigString
	sigInt8
	sigInt16
	sigInt32
	sigInt64
	sigInt
	sigByte
	sigUint8
	sigUint16
	sigUint32
	sigUint64
	sigUint
	sigFloat32
	sigFloat64
)

func (reUnfoldEmpty) OnObjectStart(_ *unfoldCtx, _ int, _ structform.BaseType) error {
	return errNotSUpported
}
func (reUnfoldEmpty) OnObjectFinished(_ *unfoldCtx) error { return errNotSUpported }
func (reUnfoldEmpty) OnKey(_ *unfoldCtx, s string) error  { return errNotSUpported }
func (reUnfoldEmpty) OnArrayStart(_ *unfoldCtx, _ int, _ structform.BaseType) error {
	return errNotSUpported
}
func (reUnfoldEmpty) OnArrayFinished(_ *unfoldCtx) error      { return errNotSUpported }
func (reUnfoldEmpty) OnNil(_ *unfoldCtx) error                { return errNotSUpported }
func (reUnfoldEmpty) OnBool(_ *unfoldCtx, b bool) error       { return errNotSUpported }
func (reUnfoldEmpty) OnString(_ *unfoldCtx, s string) error   { return errNotSUpported }
func (reUnfoldEmpty) OnInt8(_ *unfoldCtx, i int8) error       { return errNotSUpported }
func (reUnfoldEmpty) OnInt16(_ *unfoldCtx, i int16) error     { return errNotSUpported }
func (reUnfoldEmpty) OnInt32(_ *unfoldCtx, i int32) error     { return errNotSUpported }
func (reUnfoldEmpty) OnInt64(_ *unfoldCtx, i int64) error     { return errNotSUpported }
func (reUnfoldEmpty) OnInt(_ *unfoldCtx, i int) error         { return errNotSUpported }
func (reUnfoldEmpty) OnByte(_ *unfoldCtx, b byte) error       { return errNotSUpported }
func (reUnfoldEmpty) OnUint8(_ *unfoldCtx, u uint8) error     { return errNotSUpported }
func (reUnfoldEmpty) OnUint16(_ *unfoldCtx, u uint16) error   { return errNotSUpported }
func (reUnfoldEmpty) OnUint32(_ *unfoldCtx, u uint32) error   { return errNotSUpported }
func (reUnfoldEmpty) OnUint64(_ *unfoldCtx, u uint64) error   { return errNotSUpported }
func (reUnfoldEmpty) OnUint(_ *unfoldCtx, u uint) error       { return errNotSUpported }
func (reUnfoldEmpty) OnFloat32(_ *unfoldCtx, f float32) error { return errNotSUpported }
func (reUnfoldEmpty) OnFloat64(_ *unfoldCtx, f float64) error { return errNotSUpported }

// func (f reUnfoldFn) OnObjectStart(len int, baseType BaseType) error
// func (f reUnfoldFn) OnObjectFinished() error
// func (f reUnfoldFn) OnKey(s string) error
// func (f reUnfoldFn) OnArrayStart(len int, baseType BaseType) error
// func (f reUnfoldFn) OnArrayFinished() error
// func (f reUnfoldFn) OnNil() error
// func (f reUnfoldFn) OnBool(b bool) error
// func (f reUnfoldFn) OnString(s string) error
// func (f reUnfoldFn) OnInt8(i int8) error
// func (f reUnfoldFn) OnInt16(i int16) error
// func (f reUnfoldFn) OnInt32(i int32) error
// func (f reUnfoldFn) OnInt64(i int64) error
// func (f reUnfoldFn) OnInt(i int) error
// func (f reUnfoldFn) OnByte(b byte) error
// func (f reUnfoldFn) OnUint8(u uint8) error
// func (f reUnfoldFn) OnUint16(u uint16) error
// func (f reUnfoldFn) OnUint32(u uint32) error
// func (f reUnfoldFn) OnUint64(u uint64) error
// func (f reUnfoldFn) OnUint(u uint) error
// func (f reUnfoldFn) OnFloat32(f float32) error
// func (f reUnfoldFn) OnFloat64(f float64) error
