package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

var (
	// reflectMapAny     = reflectCast(map[string]interface{}(nil), visitMapInterface)
	reflectMapBool    = reflectCast(map[string]bool(nil), visitMapBool)
	reflectMapInt     = reflectCast(map[string]int(nil), visitMapInt)
	reflectMapInt8    = reflectCast(map[string]int8(nil), visitMapInt8)
	reflectMapInt16   = reflectCast(map[string]int16(nil), visitMapInt16)
	reflectMapInt32   = reflectCast(map[string]int32(nil), visitMapInt32)
	reflectMapInt64   = reflectCast(map[string]int64(nil), visitMapInt64)
	reflectMapUint    = reflectCast(map[string]uint(nil), visitMapUint)
	reflectMapUint8   = reflectCast(map[string]uint8(nil), visitMapUint8)
	reflectMapUint16  = reflectCast(map[string]uint16(nil), visitMapUint16)
	reflectMapUint32  = reflectCast(map[string]uint32(nil), visitMapUint32)
	reflectMapUint64  = reflectCast(map[string]uint64(nil), visitMapUint64)
	reflectMapFloat32 = reflectCast(map[string]float32(nil), visitMapFloat32)
	reflectMapFloat64 = reflectCast(map[string]float64(nil), visitMapFloat64)
	reflectMapString  = reflectCast(map[string]string(nil), visitMapString)
)

var tMapAny = reflect.TypeOf(map[string]interface{}(nil))

func reflectMapAny(c *context, V visitor, v reflect.Value) error {
	if v.Type().Name() != "" {
		v = v.Convert(tArrayAny)
	}
	return visitMapInterface(c, V, v.Interface())
}

func visitMapInterface(c *context, V visitor, v interface{}) error {
	m := v.(map[string]interface{})
	if err := V.OnObjectStart(len(m), structform.AnyType); err != nil {
		return err
	}

	for k, v := range m {
		if err := V.OnKey(k); err != nil {
			return err
		}
		if err := visitInterfaceValue(c, V, v); err != nil {
			return err
		}
	}
	return V.OnObjectFinished()
}

func visitMapBool(_ *context, V visitor, v interface{}) error {
	return V.OnBoolObject(v.(map[string]bool))
}

func visitMapString(_ *context, V visitor, v interface{}) error {
	return V.OnStringObject(v.(map[string]string))
}

func visitMapInt8(_ *context, V visitor, v interface{}) error {
	return V.OnInt8Object(v.(map[string]int8))
}

func visitMapInt16(_ *context, V visitor, v interface{}) error {
	return V.OnInt16Object(v.(map[string]int16))
}

func visitMapInt32(_ *context, V visitor, v interface{}) error {
	return V.OnInt32Object(v.(map[string]int32))
}

func visitMapInt64(_ *context, V visitor, v interface{}) error {
	return V.OnInt64Object(v.(map[string]int64))
}

func visitMapInt(_ *context, V visitor, v interface{}) error {
	return V.OnIntObject(v.(map[string]int))
}

func visitMapUint8(_ *context, V visitor, v interface{}) error {
	return V.OnUint8Object(v.(map[string]uint8))
}

func visitMapUint16(_ *context, V visitor, v interface{}) error {
	return V.OnUint16Object(v.(map[string]uint16))
}

func visitMapUint32(_ *context, V visitor, v interface{}) error {
	return V.OnUint32Object(v.(map[string]uint32))
}

func visitMapUint64(_ *context, V visitor, v interface{}) error {
	return V.OnUint64Object(v.(map[string]uint64))
}

func visitMapUint(_ *context, V visitor, v interface{}) error {
	return V.OnUintObject(v.(map[string]uint))
}

func visitMapFloat32(_ *context, V visitor, v interface{}) error {
	return V.OnFloat32Object(v.(map[string]float32))
}

func visitMapFloat64(_ *context, V visitor, v interface{}) error {
	return V.OnFloat64Object(v.(map[string]float64))
}
