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

package cborl

import (
	"bytes"
	"testing"

	structform "github.com/elastic/go-structform"
	"github.com/elastic/go-structform/sftest"
)

func TestEncParseConsistent(t *testing.T) {
	testEncParseConsistent(t, Parse)
}

func TestEncDecoderConsistent(t *testing.T) {
	testEncParseConsistent(t, func(content []byte, to structform.Visitor) error {
		dec := NewBytesDecoder(content, to)
		return dec.Next()
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
		return nil
	})
}

func testEncParseConsistent(
	t *testing.T,
	parse func([]byte, structform.Visitor) error,
) {
	samples := append([]sftest.Recording(nil), sftest.Samples...)
	// CBOR stores negative integers as -1-n, where n is unsigned. Exercise
	// both the signed limits and the transitions between encoded widths.
	for _, value := range []int64{
		-1, -24, -25,
		-128, -129, -200, -255, -256, -257,
		-32768, -32769, -38600, -65535, -65536, -65537,
		-2147483648, -2147483649, -4294967295, -4294967296, -4294967297,
		-9223372036854775808,
	} {
		samples = append(samples, sftest.Recording{sftest.Int64Rec{Value: value}})
	}
	samples = append(samples, sftest.Obj(1, structform.AnyType,
		"properties", sftest.Obj(1, structform.AnyType,
			"change_num", sftest.Int64Rec{Value: -38600},
		),
	))

	sftest.TestEncodeParseConsistent(t, samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := bytes.NewBuffer(nil)
			vs := NewVisitor(buf)

			return vs, func(to structform.Visitor) error {
				return parse(buf.Bytes(), to)
			}
		})
}
