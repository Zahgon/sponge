// Package wcipher is a package to encrypt and decrypt data.
package wcipher

import (
	"crypto/cipher"
)

// Cipher provides a unified interface to encrypt/decrypt data.
type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

type blockCipher struct {
	padding Padding
	encrypt cipher.BlockMode
	decrypt cipher.BlockMode
}

// NewBlockCipher new block encryption
func NewBlockCipher(padding Padding, encrypt, decrypt cipher.BlockMode) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

// Encrypt encrypted
func (blockCipher *blockCipher) Encrypt(plaintext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Decrypt decrypt
func (blockCipher *blockCipher) Decrypt(ciphertext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// ---------------------------------------------------------------------------------------

type streamCipher struct {
	encrypt cipher.Stream
	decrypt cipher.Stream
}

// NewStreamCipher new stream encryption
func NewStreamCipher(encrypt cipher.Stream, decrypt cipher.Stream) Cipher {
	_ = "STUB: not implemented"
	return *new(Cipher)
}

// Encrypt stream encryption
func (streamCipher *streamCipher) Encrypt(plaintext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Decrypt stream decryption
func (streamCipher *streamCipher) Decrypt(ciphertext []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
