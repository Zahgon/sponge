// Package proto is a protobuf encoding and decoding.
package proto

import (
	"github.com/go-dev-frame/sponge/pkg/encoding"
)

// Name is the name registered for the proto compressor.
const Name = "proto"

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with protobuf. It is the default codec for gRPC.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (codec) Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func (codec) Name() string { _ = "STUB: not implemented"; return "" }
