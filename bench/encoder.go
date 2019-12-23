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

package bench

import (
	"bytes"
	stdjson "encoding/json"
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/ugorji/go/codec"

	"github.com/elastic/go-structform/cborl"
	"github.com/elastic/go-structform/gotype"
	"github.com/elastic/go-structform/json"
	"github.com/elastic/go-structform/ubjson"
)

type encoderFactory func(io.Writer) func(interface{}) error
type decoderFactory func([]byte) func(interface{}) error
type transcodeFactory func(io.Writer) func([]byte) error

func stdJSONEncoder(w io.Writer) func(interface{}) error {
	enc := stdjson.NewEncoder(w)
	return enc.Encode
}

func stdJSONDecoder(r io.Reader) func(interface{}) error {
	dec := stdjson.NewDecoder(r)
	return dec.Decode
}

func stdJSONBufDecoder(b []byte) func(interface{}) error {
	return stdJSONDecoder(bytes.NewReader(b))
}

func gocodecJSONDecoder(r io.Reader) func(interface{}) error {
	h := &codec.JsonHandle{}
	dec := codec.NewDecoder(r, h)
	return dec.Decode
}

func jsoniterDecoder(r io.Reader) func(interface{}) error {
	dec := jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(r)
	return func(v interface{}) error {
		if !dec.More() {
			return io.EOF
		}
		return dec.Decode(v)
	}
}

func jsoniterBufDecoder(b []byte) func(interface{}) error {
	return jsoniterDecoder(bytes.NewReader(b))
}

func structformJSONEncoder(w io.Writer) func(interface{}) error {
	vs := json.NewVisitor(w)
	folder, _ := gotype.NewIterator(vs)
	return folder.Fold
}

func structformUBJSONEncoder(w io.Writer) func(interface{}) error {
	vs := ubjson.NewVisitor(w)
	folder, _ := gotype.NewIterator(vs)
	return folder.Fold
}

func structformCBORLEncoder(w io.Writer) func(interface{}) error {
	vs := cborl.NewVisitor(w)
	folder, _ := gotype.NewIterator(vs)
	return folder.Fold
}

func structformJSONBufDecoder(keyCache int) func([]byte) func(interface{}) error {
	return func(b []byte) func(interface{}) error {
		u, _ := gotype.NewUnfolder(nil)
		dec := json.NewBytesDecoder(b, u)
		return makeStructformDecoder(u, dec.Next, keyCache)
	}
}

func structformUBJSONBufDecoder(keyCache int) func([]byte) func(interface{}) error {
	return func(b []byte) func(interface{}) error {
		u, _ := gotype.NewUnfolder(nil)
		dec := ubjson.NewBytesDecoder(b, u)
		return makeStructformDecoder(u, dec.Next, keyCache)
	}
}

func structformCBORLBufDecoder(keyCache int) func([]byte) func(interface{}) error {
	return func(b []byte) func(interface{}) error {
		u, _ := gotype.NewUnfolder(nil)
		dec := cborl.NewBytesDecoder(b, u)
		return makeStructformDecoder(u, dec.Next, keyCache)
	}
}

func makeStructformDecoder(
	u *gotype.Unfolder,
	next func() error,
	keyCache int,
) func(interface{}) error {
	if keyCache > 0 {
		u.EnableKeyCache(keyCache)
	}
	return func(v interface{}) error {
		if err := u.SetTarget(v); err != nil {
			return err
		}
		return next()
	}
}

func makeCBORL2JSONTranscoder(w io.Writer) func([]byte) error {
	j := json.NewVisitor(w)
	p := cborl.NewParser(j)
	return p.Parse
}

func makeUBJSON2JSONTranscoder(w io.Writer) func([]byte) error {
	j := json.NewVisitor(w)
	p := ubjson.NewParser(j)
	return p.Parse
}
