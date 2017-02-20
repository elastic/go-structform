package gotype

import "reflect"

type options struct {
	tag string
}

var (
	tInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	tString    = reflect.TypeOf("")
	tBool      = reflect.TypeOf(true)
	tInt       = reflect.TypeOf(int(0))
	tInt8      = reflect.TypeOf(int8(0))
	tInt16     = reflect.TypeOf(int16(0))
	tInt32     = reflect.TypeOf(int32(0))
	tInt64     = reflect.TypeOf(int64(0))
	tUint      = reflect.TypeOf(uint(0))
	tByte      = reflect.TypeOf(byte(0))
	tUint8     = reflect.TypeOf(uint8(0))
	tUint16    = reflect.TypeOf(uint16(0))
	tUint32    = reflect.TypeOf(uint32(0))
	tUint64    = reflect.TypeOf(uint64(0))
	tFloat32   = reflect.TypeOf(float32(0))
	tFloat64   = reflect.TypeOf(float64(0))
)
