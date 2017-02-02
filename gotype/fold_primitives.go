package gotype

import "reflect"

func foldNil(_ *context, V visitor, v interface{}) error     { return V.OnNil() }
func foldBool(_ *context, V visitor, v interface{}) error    { return V.OnBool(v.(bool)) }
func foldInt8(_ *context, V visitor, v interface{}) error    { return V.OnInt8(v.(int8)) }
func foldInt16(_ *context, V visitor, v interface{}) error   { return V.OnInt16(v.(int16)) }
func foldInt32(_ *context, V visitor, v interface{}) error   { return V.OnInt32(v.(int32)) }
func foldInt64(_ *context, V visitor, v interface{}) error   { return V.OnInt64(v.(int64)) }
func foldInt(_ *context, V visitor, v interface{}) error     { return V.OnInt64(int64(v.(int))) }
func foldByte(_ *context, V visitor, v interface{}) error    { return V.OnByte(v.(byte)) }
func foldUint8(_ *context, V visitor, v interface{}) error   { return V.OnUint8(v.(uint8)) }
func foldUint16(_ *context, V visitor, v interface{}) error  { return V.OnUint16(v.(uint16)) }
func foldUint32(_ *context, V visitor, v interface{}) error  { return V.OnUint32(v.(uint32)) }
func foldUint64(_ *context, V visitor, v interface{}) error  { return V.OnUint64(v.(uint64)) }
func foldUint(_ *context, V visitor, v interface{}) error    { return V.OnUint(v.(uint)) }
func foldFloat32(_ *context, V visitor, v interface{}) error { return V.OnFloat32(v.(float32)) }
func foldFloat64(_ *context, V visitor, v interface{}) error { return V.OnFloat64(v.(float64)) }
func foldString(_ *context, V visitor, v interface{}) error  { return V.OnString(v.(string)) }

func reFoldNil(_ *context, V visitor, v reflect.Value) error    { return V.OnNil() }
func reFoldBool(_ *context, V visitor, v reflect.Value) error   { return V.OnBool(v.Bool()) }
func reFoldInt8(_ *context, V visitor, v reflect.Value) error   { return V.OnInt8(int8(v.Int())) }
func reFoldInt16(_ *context, V visitor, v reflect.Value) error  { return V.OnInt16(int16(v.Int())) }
func reFoldInt32(_ *context, V visitor, v reflect.Value) error  { return V.OnInt32(int32(v.Int())) }
func reFoldInt64(_ *context, V visitor, v reflect.Value) error  { return V.OnInt64(v.Int()) }
func reFoldInt(_ *context, V visitor, v reflect.Value) error    { return V.OnInt64(int64(int(v.Int()))) }
func reFoldUint8(_ *context, V visitor, v reflect.Value) error  { return V.OnUint8(uint8(v.Uint())) }
func reFoldUint16(_ *context, V visitor, v reflect.Value) error { return V.OnUint16(uint16(v.Uint())) }
func reFoldUint32(_ *context, V visitor, v reflect.Value) error { return V.OnUint32(uint32(v.Uint())) }
func reFoldUint64(_ *context, V visitor, v reflect.Value) error { return V.OnUint64(v.Uint()) }
func reFoldUint(_ *context, V visitor, v reflect.Value) error   { return V.OnUint(uint(v.Uint())) }
func reFoldFloat32(_ *context, V visitor, v reflect.Value) error {
	return V.OnFloat32(float32(v.Float()))
}
func reFoldFloat64(_ *context, V visitor, v reflect.Value) error { return V.OnFloat64(v.Float()) }
func reFoldString(_ *context, V visitor, v reflect.Value) error  { return V.OnString(v.String()) }
