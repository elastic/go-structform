package bench

import (
	"bytes"
	stdjson "encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/ugorji/go/codec"
	"github.com/urso/go-structform/cborl"
	"github.com/urso/go-structform/gotype"
	"github.com/urso/go-structform/json"
	"github.com/urso/go-structform/ubjson"
)

type encoderFactory func(io.Writer) func(interface{}) error
type decoderFactory func(io.Reader) func(interface{}) error
type decoderFactoryBuf func([]byte) func(interface{}) error

func BenchmarkDecodeBeatsEventsFile(b *testing.B) {
	// b.Run("go-codec", makeBenchmarkDecodeBeatsEvents(gocodecJSONDecoder))
	b.Run("structform", makeBenchmarkDecodeBeatsEvents(structformJSONDecoder(0)))
	b.Run("structform-keycache", makeBenchmarkDecodeBeatsEvents(structformJSONDecoder(1000)))
	b.Run("std-json", makeBenchmarkDecodeBeatsEvents(stdJSONDecoder))

	// fails with panic
	// b.Run("jsoniter", makeBenchmarkDecodeBeatsEvents(jsoniterDecoder))
}

func BenchmarkDecodeBeatsEventsMem(b *testing.B) {
	b.Run("structform", makeBenchmarkDecodeBeatsEventsBuffered(structformJSONBufDecoder(0)))
	b.Run("structform-keycache", makeBenchmarkDecodeBeatsEventsBuffered(structformJSONBufDecoder(1000)))
	b.Run("std-json", makeBenchmarkDecodeBeatsEventsBuffered(stdJSONBufDecoder))

	// fails with panic
	// b.Run("jsoniter", makeBenchmarkDecodeBeatsEventsBuffered(jsoniterBufDecoder))
}

func BenchmarkEncodeBeatsEvents(b *testing.B) {
	events := loadBeatsEvents()
	b.Run("std-json", makeBenchmarkEncodeEvents(stdJSONEncoder, events))
	b.Run("structform-json", makeBenchmarkEncodeEvents(structformJSONEncoder, events))
	b.Run("structform-ubjson", makeBenchmarkEncodeEvents(structformUBJSONEncoder, events))
	b.Run("structform-cborl", makeBenchmarkEncodeEvents(structformCBORLEncoder, events))
}

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
	iter := jsoniter.Parse(r, 4096)
	return func(v interface{}) error {
		iter.ReadVal(v)
		return iter.Error
	}
}

func jsoniterBufDecoder(b []byte) func(interface{}) error {
	iter := jsoniter.ParseBytes(b)
	return func(v interface{}) error {
		iter.ReadVal(v)
		return iter.Error
	}
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

func structformJSONDecoder(keyCache int) func(io.Reader) func(interface{}) error {
	return func(r io.Reader) func(interface{}) error {
		u, _ := gotype.NewUnfolder(nil)
		dec := json.NewDecoder(r, 2*4096, u)
		return makeStructformJSONDecoder(u, dec, keyCache)
	}
}

func structformJSONBufDecoder(keyCache int) func([]byte) func(interface{}) error {
	return func(b []byte) func(interface{}) error {
		u, _ := gotype.NewUnfolder(nil)
		dec := json.NewBytesDecoder(b, u)
		return makeStructformJSONDecoder(u, dec, keyCache)
	}
}

func makeStructformJSONDecoder(
	u *gotype.Unfolder,
	d *json.Decoder,
	keyCache int,
) func(interface{}) error {
	if keyCache > 0 {
		u.EnableKeyCache(keyCache)
	}
	return func(v interface{}) error {
		if err := u.SetTarget(v); err != nil {
			return err
		}
		return d.Next()
	}
}

func makeBenchmarkDecodeBeatsEvents(factory decoderFactory) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				file, err := os.Open("files/beats_events.json")
				// file, err := os.Open("test.json")
				if err != nil {
					b.Error(err)
					return
				}

				defer file.Close()

				decode := factory(file)
				for {
					var to map[string]interface{}
					if err := decode(&to); err != nil {
						if err != io.EOF {
							b.Error(err)
						}
						return
					}
				}
			}()
		}
	}
}

func makeBenchmarkDecodeBeatsEventsBuffered(factory decoderFactoryBuf) func(*testing.B) {
	content, err := ioutil.ReadFile("files/beats_events.json")
	if err != nil {
		panic(err)
	}

	return func(b *testing.B) {
		b.SetBytes(int64(len(content)))

		for i := 0; i < b.N; i++ {
			decode := factory(content)
			for {
				var to map[string]interface{}
				if err := decode(&to); err != nil {
					if err != io.EOF {
						b.Error(err)
					}
					break
				}
			}
		}
	}
}

func makeBenchmarkEncodeEvents(factory encoderFactory, events []map[string]interface{}) func(*testing.B) {
	var buf bytes.Buffer
	buf.Grow(16 * 1024)
	encode := factory(&buf)

	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var written int64

			for _, event := range events {
				buf.Reset()
				if err := encode(event); err != nil {
					b.Error(err)
					return
				}
				written += int64(buf.Len())
			}
			b.SetBytes(written)
		}
	}
}

func loadBeatsEvents() []map[string]interface{} {
	content, err := ioutil.ReadFile("files/beats_events.json")
	if err != nil {
		panic(err)
	}

	var events []map[string]interface{}
	dec := stdjson.NewDecoder(bytes.NewReader(content))
	for {
		var e map[string]interface{}
		if err := dec.Decode(&e); err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		events = append(events, e)
	}

	return events
}
