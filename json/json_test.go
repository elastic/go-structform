// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package json

import (
	"bytes"
	"io"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	structform "github.com/elastic/go-structform"
	"github.com/elastic/go-structform/sftest"
)

func TestEncParseConsistent(t *testing.T) {
	testEncParseConsistent(t, Parse)
}

func TestEncDecoderConsistent(t *testing.T) {
	testEncParseConsistent(t, func(content []byte, to structform.Visitor) error {
		dec := NewBytesDecoder(content, to)
		err := dec.Next()
		if err == io.EOF {
			err = nil
		}
		return err
	})
}

func TestEncParseBytesConsistent(t *testing.T) {
	testEncParseConsistent(t, func(content []byte, to structform.Visitor) error {
		p := NewParser(to)
		for _, b := range content {
			err := p.feed([]byte{b})
			if err != nil {
				return err
			}
		}
		return p.finalize()
	})
}

func testEncParseConsistent(
	t *testing.T,
	parse func([]byte, structform.Visitor) error,
) {
	sftest.TestEncodeParseConsistent(t, sftest.Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := bytes.NewBuffer(nil)
			vs := NewVisitor(buf)

			return vs, func(to structform.Visitor) error {
				return parse(buf.Bytes(), to)
			}
		})
}

func TestStringEncoding(t *testing.T) {
	type testCase struct {
		in         string
		htmlEscape bool
		expected   string
	}

	cases := map[string]testCase{
		"no escaping required": testCase{
			in:       "hello world",
			expected: `"hello world"`,
		},
		"json required escaping": testCase{
			in:         `"hello \nworld"`,
			htmlEscape: false,
			expected:   `"\"hello \\nworld\""`,
		},
		"html escape with json required escaping": testCase{
			in:         `"hello \nworld"`,
			htmlEscape: true,
			expected:   `"\"hello \\nworld\""`,
		},
		"html with html escaping enabled": testCase{
			in:         "<hello>world</hello>",
			htmlEscape: true,
			expected:   `"\u003chello\u003eworld\u003c/hello\u003e"`,
		},
		"html without html escaping": testCase{
			in:         "<hello>world</hello>",
			htmlEscape: false,
			expected:   `"<hello>world</hello>"`,
		},
		"html with double quotes": testCase{
			in:         `<hello id="hola">`,
			htmlEscape: true,
			expected:   `"\u003chello id=\"hola\"\u003e"`,
		},
		"html with single quotes": testCase{
			in:         `<hello id='hola'>`,
			htmlEscape: true,
			expected:   `"\u003chello id='hola'\u003e"`,
		},
	}

	for name, test := range cases {
		in, htmlEscape, expected := test.in, test.htmlEscape, test.expected

		t.Run(name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			visitor := NewVisitor(buf)
			visitor.SetEscapeHTML(htmlEscape)
			visitor.OnString(in)

			actual := buf.String()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestEncodeIgnoreSpecialFloatValues(t *testing.T) {
	cases := map[string]struct {
		tokens sftest.Recording
		want   string
	}{
		"float64 NaN": {
			tokens: sftest.Recording{
				sftest.Float64Rec{math.NaN()},
			},
			want: `null`,
		},
		"float64 +Inf": {
			tokens: sftest.Recording{
				sftest.Float64Rec{math.Inf(1)},
			},
			want: `null`,
		},
		"float64 -Inf": {
			tokens: sftest.Recording{
				sftest.Float64Rec{math.Inf(-1)},
			},
			want: `null`,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			var buf strings.Builder
			visitor := NewVisitor(&buf)
			visitor.SetIgnoreInvalidFloat(true)

			err := test.tokens.Replay(visitor)
			require.NoError(t, err)

			require.Equal(t, test.want, buf.String())
		})
	}
}

func TestEncodeExplicitRadixPoint(t *testing.T) {
	cases := map[string]struct {
		tokens sftest.Recording
		want   string
	}{
		"1e+10": {
			tokens: sftest.Recording{
				sftest.Float64Rec{1e+10},
			},
			want: `1.0e+10`,
		},
		"1.1e+10": {
			tokens: sftest.Recording{
				sftest.Float64Rec{1.1e+10},
			},
			want: `1.1e+10`,
		},
		"1": {
			tokens: sftest.Recording{
				sftest.Float64Rec{1},
			},
			want: `1.0`,
		},
		"1.1": {
			tokens: sftest.Recording{
				sftest.Float64Rec{1.1},
			},
			want: `1.1`,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			var buf strings.Builder
			visitor := NewVisitor(&buf)
			visitor.SetExplicitRadixPoint(true)

			err := test.tokens.Replay(visitor)
			require.NoError(t, err)

			require.Equal(t, test.want, buf.String())
		})
	}
}
