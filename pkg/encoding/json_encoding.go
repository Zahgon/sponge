package encoding

// JSONEncoding json format
type JSONEncoding struct{}

// Marshal json encode
func (j JSONEncoding) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal json decode
func (j JSONEncoding) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// JSONGzipEncoding json and gzip
type JSONGzipEncoding struct{}

// Marshal json encode and gzip
func (jz JSONGzipEncoding) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// var bufSizeBefore = len(buf)

// log.Infof("gzip_json_compress_ratio=%d/%d=%.2f", bufSizeBefore, len(buf), float64(bufSizeBefore)/float64(len(buf)))

// Unmarshal json encode and gzip
func (jz JSONGzipEncoding) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// GzipEncode encoding
func GzipEncode(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GzipDecode decode
func GzipDecode(in []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// JSONSnappyEncoding json format and snappy compression
type JSONSnappyEncoding struct{}

// Marshal serialization
func (s JSONSnappyEncoding) Marshal(v interface{}) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal deserialization
func (s JSONSnappyEncoding) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
