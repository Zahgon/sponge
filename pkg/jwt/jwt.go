package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims universal claims
type Claims struct {
	UID    string                 `json:"uid,omitempty"`    // user id
	Fields map[string]interface{} `json:"fields,omitempty"` // custom fields
	jwt.RegisteredClaims
}

// Get custom field value by key, if not found, return false
func (c *Claims) Get(key string) (val interface{}, isExist bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetString custom field value by key, if not found, return false
func (c *Claims) GetString(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// GetInt custom field value by key, if not found, return false
func (c *Claims) GetInt(key string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// GetInt64 custom field value by key, if not found, return false
func (c *Claims) GetInt64(key string) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

// GetBool custom field value by key, if not found, return false
func (c *Claims) GetBool(key string) (b bool, isExist bool) {
	_ = "STUB: not implemented"
	return false, false
}

// GetFloat64 custom field value by key, if not found, return false
func (c *Claims) GetFloat64(key string) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

// NewToken create new token with claims, duration, signing method and signing key
func (c *Claims) NewToken(d time.Duration, signMethod jwt.SigningMethod, signKey []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetClaimsUnverified get claims from token, not verifying signature
func GetClaimsUnverified(tokenString string) (*Claims, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ------------------------------- one token -------------------------------

// GenerateToken create token by uid and name, use universal Claims
func GenerateToken(uid string, opts ...GenerateTokenOption) (jwtID string, tokenStr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ValidateToken validate token, return error if token is invalid
func ValidateToken(tokenString string, opts ...ValidateTokenOption) (*Claims, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifyToken(tokenString string, opts ...ValidateTokenOption) (string, *Claims, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// RefreshToken refresh token
func RefreshToken(tokenString string, opts ...RefreshTokenOption) (jwtID string, tokenStr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// --------------------------------- two tokens ---------------------------------

type Tokens struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	JwtID        string `json:"jwtID"` // used to prevent replay attacks, identifying specific tokens
}

// GenerateTwoTokens create accessToken and refreshToken
func GenerateTwoTokens(uid string, opts ...GenerateTwoTokensOption) (*Tokens, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// forced id consistency

// RefreshTwoTokens refresh access token, if refresh token is expired time is less than 3 hours, will auto refresh token too.
// if return err is ErrTokenExpired, you need to login again to get token.
func RefreshTwoTokens(refreshToken string, accessToken string, opts ...RefreshTwoTokensOption) (*Tokens, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
