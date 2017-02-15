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

func (reUnfoldEmpty) OnObjectStart(_ int, _ structform.BaseType) error { return errNotSUpported }
func (reUnfoldEmpty) OnObjectFinished() error                          { return errNotSUpported }
func (reUnfoldEmpty) OnKey(s string) error                             { return errNotSUpported }
func (reUnfoldEmpty) OnArrayStart(_ int, _ structform.BaseType) error  { return errNotSUpported }
func (reUnfoldEmpty) OnArrayFinished() error                           { return errNotSUpported }
func (reUnfoldEmpty) OnNil() error                                     { return errNotSUpported }
func (reUnfoldEmpty) OnBool(b bool) error                              { return errNotSUpported }
func (reUnfoldEmpty) OnString(s string) error                          { return errNotSUpported }
func (reUnfoldEmpty) OnInt8(i int8) error                              { return errNotSUpported }
func (reUnfoldEmpty) OnInt16(i int16) error                            { return errNotSUpported }
func (reUnfoldEmpty) OnInt32(i int32) error                            { return errNotSUpported }
func (reUnfoldEmpty) OnInt64(i int64) error                            { return errNotSUpported }
func (reUnfoldEmpty) OnInt(i int) error                                { return errNotSUpported }
func (reUnfoldEmpty) OnByte(b byte) error                              { return errNotSUpported }
func (reUnfoldEmpty) OnUint8(u uint8) error                            { return errNotSUpported }
func (reUnfoldEmpty) OnUint16(u uint16) error                          { return errNotSUpported }
func (reUnfoldEmpty) OnUint32(u uint32) error                          { return errNotSUpported }
func (reUnfoldEmpty) OnUint64(u uint64) error                          { return errNotSUpported }
func (reUnfoldEmpty) OnUint(u uint) error                              { return errNotSUpported }
func (reUnfoldEmpty) OnFloat32(f float32) error                        { return errNotSUpported }
func (reUnfoldEmpty) OnFloat64(f float64) error                        { return errNotSUpported }

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
