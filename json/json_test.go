package json

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var samples1 = []string{
	// // primitives
	"null",
	"true",
	"false",
	`"test"`,
	`"test with \" being escaped"`,
	`1`,
	`100`,
	`-10`,
	`7e+09`,

	// objects
	`{}`,
	`{"a":null}`,
	`{"a":true,"b":false}`,
	`{"a":1,"b":3.14,"c":"test"}`,

	// arrays
	`[]`,
	`[[]]`,
	`[null,true,false,1,3.14,7e+09]`,

	// mix of arrays and objects
	`[{}]`,
	`[{"a":[5,{}]}]`,

	// // streaming multiple objects
	`{}{}`,
	`{"a":1}{"a":3.14}{"a":"test"}`,
}

type noopWriter struct{}

func (noopWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func BenchmarkParserWriter(b *testing.B) {
	tests := samples1
	for i, test := range tests {
		b.Logf("run test (%v): %v", i, test)
		b.Run(test, func(b *testing.B) {
			buf := noopWriter{}
			parser := NewParser(NewVisitor(buf))
			for i := 0; i < b.N; i++ {
				err := parser.ParseString(test)
				if err != nil {
					b.Error(err)
					return
				}
			}
		})
	}
}

func TestCheckParserWriterConsistent(t *testing.T) {
	tests := samples1
	t.Logf("parse complete buffer")
	for i, test := range tests {
		t.Logf("run test (%v): %v", i, test)

		// parse and unparse the json
		var buf bytes.Buffer
		visitor := NewVisitor(&buf)
		err := ParseString(test, visitor)
		if err != nil {
			t.Error(err)
			continue
		}
		actual := buf.String()
		assert.Equal(t, test, actual)

		// check encoding by parsing already serialized content
		buf.Reset()
		err = ParseString(actual, visitor)
		if err != nil {
			t.Error(err)
			continue
		}
		actual = buf.String()
		assert.Equal(t, test, actual)
	}

	t.Logf("parse byte by byte")
byteTests:
	for i, test := range tests {
		t.Logf("run test (%v): `%v`", i, test)

		var buf bytes.Buffer
		visitor := NewVisitor(&buf)
		parser := NewParser(visitor)
		for _, c := range []byte(test) {
			err := parser.feed([]byte{c})
			if err != nil {
				t.Error(err)
				continue byteTests
			}
		}

		if err := parser.finalize(); err != nil {
			t.Error(err)
			continue
		}

		assert.Equal(t, test, buf.String())
	}
}
