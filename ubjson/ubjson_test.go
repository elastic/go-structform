package ubjson

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/urso/go-structform"
	"github.com/urso/go-structform/json"
)

func TestCheckWriterParserConsistent(t *testing.T) {
	tests := []struct {
		json   string
		driver func(vs structform.Visitor)
	}{
		// primitives
		{
			"null",
			func(vs structform.Visitor) { vs.OnNil() },
		},
		{
			"true",
			func(vs structform.Visitor) { vs.OnBool(true) },
		},
		{
			"false",
			func(vs structform.Visitor) { vs.OnBool(false) },
		},
		{
			"10",
			func(vs structform.Visitor) { vs.OnByte(10) },
		},
		{
			"-10",
			func(vs structform.Visitor) { vs.OnInt8(-10) },
		},
		{
			"1024",
			func(vs structform.Visitor) { vs.OnInt64(1024) },
		},
		{
			"3.14",
			func(vs structform.Visitor) { vs.OnFloat32(3.14) },
		},
		{
			"7.1e+12",
			func(vs structform.Visitor) { vs.OnFloat64(7.1e12) },
		},
		{
			`"test"`,
			func(vs structform.Visitor) { vs.OnString("test") },
		},
		{
			`"test with \" being escaped"`,
			func(vs structform.Visitor) { vs.OnString("test with \" being escaped") },
		},

		// arrays
		{
			`[]`,
			func(vs structform.Visitor) {
				vs.OnArrayStart(-1, structform.AnyType)
				vs.OnArrayFinished()
			},
		},
		{
			`[[]]`,
			func(vs structform.Visitor) {
				vs.OnArrayStart(-1, structform.AnyType)
				vs.OnArrayStart(-1, structform.AnyType)
				vs.OnArrayFinished()
				vs.OnArrayFinished()
			},
		},
		{
			`[null,true,false,12345678910,3.14,"test"]`,
			func(vs structform.Visitor) {
				vs.OnArrayStart(-1, structform.AnyType)
				vs.OnNil()
				vs.OnBool(true)
				vs.OnBool(false)
				vs.OnInt64(12345678910)
				vs.OnFloat64(3.14)
				vs.OnString("test")
				vs.OnArrayFinished()
			},
		},
		{
			`[1,true]`,
			func(vs structform.Visitor) {
				vs.OnArrayStart(2, structform.AnyType)
				vs.OnInt8(1)
				vs.OnBool(true)
				vs.OnArrayFinished()
			},
		},
		{
			`[1,2,3]`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnInt8Array([]int8{1, 2, 3})
			},
		},
		{
			`["a","b","c"]`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnStringArray([]string{"a", "b", "c"})
			},
		},

		// objects
		{
			`{}`,
			func(vs structform.Visitor) {
				vs.OnObjectStart(-1, structform.AnyType)
				vs.OnObjectFinished()
			},
		},
		{
			`{"a":true,"b":1,"c":"test"}`,
			func(vs structform.Visitor) {
				vs.OnObjectStart(-1, structform.AnyType)
				vs.OnKey("a")
				vs.OnBool(true)
				vs.OnFieldNext()
				vs.OnKey("b")
				vs.OnInt8(1)
				vs.OnFieldNext()
				vs.OnKey("c")
				vs.OnString("test")
				vs.OnObjectFinished()
			},
		},
		{
			`{"a":[1,2,3]}`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnObjectStart(-1, structform.AnyType)
				ev.OnKey("a")
				ev.OnInt8Array([]int8{1, 2, 3})
				ev.OnObjectFinished()
			},
		},
		{
			`{}`,
			func(vs structform.Visitor) {
				vs.OnObjectStart(0, structform.AnyType)
				vs.OnObjectFinished()
			},
		},
		{
			`{"a":true,"b":1,"c":"test"}`,
			func(vs structform.Visitor) {
				vs.OnObjectStart(3, structform.AnyType)
				vs.OnKey("a")
				vs.OnBool(true)
				vs.OnFieldNext()
				vs.OnKey("b")
				vs.OnInt8(1)
				vs.OnFieldNext()
				vs.OnKey("c")
				vs.OnString("test")
				vs.OnObjectFinished()
			},
		},
		{
			`{"a":[1,2,3]}`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnObjectStart(1, structform.AnyType)
				ev.OnKey("a")
				ev.OnInt8Array([]int8{1, 2, 3})
				ev.OnObjectFinished()
			},
		},
		{
			`{"a":"test1","b":"test2","c":"test3"}`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnStringObject(map[string]string{
					"a": "test1",
					"b": "test2",
					"c": "test3",
				})
			},
		},
		{
			`{"a":1,"b":2,"c":3}`,
			func(vs structform.Visitor) {
				ev := structform.EnsureExtVisitor(vs)
				ev.OnUintObject(map[string]uint{
					"a": 1,
					"b": 2,
					"c": 3,
				})
			},
		},
	}

	for i, test := range tests {
		t.Logf("run test (%v): %v", i, test.json)

		var buf bytes.Buffer
		vs := NewVisitor(&buf)
		test.driver(vs)

		// try to parse into json
		var jsonBuf bytes.Buffer
		err := Parse(buf.Bytes(), json.NewVisitor(&jsonBuf))
		if err != nil {
			t.Error(err)
			continue
		}

		assert.Equal(t, test.json, jsonBuf.String())
	}
}
