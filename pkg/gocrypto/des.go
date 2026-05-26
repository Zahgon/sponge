// Symmetric encryption DES, one of the most popular encryption algorithms,
// is gradually being replaced by AES.

package gocrypto

// DesEncrypt des encryption, the returned ciphertext is not transcoded
func DesEncrypt(rawData []byte, opts ...DesOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DesDecrypt des decryption, parameter input untranscoded cipher text
func DesDecrypt(cipherData []byte, opts ...DesOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DesEncryptHex des encrypts and returns a ciphertext that has been transcoded
func DesEncryptHex(rawData string, opts ...DesOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DesDecryptHex des decryption, parameter input has been transcoded ciphertext string
func DesDecryptHex(cipherStr string, opts ...DesOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func desEncryptByMode(mode string, rawData []byte, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func desDecryptByMode(mode string, cipherData []byte, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
