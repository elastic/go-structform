package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type visitingFn func(c *context, V visitor, v interface{}) error

type reflectFn func(c *context, V visitor, v reflect.Value) error

type visitor interface {
	structform.ExtVisitor
}

type Iterator struct {
	v   visitor
	ctx context
}

type context struct {
	opts options
}

type options struct {
	tag string
}

func Iter(v interface{}, vs structform.Visitor) error {
	return NewIterator(vs).Iter(v)
}

func NewIterator(vs structform.Visitor) *Iterator {
	return &Iterator{
		v: structform.EnsureExtVisitor(vs).(visitor),
		ctx: context{
			opts: options{
				tag: "struct",
			},
		},
	}
}

func (i *Iterator) Iter(v interface{}) error {
	return visitInterfaceValue(&i.ctx, i.v, v)
}

func visitInterfaceValue(c *context, V visitor, v interface{}) error {
	if f := getVisitorGoTypes(v); f != nil {
		return f(c, V, v)
	}

	if tmp, f := getVisitorConvert(v); f != nil {
		return f(c, V, tmp)
	}

	return visitAnyReflect(c, V, reflect.ValueOf(v))
}

func getVisitorConvert(v interface{}) (interface{}, visitingFn) {
	t := reflect.TypeOf(v)
	cast := false

	switch t.Kind() {
	case reflect.Map:
		if cast = t.Name() != ""; cast {
			mt := reflect.MapOf(t.Key(), t.Elem())
			v = reflect.ValueOf(v).Convert(mt).Interface()
		}
	case reflect.Slice:
		if cast = t.Name() != ""; cast {
			mt := reflect.SliceOf(t.Elem())
			v = reflect.ValueOf(v).Convert(mt).Interface()
		}
	case reflect.Array:
		if cast = t.Name() != ""; cast {
			mt := reflect.ArrayOf(t.Len(), t.Elem())
			v = reflect.ValueOf(v).Convert(mt).Interface()
		}
	}

	return v, getVisitorGoTypes(v)
}

func getVisitorGoTypes(v interface{}) visitingFn {
	switch v.(type) {
	case nil:
		return visitNil

	case bool:
		return visitBool
	case []bool:
		return visitArrBool
	case map[string]bool:
		return visitMapBool

	case int8:
		return visitInt8
	case int16:
		return visitInt16
	case int32:
		return visitInt32
	case int64:
		return visitInt64
	case int:
		return visitInt

	case []int8:
		return visitArrInt8
	case []int16:
		return visitArrInt16
	case []int32:
		return visitArrInt32
	case []int64:
		return visitArrInt64
	case []int:
		return visitArrInt

	case map[string]int8:
		return visitMapInt8
	case map[string]int16:
		return visitMapInt16
	case map[string]int32:
		return visitMapInt32
	case map[string]int64:
		return visitMapInt64
	case map[string]int:
		return visitMapInt

		/*
			case byte:
				return visitByte
		*/
	case uint8:
		return visitUint8
	case uint16:
		return visitUint16
	case uint32:
		return visitUint32
	case uint64:
		return visitUint64
	case uint:
		return visitUint

	case []byte:
		return visitBytes
		/*
			case []uint8:
				return visitArrUint8
		*/
	case []uint16:
		return visitArrUint16
	case []uint32:
		return visitArrUint32
	case []uint64:
		return visitArrUint64
	case []uint:
		return visitArrUint

	case map[string]uint8:
		return visitMapUint8
	case map[string]uint16:
		return visitMapUint16
	case map[string]uint32:
		return visitMapUint32
	case map[string]uint64:
		return visitMapUint64
	case map[string]uint:
		return visitMapUint

	case float32:
		return visitFloat32
	case float64:
		return visitFloat64

	case []float32:
		return visitArrFloat32
	case []float64:
		return visitArrFloat64

	case map[string]float32:
		return visitMapFloat32
	case map[string]float64:
		return visitMapFloat64

	case string:
		return visitString

	case []string:
		return visitArrString
	case map[string]string:
		return visitMapString

	case []interface{}:
		return visitArrInterface
	case map[string]interface{}:
		return visitMapInterface
	}

	return nil
}
