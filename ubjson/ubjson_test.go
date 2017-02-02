package ubjson

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/urso/go-structform"
	"github.com/urso/go-structform/json"
)

var samples1 = []struct {
	json   string
	driver func(vs structform.ExtVisitor)
}{
	// primitives
	{
		"null",
		func(vs structform.ExtVisitor) { vs.OnNil() },
	},
	{
		"true",
		func(vs structform.ExtVisitor) { vs.OnBool(true) },
	},
	{
		"false",
		func(vs structform.ExtVisitor) { vs.OnBool(false) },
	},
	{
		"10",
		func(vs structform.ExtVisitor) { vs.OnByte(10) },
	},
	{
		"-10",
		func(vs structform.ExtVisitor) { vs.OnInt8(-10) },
	},
	{
		"1024",
		func(vs structform.ExtVisitor) { vs.OnInt64(1024) },
	},
	{
		"3.14",
		func(vs structform.ExtVisitor) { vs.OnFloat32(3.14) },
	},
	{
		"7.1e+12",
		func(vs structform.ExtVisitor) { vs.OnFloat64(7.1e12) },
	},
	{
		`"test"`,
		func(vs structform.ExtVisitor) { vs.OnString("test") },
	},
	{
		`"test with \" being escaped"`,
		func(vs structform.ExtVisitor) { vs.OnString("test with \" being escaped") },
	},

	// arrays
	{
		`[]`,
		func(vs structform.ExtVisitor) {
			vs.OnArrayStart(-1, structform.AnyType)
			vs.OnArrayFinished()
		},
	},
	{
		`[[]]`,
		func(vs structform.ExtVisitor) {
			vs.OnArrayStart(-1, structform.AnyType)
			vs.OnArrayStart(-1, structform.AnyType)
			vs.OnArrayFinished()
			vs.OnArrayFinished()
		},
	},
	{
		`[null,true,false,12345678910,3.14,"test"]`,
		func(vs structform.ExtVisitor) {
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
		func(vs structform.ExtVisitor) {
			vs.OnArrayStart(2, structform.AnyType)
			vs.OnInt8(1)
			vs.OnBool(true)
			vs.OnArrayFinished()
		},
	},
	{
		`[1,2,3]`,
		func() func(vs structform.ExtVisitor) {
			a := []int8{1, 2, 3}
			return func(vs structform.ExtVisitor) {
				vs.OnInt8Array(a)
			}
		}(),
	},
	{
		`["a","b","c"]`,
		func() func(vs structform.ExtVisitor) {
			a := []string{"a", "b", "c"}
			return func(vs structform.ExtVisitor) {
				vs.OnStringArray(a)
			}
		}(),
	},

	// objects
	{
		`{}`,
		func(vs structform.ExtVisitor) {
			vs.OnObjectStart(-1, structform.AnyType)
			vs.OnObjectFinished()
		},
	},
	{
		`{"a":true,"b":1,"c":"test"}`,
		func(vs structform.ExtVisitor) {
			vs.OnObjectStart(-1, structform.AnyType)
			vs.OnKey("a")
			vs.OnBool(true)
			vs.OnKey("b")
			vs.OnInt8(1)
			vs.OnKey("c")
			vs.OnString("test")
			vs.OnObjectFinished()
		},
	},
	{
		`{"a":[1,2,3]}`,
		func() func(structform.ExtVisitor) {
			a := []int8{1, 2, 3}
			return func(vs structform.ExtVisitor) {
				vs.OnObjectStart(-1, structform.AnyType)
				vs.OnKey("a")
				vs.OnInt8Array(a)
				vs.OnObjectFinished()
			}
		}(),
	},
	{
		`{}`,
		func(vs structform.ExtVisitor) {
			vs.OnObjectStart(0, structform.AnyType)
			vs.OnObjectFinished()
		},
	},
	{
		`{"a":true,"b":1,"c":"test"}`,
		func(vs structform.ExtVisitor) {
			vs.OnObjectStart(3, structform.AnyType)
			vs.OnKey("a")
			vs.OnBool(true)
			vs.OnKey("b")
			vs.OnInt8(1)
			vs.OnKey("c")
			vs.OnString("test")
			vs.OnObjectFinished()
		},
	},
	{
		`{"a":[1,2,3]}`,

		func() func(structform.ExtVisitor) {
			a := []int8{1, 2, 3}
			return func(vs structform.ExtVisitor) {
				vs.OnObjectStart(1, structform.AnyType)
				vs.OnKey("a")
				vs.OnInt8Array(a)
				vs.OnObjectFinished()
			}
		}(),
	},
	{
		`{"a":"test1","b":"test2","c":"test3"}`,
		func() func(vs structform.ExtVisitor) {
			m := map[string]string{"a": "test1", "b": "test2", "c": "test3"}
			return func(vs structform.ExtVisitor) {
				vs.OnStringObject(m)
			}
		}(),
	},
	{
		`{"a":1,"b":2,"c":3}`,
		func() func(vs structform.ExtVisitor) {
			m := map[string]uint{"a": 1, "b": 2, "c": 3}
			return func(vs structform.ExtVisitor) {
				vs.OnUintObject(m)
			}
		}(),
	},
}

type noopWriter struct{}

func (noopWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkUBJSONVisitor(b *testing.B) {
	tests := samples1
	for _, test := range tests {
		vs := NewVisitor(noopWriter{})
		b.Run(test.json, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.driver(vs)
			}
		})
	}
}

func BenchmarkTranscodeUBJSONToJSON(b *testing.B) {
	tests := samples1
	buffers := make([]bytes.Buffer, len(tests))
	for i, test := range tests {
		vs := NewVisitor(&buffers[i])
		test.driver(vs)
	}

	for i, test := range tests {
		transcoder := NewParser(json.NewVisitor(noopWriter{}))
		buf := buffers[i].Bytes()
		b.Run(test.json, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err := transcoder.Parse(buf)
				if err != nil {
					b.Error(err)
					return
				}
			}
		})
	}
}

func TestCheckWriterParserConsistent(t *testing.T) {
	tests := samples1
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
