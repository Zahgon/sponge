package encoding

// GobEncoding gob encode
type GobEncoding struct{}

// Marshal gob encode
func (g GobEncoding) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal gob encode
func (g GobEncoding) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
