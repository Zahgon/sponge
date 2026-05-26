// Asymmetric encryption and decryption rsa
// 1. public key encryption, private key decryption to get the original text
// 2. private key signature, public key signature verification

package gocrypto

import (
	"crypto"
	"crypto/rsa"
)

const (
	pkcs1 = "PKCS#1"
	pkcs8 = "PKCS#8"
)

// RsaEncrypt rsa encryption, the returned ciphertext is not transcoded
func RsaEncrypt(publicKey []byte, rawData []byte, opts ...RsaOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RsaDecrypt rsa decryption, parameter input untranscoded cipher text
func RsaDecrypt(privateKey []byte, cipherData []byte, opts ...RsaOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RsaEncryptHex rsa encryption, return hex
func RsaEncryptHex(publicKey []byte, rawData []byte, opts ...RsaOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RsaDecryptHex rsa decryption, return to original
func RsaDecryptHex(privateKey []byte, cipherHex string, opts ...RsaOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RsaSign rsa signature, the returned ciphertext is not transcoded
func RsaSign(privateKey []byte, rawData []byte, opts ...RsaOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RsaVerify rsa signature verification
func RsaVerify(publicKey []byte, rawData []byte, signData []byte, opts ...RsaOption) error {
	_ = "STUB: not implemented"
	return nil
}

// RsaSignBase64 rsa signature, return base64
func RsaSignBase64(privateKey []byte, rawData []byte, opts ...RsaOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RsaVerifyBase64 rsa signature verification
func RsaVerifyBase64(publicKey []byte, rawData []byte, signBase64 string, opts ...RsaOption) error {
	_ = "STUB: not implemented"
	return nil
}

// encrypt with public key
func rsaEncryptWithPublicKey(publicKey []byte, rawData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// decrypt with private key
func rsaDecryptWithPrivateKey(privateKey []byte, cipherData []byte, format string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sign with private key
func rsaSignWithPrivateKey(privateKey []byte, hash crypto.Hash, rawData []byte, format string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verify with public key
func rsaVerifyWithPublicKey(publicKey []byte, hash crypto.Hash, rawData []byte, signData []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func getPrivateKey(der []byte, format string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
