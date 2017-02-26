package gotype

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"

	structform "github.com/urso/go-structform"
)

type unfolderStruct struct {
	unfolderErrExpectKey
	fields map[string]fieldUnfolder
}

type unfolderStructStart struct {
	unfolderErrObjectStart
}

type fieldUnfolder struct {
	offset uintptr
	size   int

	initState func(ctx *unfoldCtx, sp unsafe.Pointer)
}

var (
	_singletonUnfolderStructStart = &unfolderStructStart{}

	_ignoredField = &fieldUnfolder{
		initState: _singletonUnfoldIgnorePtr.initState,
	}
)

func createUnfolderReflStruct(ctx *unfoldCtx, t reflect.Type) (*unfolderStruct, error) {
	// assume t is pointer to struct
	t = t.Elem()

	count := t.NumField()
	fields := map[string]fieldUnfolder{}

	for i := 0; i < count; i++ {
		name, fu, ok, err := createFieldUnfolder(ctx, t, i)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		fields[name] = fu
	}

	u := &unfolderStruct{fields: fields}
	return u, nil
}

func createFieldUnfolder(ctx *unfoldCtx, t reflect.Type, idx int) (string, fieldUnfolder, bool, error) {
	st := t.Field(idx)

	name := st.Name
	rune, _ := utf8.DecodeRuneInString(name)
	if !unicode.IsUpper(rune) {
		return "", fieldUnfolder{}, false, nil
	}

	tagName, _ := parseTags(st.Tag.Get(ctx.opts.tag))
	if tagName != "" {
		name = tagName
	} else {
		name = strings.ToLower(name)
	}

	if pu := lookupGoPtrUnfolder(st.Type); pu != nil {
		fu := fieldUnfolder{
			offset:    st.Offset,
			size:      int(st.Type.Size()),
			initState: pu.initState,
		}
		return name, fu, true, nil
	}

	targetType := reflect.PtrTo(st.Type)
	ru, err := lookupReflUnfolder(ctx, targetType)
	if err != nil {
		return "", fieldUnfolder{}, false, err
	}

	fu := fieldUnfolder{
		offset:    st.Offset,
		size:      int(st.Type.Size()),
		initState: wrapReflUnfolder(targetType, ru),
	}
	return name, fu, true, nil
}

func wrapReflUnfolder(t reflect.Type, ru reflUnfolder) func(*unfoldCtx, unsafe.Pointer) {
	return func(ctx *unfoldCtx, ptr unsafe.Pointer) {
		v := reflect.NewAt(reflect.PtrTo(t), ptr)
		ru.initState(ctx, v)
	}
}

func (u *unfolderStruct) initState(ctx *unfoldCtx, v reflect.Value) {
	u.initStatePtr(ctx, unsafe.Pointer(v.Pointer()))
}

func (u *unfolderStruct) initStatePtr(ctx *unfoldCtx, ptr unsafe.Pointer) {
	ctx.ptr.push(ptr)
	ctx.unfolder.push(u)
	ctx.unfolder.push(_singletonUnfolderStructStart)
}

func (u *unfolderStructStart) OnObjectStart(ctx *unfoldCtx, l int, bt structform.BaseType) error {
	ctx.unfolder.pop()
	return nil
}

func (u *unfolderStruct) OnObjectFinished(ctx *unfoldCtx) error {
	ctx.unfolder.pop()
	ctx.ptr.pop()
	return nil
}

func (u *unfolderStruct) OnKey(ctx *unfoldCtx, key string) error {
	field, exists := u.fields[key]
	if !exists {
		_ignoredField.initState(ctx, nil)
		return nil
	}

	structPtr := ctx.ptr.current
	fieldAddr := uintptr(structPtr) + field.offset
	fieldPtr := unsafe.Pointer(fieldAddr)
	field.initState(ctx, fieldPtr)
	return nil
}
