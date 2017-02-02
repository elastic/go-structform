package gotype

import (
	"reflect"
	"strings"
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

func getVisitorReflect(c *context, t reflect.Type) (reflectFn, error) {
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
		f, err = getVisitorPointer(c, t)
	case reflect.Struct:
		f, err = getVisitorReflectStruct(c, t, false)
	case reflect.Map:
		f, err = getVisitorReflectMap(c, t)
	case reflect.Slice, reflect.Array:
		f, err = getVisitorReflectSlice(c, t)
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

func getVisitorReflectMap(c *context, t reflect.Type) (reflectFn, error) {
	iterVisitor, err := getVisitorReflectMapKeys(c, t)
	if err != nil {
		return nil, err
	}

	return func(c *context, V visitor, rv reflect.Value) error {
		if err := V.OnObjectStart(rv.Len(), structform.AnyType); err != nil {
			return err
		}
		if err := iterVisitor(c, V, rv); err != nil {
			return err
		}
		return V.OnObjectFinished()
	}, nil
}

func getVisitorReflectMapKeys(c *context, t reflect.Type) (reflectFn, error) {
	if t.Key().Kind() != reflect.String {
		return nil, errMapRequiresStringKey
	}

	elemVisitor, err := getVisitorReflect(c, t.Elem())
	if err != nil {
		return nil, err
	}

	return makeMapKeysVisitor(elemVisitor), nil
}

func makeMapKeysVisitor(elemVisitor reflectFn) reflectFn {
	return func(c *context, V visitor, rv reflect.Value) error {
		for _, k := range rv.MapKeys() {
			if err := V.OnKey(k.String()); err != nil {
				return err
			}
			if err := elemVisitor(c, V, rv.MapIndex(k)); err != nil {
				return err
			}
		}
		return nil
	}
}

func getVisitorPointer(c *context, t reflect.Type) (reflectFn, error) {
	N, bt := baseType(t)
	elemVisitor, err := getVisitorReflect(c, bt)
	if err != nil {
		return nil, err
	}
	return makePointerVisitor(N, elemVisitor), nil
}

func makePointerVisitor(N int, elemVisitor reflectFn) reflectFn {
	if N == 0 {
		return elemVisitor
	}

	return func(c *context, V visitor, v reflect.Value) error {
		for i := 0; i < N; i++ {
			if v.IsNil() {
				return V.OnNil()
			}
			v = v.Elem()
		}
		return elemVisitor(c, V, v)
	}
}

func getVisitorReflectStruct(c *context, t reflect.Type, inline bool) (reflectFn, error) {
	fields, err := getStructFieldsVisitors(c, t)
	if err != nil {
		return nil, err
	}

	if inline {
		return makeFieldsVisitor(fields), nil
	}
	return makeStructVisitor(fields), nil
}

func getStructFieldsVisitors(c *context, t reflect.Type) ([]reflectFn, error) {
	count := t.NumField()
	fields := make([]reflectFn, 0, count)

	for i := 0; i < count; i++ {
		fv, err := fieldVisitor(c, t, i)
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

	return fields, nil
}

func makeStructVisitor(fields []reflectFn) reflectFn {
	fieldsVisitor := makeFieldsVisitor(fields)
	return func(c *context, V visitor, v reflect.Value) error {
		if err := V.OnObjectStart(len(fields), structform.AnyType); err != nil {
			return err
		}
		if err := fieldsVisitor(c, V, v); err != nil {
			return err
		}
		return V.OnObjectFinished()
	}
}

func makeFieldsVisitor(fields []reflectFn) reflectFn {
	return func(c *context, V visitor, v reflect.Value) error {
		for _, fv := range fields {
			if err := fv(c, V, v); err != nil {
				return err
			}
		}
		return nil
	}
}

func fieldVisitor(c *context, t reflect.Type, idx int) (reflectFn, error) {
	st := t.Field(idx)

	name := st.Name
	rune, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsUpper(rune) {
		// ignore non exported fields
		return nil, nil
	}

	tagName, tagOpts := parseTags(st.Tag.Get(c.opts.tag))
	if tagName != "" {
		name = tagName
	} else {
		name = strings.ToLower(name)
	}

	valueVisitor, err := fieldValueVisitor(c, tagOpts, st.Type)
	if err != nil {
		return nil, err
	}

	return makeFieldVisitor(name, idx, valueVisitor)
}

func makeFieldVisitor(name string, idx int, fn reflectFn) (reflectFn, error) {
	return func(c *context, V visitor, v reflect.Value) error {
		if err := V.OnKey(name); err != nil {
			return err
		}
		return fn(c, V, v.Field(idx))
	}, nil
}

func fieldValueVisitor(c *context, opts tagOptions, t reflect.Type) (reflectFn, error) {
	if !opts.squash {
		return getVisitorReflect(c, t)
	}

	var (
		N, bt       = baseType(t)
		baseVisitor reflectFn
		err         error
	)

	switch bt.Kind() {
	case reflect.Struct:
		baseVisitor, err = getVisitorReflectStruct(c, bt, true)
	case reflect.Map:
		baseVisitor, err = getVisitorReflectMapKeys(c, bt)
	default:
		err = errSquashNeedObject
	}
	if err != nil {
		return nil, err
	}

	return makePointerVisitor(N, baseVisitor), nil
}

func getVisitorReflectSlice(c *context, t reflect.Type) (reflectFn, error) {
	elemVisitor, err := getVisitorReflect(c, t.Elem())
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
	f, err := getVisitorReflect(c, v.Type())
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

func baseType(t reflect.Type) (int, reflect.Type) {
	i := 0
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		i++
	}
	return i, t
}
