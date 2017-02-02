package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

var (
	// reFoldArrAny  = liftFold([]interface{}(nil), visitArrInterface)
	reFoldArrBool    = liftFold([]bool(nil), foldArrBool)
	reFoldArrInt     = liftFold([]int(nil), foldArrInt)
	reFoldArrInt8    = liftFold([]int8(nil), foldArrInt8)
	reFoldArrInt16   = liftFold([]int16(nil), foldArrInt16)
	reFoldArrInt32   = liftFold([]int32(nil), foldArrInt32)
	reFoldArrInt64   = liftFold([]int64(nil), foldArrInt64)
	reFoldArrUint    = liftFold([]uint(nil), foldArrUint)
	reFoldArrUint8   = liftFold([]uint8(nil), foldArrUint8)
	reFoldArrUint16  = liftFold([]uint16(nil), foldArrUint16)
	reFoldArrUint32  = liftFold([]uint32(nil), foldArrUint32)
	reFoldArrUint64  = liftFold([]uint64(nil), foldArrUint64)
	reFoldArrFloat32 = liftFold([]float32(nil), foldArrFloat32)
	reFoldArrFloat64 = liftFold([]float64(nil), foldArrFloat64)
	reFoldArrString  = liftFold([]string(nil), foldArrString)
)

var tArrayAny = reflect.TypeOf([]interface{}(nil))

func reFoldArrAny(c *context, V visitor, v reflect.Value) error {
	if v.Type().Name() != "" {
		v = v.Convert(tArrayAny)
	}
	return foldArrInterface(c, V, v.Interface())
}

func foldArrInterface(c *context, V visitor, v interface{}) error {
	a := v.([]interface{})
	if err := V.OnArrayStart(len(a), structform.AnyType); err != nil {
		return err
	}

	for _, v := range a {
		if err := foldInterfaceValue(c, V, v); err != nil {
			return err
		}
	}
	return V.OnArrayFinished()
}

func foldArrBool(_ *context, V visitor, v interface{}) error   { return V.OnBoolArray(v.([]bool)) }
func foldArrString(_ *context, V visitor, v interface{}) error { return V.OnStringArray(v.([]string)) }
func foldArrInt8(_ *context, V visitor, v interface{}) error   { return V.OnInt8Array(v.([]int8)) }
func foldArrInt16(_ *context, V visitor, v interface{}) error  { return V.OnInt16Array(v.([]int16)) }
func foldArrInt32(_ *context, V visitor, v interface{}) error  { return V.OnInt32Array(v.([]int32)) }
func foldArrInt64(_ *context, V visitor, v interface{}) error  { return V.OnInt64Array(v.([]int64)) }
func foldArrInt(_ *context, V visitor, v interface{}) error    { return V.OnIntArray(v.([]int)) }
func foldBytes(_ *context, V visitor, v interface{}) error     { return V.OnBytes(v.([]byte)) }
func foldArrUint8(_ *context, V visitor, v interface{}) error  { return V.OnUint8Array(v.([]uint8)) }
func foldArrUint16(_ *context, V visitor, v interface{}) error { return V.OnUint16Array(v.([]uint16)) }
func foldArrUint32(_ *context, V visitor, v interface{}) error { return V.OnUint32Array(v.([]uint32)) }
func foldArrUint64(_ *context, V visitor, v interface{}) error { return V.OnUint64Array(v.([]uint64)) }
func foldArrUint(_ *context, V visitor, v interface{}) error   { return V.OnUintArray(v.([]uint)) }
func foldArrFloat32(_ *context, V visitor, v interface{}) error {
	return V.OnFloat32Array(v.([]float32))
}
func foldArrFloat64(_ *context, V visitor, v interface{}) error {
	return V.OnFloat64Array(v.([]float64))
}
