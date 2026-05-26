package gocrypto

import (
	"crypto"
)

var hashKey = []byte("fVy7UjMkO9_pLqs3")

// Md5 hash
func Md5(rawData []byte) string { _ = "STUB: not implemented"; return "" }

// Sha1 hash
func Sha1(rawData []byte) string { _ = "STUB: not implemented"; return "" }

// Sha256 hash
func Sha256(rawData []byte) string { _ = "STUB: not implemented"; return "" }

// Sha512 hash
func Sha512(rawData []byte) string { _ = "STUB: not implemented"; return "" }

func sha1Hash(slices [][]byte) []byte { _ = "STUB: not implemented"; return nil }

func md5Sha1(slices [][]byte) string { _ = "STUB: not implemented"; return "" }

// Hash commonly used hash sets
func Hash(hashType crypto.Hash, rawData []byte) (string, error) {
	_ = "STUB: not implemented" //nolint
	return "", nil
}
