package gotype

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
	"github.com/urso/go-structform/json"
	"github.com/urso/go-structform/ubjson"

	gojson "encoding/json"
)

type mapstr map[string]interface{}

var foldSamples = []struct {
	json  string
	value interface{}
}{
	// // primitives
	// {`null`, nil},
	// {`true`, true},
	// {`false`, false},
	// {`10`, int8(10)},
	// {`10`, int32(10)},
	// {`10`, int(10)},
	// {`10`, uint(10)},
	// {`10`, uint8(10)},
	// {`10`, uint16(10)},
	// {`10`, uint32(10)},
	// {`12340`, uint16(12340)},
	// {`1234567`, uint32(1234567)},
	// {`12345678190`, uint64(12345678190)},
	// {`-10`, int8(-10)},
	// {`-10`, int32(-10)},
	// {`-10`, int(-10)},
	// {`3.14`, float32(3.14)},
	// {`3.14`, float64(3.14)},
	// {`"test"`, "test"},
	// {`"test with \" being escaped"`, "test with \" being escaped"},

	// // arrays
	// {`[]`, []uint8{}},
	// {`[]`, []string{}},
	// {`[]`, []interface{}{}},
	// {`[]`, []struct{ A string }{}},
	// {`[[]]`, [][]uint8{{}}},
	// {`[[]]`, [][]string{{}}},
	// {`[[]]`, [][]interface{}{{}}},
	// {`[[]]`, [][]struct{ A string }{{}}},
	// {
	// 	`[null,true,false,12345678910,3.14,"test"]`,
	// 	[]interface{}{nil, true, false, uint64(12345678910), 3.14, "test"},
	// },
	// {`[1,2,3,4,5,6,7,8,9,10]`, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	// {`[1,2,3,4,5,6,7,8,9,10]`, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	// {`[1,2,3,4,5,6,7,8,9,10]`, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	// {`[1,2,3,4,5,6,7,8,9,10]`, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	// {`[1,2,3,4,5,6,7,8,9,10]`, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	// {`["testa","testb","testc"]`, []string{"testa", "testb", "testc"}},
	// {`["testa","testb","testc"]`, []interface{}{"testa", "testb", "testc"}},

	// // objects
	// {`{}`, map[string]interface{}{}},
	// {`{}`, map[string]int8{}},
	// {`{}`, map[string]int64{}},
	// {`{}`, map[string]struct{ A *int }{}},
	// {`{}`, mapstr{}},
	// {`{"a":null}`, map[string]interface{}{"a": nil}},
	// {`{"a":null}`, mapstr{"a": nil}},
	// {`{"a":null}`, struct{ A *int }{}},
	// {`{"a":null}`, struct{ A *struct{ B int } }{}},

	// unstable tests due to randomize field order in maps
	// {`{"a":true,"b":1,"c":"test"}`, map[string]interface{}{"a": true, "b": 1, "c": "test"}},
	// {`{"a":true,"b":1,"c":"test"}`, mapstr{"a": true, "b": 1, "c": "test"}},
	{`{"a":true,"b":1,"c":"test"}`, struct {
		A bool
		B int
		C string
	}{true, 1, "test"}},

	// {`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
	// 	map[string]interface{}{
	// 		"field": []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	},
	// },
	// {`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
	// 	map[string]interface{}{
	// 		"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	},
	// },
	// {`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
	// 	mapstr{
	// 		"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	},
	// },
	{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
		map[string][]int{
			"field": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	},
	{`{"field":[1,2,3,4,5,6,7,8,9,10]}`,
		struct {
			Field []int
		}{
			Field: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	},
}

type countWriter struct{ N int64 }

func (c *countWriter) Write(b []byte) (int, error) {
	n := len(b)
	c.N += int64(n)
	return n, nil
}

func TestIter2JsonConsistent(t *testing.T) {
	tests := foldSamples
	for i, test := range tests {
		t.Logf("run test (%v): %v (%T)", i, test.json, test.value)

		var buf bytes.Buffer
		iter := NewIterator(json.NewVisitor(&buf))
		err := iter.Fold(test.value)
		if err != nil {
			t.Error(err)
			continue
		}

		assert.Equal(t, test.json, buf.String())
	}
}

func BenchmarkCompareEncode(b *testing.B) {
	tests := foldSamples

	buf := &countWriter{}
	makeRun := func(v interface{}, enc func(interface{}) error) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf.N = 0
				if err := enc(v); err != nil {
					b.Error(err)
					return
				}
				b.SetBytes(buf.N)
			}
		}
	}

	for i, test := range tests {

		name := fmt.Sprintf("%v (%T)", test.json, test.value)
		b.Logf("run test (%v): %v", i, name)

		b.Run("ubjson "+name, makeRun(test.value, NewIterator(ubjson.NewVisitor(buf)).Fold))
		b.Run("json "+name, makeRun(test.value, NewIterator(json.NewVisitor(buf)).Fold))
		b.Run("go-codec-json "+name, makeRun(test.value, codec.NewEncoder(buf, &codec.JsonHandle{}).Encode))
		b.Run("go-codec-msgpack "+name, makeRun(test.value, codec.NewEncoder(buf, &codec.MsgpackHandle{}).Encode))
		b.Run("go-codec-cbor "+name, makeRun(test.value, codec.NewEncoder(buf, &codec.CborHandle{}).Encode))
		b.Run("go-json"+name, makeRun(test.value, gojson.NewEncoder(buf).Encode))
	}
}
