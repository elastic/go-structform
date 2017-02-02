package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type foldFn func(c *context, V visitor, v interface{}) error

type reFoldFn func(c *context, V visitor, v reflect.Value) error

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

func Fold(v interface{}, vs structform.Visitor) error {
	return NewIterator(vs).Fold(v)
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

func (i *Iterator) Fold(v interface{}) error {
	return foldInterfaceValue(&i.ctx, i.v, v)
}

func foldInterfaceValue(c *context, V visitor, v interface{}) error {
	if f := getFoldGoTypes(v); f != nil {
		return f(c, V, v)
	}

	if tmp, f := getFoldConvert(v); f != nil {
		return f(c, V, tmp)
	}

	return foldAnyReflect(c, V, reflect.ValueOf(v))
}

func getFoldConvert(v interface{}) (interface{}, foldFn) {
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

	return v, getFoldGoTypes(v)
}

func getFoldGoTypes(v interface{}) foldFn {
	switch v.(type) {
	case nil:
		return foldNil

	case bool:
		return foldBool
	case []bool:
		return foldArrBool
	case map[string]bool:
		return foldMapBool

	case int8:
		return foldInt8
	case int16:
		return foldInt16
	case int32:
		return foldInt32
	case int64:
		return foldInt64
	case int:
		return foldInt

	case []int8:
		return foldArrInt8
	case []int16:
		return foldArrInt16
	case []int32:
		return foldArrInt32
	case []int64:
		return foldArrInt64
	case []int:
		return foldArrInt

	case map[string]int8:
		return foldMapInt8
	case map[string]int16:
		return foldMapInt16
	case map[string]int32:
		return foldMapInt32
	case map[string]int64:
		return foldMapInt64
	case map[string]int:
		return foldMapInt

		/*
			case byte:
				return visitByte
		*/
	case uint8:
		return foldUint8
	case uint16:
		return foldUint16
	case uint32:
		return foldUint32
	case uint64:
		return foldUint64
	case uint:
		return foldUint

	case []byte:
		return foldBytes
		/*
			case []uint8:
				return visitArrUint8
		*/
	case []uint16:
		return foldArrUint16
	case []uint32:
		return foldArrUint32
	case []uint64:
		return foldArrUint64
	case []uint:
		return foldArrUint

	case map[string]uint8:
		return foldMapUint8
	case map[string]uint16:
		return foldMapUint16
	case map[string]uint32:
		return foldMapUint32
	case map[string]uint64:
		return foldMapUint64
	case map[string]uint:
		return foldMapUint

	case float32:
		return foldFloat32
	case float64:
		return foldFloat64

	case []float32:
		return foldArrFloat32
	case []float64:
		return foldArrFloat64

	case map[string]float32:
		return foldMapFloat32
	case map[string]float64:
		return foldMapFloat64

	case string:
		return foldString

	case []string:
		return foldArrString
	case map[string]string:
		return foldMapString

	case []interface{}:
		return foldArrInterface
	case map[string]interface{}:
		return foldMapInterface
	}

	return nil
}
