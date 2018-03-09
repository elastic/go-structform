package sftest

import (
	"testing"

	structform "github.com/elastic/go-structform"
)

func TestRecordingConsistent(t *testing.T) {
	TestEncodeParseConsistent(t, Samples,
		func() (structform.Visitor, func(structform.Visitor) error) {
			buf := &Recording{}
			return buf, buf.Replay
		})
}
