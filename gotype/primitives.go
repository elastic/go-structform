package gotype

import "reflect"

func visitNil(_ *context, V visitor, v interface{}) error     { return V.OnNil() }
func visitBool(_ *context, V visitor, v interface{}) error    { return V.OnBool(v.(bool)) }
func visitInt8(_ *context, V visitor, v interface{}) error    { return V.OnInt8(v.(int8)) }
func visitInt16(_ *context, V visitor, v interface{}) error   { return V.OnInt16(v.(int16)) }
func visitInt32(_ *context, V visitor, v interface{}) error   { return V.OnInt32(v.(int32)) }
func visitInt64(_ *context, V visitor, v interface{}) error   { return V.OnInt64(v.(int64)) }
func visitInt(_ *context, V visitor, v interface{}) error     { return V.OnInt64(int64(v.(int))) }
func visitByte(_ *context, V visitor, v interface{}) error    { return V.OnByte(v.(byte)) }
func visitUint8(_ *context, V visitor, v interface{}) error   { return V.OnUint8(v.(uint8)) }
func visitUint16(_ *context, V visitor, v interface{}) error  { return V.OnUint16(v.(uint16)) }
func visitUint32(_ *context, V visitor, v interface{}) error  { return V.OnUint32(v.(uint32)) }
func visitUint64(_ *context, V visitor, v interface{}) error  { return V.OnUint64(v.(uint64)) }
func visitUint(_ *context, V visitor, v interface{}) error    { return V.OnUint(v.(uint)) }
func visitFloat32(_ *context, V visitor, v interface{}) error { return V.OnFloat32(v.(float32)) }
func visitFloat64(_ *context, V visitor, v interface{}) error { return V.OnFloat64(v.(float64)) }
func visitString(_ *context, V visitor, v interface{}) error  { return V.OnString(v.(string)) }

func reflectNil(_ *context, V visitor, v reflect.Value) error    { return V.OnNil() }
func reflectBool(_ *context, V visitor, v reflect.Value) error   { return V.OnBool(v.Bool()) }
func reflectInt8(_ *context, V visitor, v reflect.Value) error   { return V.OnInt8(int8(v.Int())) }
func reflectInt16(_ *context, V visitor, v reflect.Value) error  { return V.OnInt16(int16(v.Int())) }
func reflectInt32(_ *context, V visitor, v reflect.Value) error  { return V.OnInt32(int32(v.Int())) }
func reflectInt64(_ *context, V visitor, v reflect.Value) error  { return V.OnInt64(int64(v.Int())) }
func reflectInt(_ *context, V visitor, v reflect.Value) error    { return V.OnInt64(int64(int(v.Int()))) }
func reflectUint8(_ *context, V visitor, v reflect.Value) error  { return V.OnUint8(uint8(v.Uint())) }
func reflectUint16(_ *context, V visitor, v reflect.Value) error { return V.OnUint16(uint16(v.Uint())) }
func reflectUint32(_ *context, V visitor, v reflect.Value) error { return V.OnUint32(uint32(v.Uint())) }
func reflectUint64(_ *context, V visitor, v reflect.Value) error { return V.OnUint64(uint64(v.Uint())) }
func reflectUint(_ *context, V visitor, v reflect.Value) error   { return V.OnUint(uint(v.Uint())) }
func reflectFloat32(_ *context, V visitor, v reflect.Value) error {
	return V.OnFloat32(float32(v.Float()))
}
func reflectFloat64(_ *context, V visitor, v reflect.Value) error { return V.OnFloat64(v.Float()) }
func reflectString(_ *context, V visitor, v reflect.Value) error  { return V.OnString(v.String()) }
