package encoding

// MsgPackEncoding msgpack format
type MsgPackEncoding struct{}

// Marshal msgpack encode
func (mp MsgPackEncoding) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal msgpack decode
func (mp MsgPackEncoding) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
