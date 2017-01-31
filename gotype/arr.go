package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

var (
	// reflectArrAny     = reflectCast([]interface{}(nil), visitArrInterface)
	reflectArrBool    = reflectCast([]bool(nil), visitArrBool)
	reflectArrInt     = reflectCast([]int(nil), visitArrInt)
	reflectArrInt8    = reflectCast([]int8(nil), visitArrInt8)
	reflectArrInt16   = reflectCast([]int16(nil), visitArrInt16)
	reflectArrInt32   = reflectCast([]int32(nil), visitArrInt32)
	reflectArrInt64   = reflectCast([]int64(nil), visitArrInt64)
	reflectArrUint    = reflectCast([]uint(nil), visitArrUint)
	reflectArrUint8   = reflectCast([]uint8(nil), visitArrUint8)
	reflectArrUint16  = reflectCast([]uint16(nil), visitArrUint16)
	reflectArrUint32  = reflectCast([]uint32(nil), visitArrUint32)
	reflectArrUint64  = reflectCast([]uint64(nil), visitArrUint64)
	reflectArrFloat32 = reflectCast([]float32(nil), visitArrFloat32)
	reflectArrFloat64 = reflectCast([]float64(nil), visitArrFloat64)
	reflectArrString  = reflectCast([]string(nil), visitArrString)
)

var tArrayAny = reflect.TypeOf([]interface{}(nil))

func reflectArrAny(c *context, V visitor, v reflect.Value) error {
	if v.Type().Name() != "" {
		v = v.Convert(tArrayAny)
	}
	return visitArrInterface(c, V, v.Interface())
}

func visitArrInterface(c *context, V visitor, v interface{}) error {
	a := v.([]interface{})
	if err := V.OnArrayStart(len(a), structform.AnyType); err != nil {
		return err
	}

	for _, v := range a {
		if err := visitInterfaceValue(c, V, v); err != nil {
			return err
		}
	}
	return V.OnArrayFinished()
}

func visitArrBool(_ *context, V visitor, v interface{}) error   { return V.OnBoolArray(v.([]bool)) }
func visitArrString(_ *context, V visitor, v interface{}) error { return V.OnStringArray(v.([]string)) }
func visitArrInt8(_ *context, V visitor, v interface{}) error   { return V.OnInt8Array(v.([]int8)) }
func visitArrInt16(_ *context, V visitor, v interface{}) error  { return V.OnInt16Array(v.([]int16)) }
func visitArrInt32(_ *context, V visitor, v interface{}) error  { return V.OnInt32Array(v.([]int32)) }
func visitArrInt64(_ *context, V visitor, v interface{}) error  { return V.OnInt64Array(v.([]int64)) }
func visitArrInt(_ *context, V visitor, v interface{}) error    { return V.OnIntArray(v.([]int)) }
func visitBytes(_ *context, V visitor, v interface{}) error     { return V.OnBytes(v.([]byte)) }
func visitArrUint8(_ *context, V visitor, v interface{}) error  { return V.OnUint8Array(v.([]uint8)) }
func visitArrUint16(_ *context, V visitor, v interface{}) error { return V.OnUint16Array(v.([]uint16)) }
func visitArrUint32(_ *context, V visitor, v interface{}) error { return V.OnUint32Array(v.([]uint32)) }
func visitArrUint64(_ *context, V visitor, v interface{}) error { return V.OnUint64Array(v.([]uint64)) }
func visitArrUint(_ *context, V visitor, v interface{}) error   { return V.OnUintArray(v.([]uint)) }
func visitArrFloat32(_ *context, V visitor, v interface{}) error {
	return V.OnFloat32Array(v.([]float32))
}
func visitArrFloat64(_ *context, V visitor, v interface{}) error {
	return V.OnFloat64Array(v.([]float64))
}
