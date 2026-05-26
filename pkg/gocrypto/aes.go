// Symmetric encryption AES, the advanced encryption standard with
// the highest level of security, has gradually replaced DES as the new
// generation of symmetric encryption standard.

package gocrypto

import (
	"github.com/go-dev-frame/sponge/pkg/gocrypto/wcipher"
)

// AesEncrypt aes encryption, returns ciphertext is not transcoded
func AesEncrypt(rawData []byte, opts ...AesOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AesDecrypt aes decryption, parameter input un-transcode cipher text
func AesDecrypt(cipherData []byte, opts ...AesOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AesEncryptHex aes encryption, the returned ciphertext is transcoded
func AesEncryptHex(rawData string, opts ...AesOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AesDecryptHex aes decryption, parameter input has been transcoded ciphertext string
func AesDecryptHex(cipherStr string, opts ...AesOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCipherMode(mode string) (wcipher.CipherMode, error) {
	_ = "STUB: not implemented"
	return *new(wcipher.CipherMode), nil
}

func aesEncryptByMode(mode string, rawData []byte, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func aesDecryptByMode(mode string, cipherData []byte, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
