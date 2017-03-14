package ubjson

import (
	"bytes"
	"testing"

	"github.com/urso/go-structform"
	"github.com/urso/go-structform/sftest"
)

func TestEncParseConsistent(t *testing.T) {
	sftest.TestEncodeParseConsistent(t, sftest.Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := bytes.NewBuffer(nil)
			vs := NewVisitor(buf)

			return vs, func(to structform.Visitor) error {
				return Parse(buf.Bytes(), to)
			}
		})
}

func TestEncParseBytesConsistent(t *testing.T) {
	sftest.TestEncodeParseConsistent(t, sftest.Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := bytes.NewBuffer(nil)
			vs := NewVisitor(buf)

			return vs, func(to structform.Visitor) error {
				p := NewParser(to)
				for _, b := range buf.Bytes() {
					err := p.feed([]byte{b})
					if err != nil {
						return err
					}
				}
				return p.finalize()
			}
		})
}
