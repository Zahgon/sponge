// Package json is a JSON encoding and decoding.
package json

import (
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/go-dev-frame/sponge/pkg/encoding"
)

// Name is the name registered for the json codec.
const Name = "json"

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with json.
type codec struct{}

// Marshal object to data
func (codec) Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal data to bytes
func (codec) Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

// Name get name
func (codec) Name() string { _ = "STUB: not implemented"; return "" }
