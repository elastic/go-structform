package unsafe

import (
	"reflect"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	b := *(*[]byte)(unsafe.Pointer(&bh))
	return b
}

func Bytes2Str(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *((*string)(unsafe.Pointer(&sh)))
}
