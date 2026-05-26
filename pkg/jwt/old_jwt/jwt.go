// Package jwt is deprecated, old package path is "github.com/go-dev-frame/sponge/pkg/jwt/old_jwt"
// Please use new jwt package instead, new package path is "github.com/go-dev-frame/sponge/pkg/jwt"
package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// ErrTokenExpired expired
var ErrTokenExpired = jwt.ErrTokenExpired

var opt *options

// Init initialize jwt
// Deprecated: build jwt.Init() before use
func Init(opts ...Option) { _ = "STUB: not implemented"; return }

// Claims standard claims, include uid, name, and RegisteredClaims
type Claims struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken generate token by uid and name, use universal Claims
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" GenerateToken instead
func GenerateToken(uid string, name ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseToken parse token, return universal Claims
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" ValidateToken instead
func ParseToken(tokenString string) (*Claims, error) { _ = "STUB: not implemented"; return nil, nil }

// RefreshToken refresh token
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" RefreshToken instead
func RefreshToken(tokenString string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// -------------------------------------------------------------------------------------------

// KV map type
type KV = map[string]interface{}

// CustomClaims custom fields claims
type CustomClaims struct {
	Fields KV `json:"fields"`
	jwt.RegisteredClaims
}

// Get custom field value by key, if not found, return false
func (c *CustomClaims) Get(key string) (val interface{}, isExist bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetString custom field value by key, if not found, return false
func (c *CustomClaims) GetString(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GetInt custom field value by key, if not found, return false
func (c *CustomClaims) GetInt(key string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// GetUint64 custom field value by key, if not found, return false
func (c *CustomClaims) GetUint64(key string) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// GenerateCustomToken generate token by custom fields, use CustomClaims
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" GenerateToken instead
func GenerateCustomToken(kv map[string]interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseCustomToken parse token, return CustomClaims
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" ValidateToken instead
func ParseCustomToken(tokenString string) (*CustomClaims, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RefreshCustomToken refresh custom token
// Deprecated: use "github.com/go-dev-frame/sponge/pkg/jwt" RefreshToken instead
func RefreshCustomToken(tokenString string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
