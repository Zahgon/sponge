package wcipher

import (
	"crypto/cipher"
)

// CipherMode provides a uniform interface to set the filling method for different operating modes.
type CipherMode interface {
	SetPadding(padding Padding) CipherMode
	Cipher(block cipher.Block, iv []byte) Cipher
}

type cipherMode struct {
	padding Padding
}

// SetPadding set padding
func (c *cipherMode) SetPadding(padding Padding) CipherMode {
	_ = "STUB: not implemented"
	return *

	// Cipher mode cipher
	new(CipherMode)
}

func (c *cipherMode) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

type ecbCipherModel cipherMode

// NewECBMode new ecb mode
func NewECBMode() CipherMode { _ = "STUB: not implemented"; return *new(CipherMode) }

// SetPadding set ecb padding
func (ecb *ecbCipherModel) SetPadding(padding Padding) CipherMode {
	_ = "STUB: not implemented"
	return *new(CipherMode)
}

// Cipher ecb cipher
func (ecb *ecbCipherModel) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

type cbcCipherModel cipherMode

// NewCBCMode new cbc mode
func NewCBCMode() CipherMode { _ = "STUB: not implemented"; return *new(CipherMode) }

// SetPadding set cbc padding
func (cbc *cbcCipherModel) SetPadding(padding Padding) CipherMode {
	_ = "STUB: not implemented"
	return *new(CipherMode)
}

// Cipher cbc cipher
func (cbc *cbcCipherModel) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

type cfbCipherModel cipherMode //nolint

// NewCFBMode new cfb mode
func NewCFBMode() CipherMode {
	_ = "STUB: not implemented"
	return *

	// Cipher cfb cipher
	new(CipherMode)
}

func (cfb *cfbCipherModel) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented" //nolint
	return *new(Cipher)
}

type ofbCipherModel struct {
	cipherMode
}

// NewOFBMode new ofb mode
func NewOFBMode() CipherMode {
	_ = "STUB: not implemented"
	return *

	// Cipher ofb cipher
	new(CipherMode)
}

func (ofb *ofbCipherModel) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

type ctrCipherModel struct {
	cipherMode
}

// NewCTRMode new ctr mode
func NewCTRMode() CipherMode {
	_ = "STUB: not implemented"
	return *

	// Cipher ctr cipher
	new(CipherMode)
}

func (ctr *ctrCipherModel) Cipher(block cipher.Block, iv []byte) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}
