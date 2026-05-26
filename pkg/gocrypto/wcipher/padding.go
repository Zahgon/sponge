package wcipher

// Padding provides a unified interface to populate/restore data for various population methods.
type Padding interface {
	Padding(src []byte, blockSize int) []byte
	UnPadding(src []byte) []byte
}

type padding struct{}

type pkcs57Padding padding

// NewPKCS57Padding new pkcs57 padding
func NewPKCS57Padding() Padding {
	_ = "STUB: not implemented"
	return *

	// Padding pkcs57 padding
	new(Padding)
}

func (p *pkcs57Padding) Padding(src []byte, blockSize int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// UnPadding pkcs57 un-padding
func (p *pkcs57Padding) UnPadding(src []byte) []byte { _ = "STUB: not implemented"; return nil }
