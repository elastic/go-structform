package gotype

import (
	"reflect"
	"sync"
	"unicode"
	"unicode/utf8"

	structform "github.com/urso/go-structform"
)

type typeVisitorRegistry struct {
	mu sync.RWMutex
	m  map[reflect.Type]reflectFn
}

var visitorRegistry = newTypeVisitorRegistry()

func getVisitorReflect(t reflect.Type) (reflectFn, error) {
	var err error

	f := visitorRegistry.find(t)
	if f != nil {
		return f, nil
	}

	f = getVisitorReflectPrimitive(t)
	if f != nil {
		return f, nil
	}

	switch t.Kind() {
	case reflect.Ptr:
		f, err = getVisitorPointer(t)
	case reflect.Struct:
		f, err = getVisitorReflectStruct(t)
	case reflect.Map:
		f, err = getVisitorReflectMap(t)
	case reflect.Slice:
		f, err = getVisitorReflectSlice(t)
	case reflect.Array:
		f, err = getVisitorReflectArray(t)
	case reflect.Interface:
		f = visitAnyReflect
	default:
		return nil, errUnsupported
	}

	if err != nil {
		return nil, err
	}
	visitorRegistry.set(t, f)
	return f, nil
}

func getVisitorReflectMap(t reflect.Type) (reflectFn, error) {
	if t.Key().Kind() != reflect.String {
		return nil, errMapRequiresStringKey
	}

	elemVisitor, err := getVisitorReflect(t.Elem())
	if err != nil {
		return nil, err
	}

	return func(c *context, V visitor, rv reflect.Value) error {
		count := rv.Len()

		if err := V.OnObjectStart(count, structform.AnyType); err != nil {
			return err
		}

		for _, k := range rv.MapKeys() {
			if err := V.OnKey(k.String()); err != nil {
				return err
			}
			if err := elemVisitor(c, V, rv.MapIndex(k)); err != nil {
				return err
			}
		}

		// iter keys/values

		return V.OnObjectFinished()
	}, nil
}

func getVisitorPointer(t reflect.Type) (reflectFn, error) {
	elemVisitor, err := getVisitorReflect(t.Elem())
	if err != nil {
		return nil, err
	}

	return func(c *context, V visitor, v reflect.Value) error {
		if v.IsNil() {
			return V.OnNil()
		}
		return elemVisitor(c, V, v.Elem())
	}, nil
}

func getVisitorReflectStruct(t reflect.Type) (reflectFn, error) {
	count := t.NumField()
	fields := make([]reflectFn, 0, count)

	for i := 0; i < count; i++ {
		fv, err := fieldVisitor(t, i)
		if err != nil {
			return nil, err
		}

		if fv == nil {
			continue
		}

		fields = append(fields, fv)
	}

	if len(fields) < cap(fields) {
		tmp := make([]reflectFn, len(fields))
		copy(tmp, fields)
		fields = tmp
	}

	return func(c *context, V visitor, v reflect.Value) error {
		if err := V.OnObjectStart(len(fields), structform.AnyType); err != nil {
			return err
		}

		for _, fv := range fields {
			if err := fv(c, V, v); err != nil {
				return err
			}
		}

		return V.OnObjectFinished()
	}, nil
}

func fieldVisitor(t reflect.Type, idx int) (reflectFn, error) {
	st := t.Field(idx)
	valueVisitor, err := getVisitorReflect(st.Type)
	if err != nil {
		return nil, err
	}

	name := st.Name
	rune, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsUpper(rune) {
		// ignore not exported fields
		return nil, nil
	}

	return func(c *context, V visitor, v reflect.Value) error {
		if err := V.OnKey(name); err != nil {
			return err
		}
		return valueVisitor(c, V, v.Field(idx))
	}, nil
}

func getVisitorReflectArray(t reflect.Type) (reflectFn, error) {
	return getVisitorReflectSlice(t)
}

func getVisitorReflectSlice(t reflect.Type) (reflectFn, error) {
	elemVisitor, err := getVisitorReflect(t.Elem())
	if err != nil {
		return nil, err
	}

	return func(c *context, V visitor, rv reflect.Value) error {
		count := rv.Len()

		if err := V.OnArrayStart(count, structform.AnyType); err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			if err := elemVisitor(c, V, rv.Index(i)); err != nil {
				return err
			}
		}

		return V.OnArrayFinished()
	}, nil
}

// TODO: create visitors casting the actual values via reflection instead of
//       golang type conversion:
func getVisitorReflectPrimitive(t reflect.Type) reflectFn {
	switch t.Kind() {
	case reflect.Bool:
		return reflectBool
	case reflect.Int:
		return reflectInt
	case reflect.Int8:
		return reflectInt8
	case reflect.Int16:
		return reflectInt16
	case reflect.Int32:
		return reflectInt32
	case reflect.Int64:
		return reflectInt64
	case reflect.Uint:
		return reflectUint
	case reflect.Uint8:
		return reflectUint8
	case reflect.Uint16:
		return reflectUint16
	case reflect.Uint32:
		return reflectUint32
	case reflect.Uint64:
		return reflectUint64
	case reflect.Float32:
		return reflectFloat32
	case reflect.Float64:
		return reflectFloat64
	case reflect.String:
		return reflectString

	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Interface:
			return reflectArrAny
		case reflect.Bool:
			return reflectArrBool
		case reflect.Int:
			return reflectArrInt
		case reflect.Int8:
			return reflectArrInt8
		case reflect.Int16:
			return reflectArrInt16
		case reflect.Int32:
			return reflectArrInt32
		case reflect.Int64:
			return reflectArrInt64
		case reflect.Uint:
			return reflectArrUint
		case reflect.Uint8:
			return reflectArrUint8
		case reflect.Uint16:
			return reflectArrUint16
		case reflect.Uint32:
			return reflectArrUint32
		case reflect.Uint64:
			return reflectArrUint64
		case reflect.Float32:
			return reflectArrFloat32
		case reflect.Float64:
			return reflectArrFloat64
		case reflect.String:
			return reflectArrString
		}

	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			return nil
		}

		switch t.Elem().Kind() {
		case reflect.Interface:
			return reflectMapAny
		case reflect.Bool:
			return reflectMapBool
		case reflect.Int:
			return reflectMapInt
		case reflect.Int8:
			return reflectMapInt8
		case reflect.Int16:
			return reflectMapInt16
		case reflect.Int32:
			return reflectMapInt32
		case reflect.Int64:
			return reflectMapInt64
		case reflect.Uint:
			return reflectMapUint
		case reflect.Uint8:
			return reflectMapUint8
		case reflect.Uint16:
			return reflectMapUint16
		case reflect.Uint32:
			return reflectMapUint32
		case reflect.Uint64:
			return reflectMapUint64
		case reflect.Float32:
			return reflectMapFloat32
		case reflect.Float64:
			return reflectMapFloat64
		case reflect.String:
			return reflectMapString
		}
	}

	return nil
}

func visitAnyReflect(c *context, V visitor, v reflect.Value) error {
	f, err := getVisitorReflect(v.Type())
	if err != nil {
		return err
	}
	return f(c, V, v)
}

func newTypeVisitorRegistry() *typeVisitorRegistry {
	return &typeVisitorRegistry{m: map[reflect.Type]reflectFn{}}
}

func (r *typeVisitorRegistry) find(t reflect.Type) reflectFn {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.m[t]
}

func (r *typeVisitorRegistry) set(t reflect.Type, f reflectFn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[t] = f
}

func reflectCast(sample interface{}, fn visitingFn) reflectFn {
	t := reflect.TypeOf(sample)
	return func(c *context, V visitor, v reflect.Value) error {
		if v.Type().Name() != "" {
			v = v.Convert(t)
		}
		return fn(c, V, v.Interface())
	}
}
