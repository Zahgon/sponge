package auth

// 1. Universal session management is implemented using the library available at https://github.com/gin-contrib/sessions,
// which provides Gin middleware for session management with support:
//
// cookie-based
// Redis
// memcached
// MongoDB
// GORM
// memstore
// PostgreSQL
// Filesystem

// -------------------------------------------------------------------------------------------

// 2. Special session for rails

// DecodeSignedCookie decrypts a Rails 7.1+ encrypted cookie using the provided
// secretKeyBase and validates that its purpose matches the given cookieName.
// It returns the decoded session payload (the JSON contained in _rails.message).
//
// The Rails encrypted cookie format is: base64(data)--base64(iv)--base64(authTag)
// Key derivation: PBKDF2-HMAC-SHA256(secret_key_base, "authenticated encrypted cookie", 1000, 32)
// Cipher: AES-256-GCM, AAD: empty
func DecodeSignedCookie(secretKeyBase string, decodedCookie string, cookieName string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GCM tag size (bytes)

// Derive key

// AES-256

// AES-GCM decrypt

// In Go, GCM expects ciphertext || tag

// Parse envelope

// Decode inner message (base64 JSON)

// UserIDFromSession tries to extract the warden user id from a Rails session.
// It returns the id and true if found, otherwise (nil, false).
func UserIDFromSession(session map[string]any) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Expecting [[id], ...]
