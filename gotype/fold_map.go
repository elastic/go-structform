package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

var (
	// reflectMapAny     = reflectCast(map[string]interface{}(nil), visitMapInterface)
	reFoldMapBool    = liftFold(map[string]bool(nil), foldMapBool)
	reFoldMapInt     = liftFold(map[string]int(nil), foldMapInt)
	reFoldMapInt8    = liftFold(map[string]int8(nil), foldMapInt8)
	reFoldMapInt16   = liftFold(map[string]int16(nil), foldMapInt16)
	reFoldMapInt32   = liftFold(map[string]int32(nil), foldMapInt32)
	reFoldMapInt64   = liftFold(map[string]int64(nil), foldMapInt64)
	reFoldMapUint    = liftFold(map[string]uint(nil), foldMapUint)
	reFoldMapUint8   = liftFold(map[string]uint8(nil), foldMapUint8)
	reFoldMapUint16  = liftFold(map[string]uint16(nil), foldMapUint16)
	reFoldMapUint32  = liftFold(map[string]uint32(nil), foldMapUint32)
	reFoldMapUint64  = liftFold(map[string]uint64(nil), foldMapUint64)
	reFoldMapFloat32 = liftFold(map[string]float32(nil), foldMapFloat32)
	reFoldMapFloat64 = liftFold(map[string]float64(nil), foldMapFloat64)
	reFoldMapString  = liftFold(map[string]string(nil), foldMapString)
)

var tMapAny = reflect.TypeOf(map[string]interface{}(nil))

func reflectMapAny(c *context, V visitor, v reflect.Value) error {
	if v.Type().Name() != "" {
		v = v.Convert(tArrayAny)
	}
	return foldMapInterface(c, V, v.Interface())
}

func foldMapInterface(c *context, V visitor, v interface{}) error {
	m := v.(map[string]interface{})
	if err := V.OnObjectStart(len(m), structform.AnyType); err != nil {
		return err
	}

	for k, v := range m {
		if err := V.OnKey(k); err != nil {
			return err
		}
		if err := foldInterfaceValue(c, V, v); err != nil {
			return err
		}
	}
	return V.OnObjectFinished()
}

func foldMapBool(_ *context, V visitor, v interface{}) error {
	return V.OnBoolObject(v.(map[string]bool))
}

func foldMapString(_ *context, V visitor, v interface{}) error {
	return V.OnStringObject(v.(map[string]string))
}

func foldMapInt8(_ *context, V visitor, v interface{}) error {
	return V.OnInt8Object(v.(map[string]int8))
}

func foldMapInt16(_ *context, V visitor, v interface{}) error {
	return V.OnInt16Object(v.(map[string]int16))
}

func foldMapInt32(_ *context, V visitor, v interface{}) error {
	return V.OnInt32Object(v.(map[string]int32))
}

func foldMapInt64(_ *context, V visitor, v interface{}) error {
	return V.OnInt64Object(v.(map[string]int64))
}

func foldMapInt(_ *context, V visitor, v interface{}) error {
	return V.OnIntObject(v.(map[string]int))
}

func foldMapUint8(_ *context, V visitor, v interface{}) error {
	return V.OnUint8Object(v.(map[string]uint8))
}

func foldMapUint16(_ *context, V visitor, v interface{}) error {
	return V.OnUint16Object(v.(map[string]uint16))
}

func foldMapUint32(_ *context, V visitor, v interface{}) error {
	return V.OnUint32Object(v.(map[string]uint32))
}

func foldMapUint64(_ *context, V visitor, v interface{}) error {
	return V.OnUint64Object(v.(map[string]uint64))
}

func foldMapUint(_ *context, V visitor, v interface{}) error {
	return V.OnUintObject(v.(map[string]uint))
}

func foldMapFloat32(_ *context, V visitor, v interface{}) error {
	return V.OnFloat32Object(v.(map[string]float32))
}

func foldMapFloat64(_ *context, V visitor, v interface{}) error {
	return V.OnFloat64Object(v.(map[string]float64))
}
