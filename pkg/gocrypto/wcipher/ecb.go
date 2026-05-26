package wcipher

import (
	"crypto/cipher"
)

type ecb struct {
	block     cipher.Block
	blockSize int
}

type ecbEncrypt ecb

func (e *ecbEncrypt) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (e *ecbEncrypt) CryptBlocks(dst, src []byte) { _ = "STUB: not implemented"; return }

type ecbDecrypt ecb

func (e *ecbDecrypt) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (e *ecbDecrypt) CryptBlocks(dst, src []byte) { _ = "STUB: not implemented"; return }

// NewECBEncrypt ecb encrypt
func NewECBEncrypt(block cipher.Block) cipher.BlockMode {
	_ = "STUB: not implemented"
	return *new(cipher.BlockMode)
}

// NewECBDecrypt ecb decrypt
func NewECBDecrypt(block cipher.Block) cipher.BlockMode {
	_ = "STUB: not implemented"
	return *new(cipher.BlockMode)
}
