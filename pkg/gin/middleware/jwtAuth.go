// Package middleware is gin middleware plugin.
package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/errcode"
	"github.com/go-dev-frame/sponge/pkg/jwt"
)

// HeaderAuthorizationKey http header authorization key, value is "Bearer token"
const HeaderAuthorizationKey = "Authorization"

// ExtraVerifyFn extra verify function
type ExtraVerifyFn = func(claims *jwt.Claims, c *gin.Context) error

// AuthOption set the auth options.
type AuthOption func(*authOptions)

type authOptions struct {
	signKey           []byte // sign key for jwt
	isReturnErrReason bool
	extraVerifyFn     ExtraVerifyFn
}

func defaultAuthOptions() *authOptions { _ = "STUB: not implemented"; return nil }

func (o *authOptions) apply(opts ...AuthOption) { _ = "STUB: not implemented"; return }

// WithSignKey set jwt sign key
func WithSignKey(key []byte) AuthOption { _ = "STUB: not implemented"; return *new(AuthOption) }

// WithReturnErrReason set return error reason
func WithReturnErrReason() AuthOption { _ = "STUB: not implemented"; return *new(AuthOption) }

// WithExtraVerify set extra verify function
func WithExtraVerify(fn ExtraVerifyFn) AuthOption {
	_ = "STUB: not implemented"
	return *new(AuthOption)
}

// WithVerify alias of WithExtraVerify
var WithVerify = WithExtraVerify

func responseUnauthorized(isReturnErrReason bool, errMsg string) *errcode.Error {
	_ = "STUB: not implemented"
	return nil
}

// -------------------------------------------------------------------------------------------

// Auth authorization middleware, support custom extra verify.
func Auth(opts ...AuthOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// remove Bearer prefix

// extra verify function

// GetClaims get jwt claims from gin context.
func GetClaims(c *gin.Context) (*jwt.Claims, bool) { _ = "STUB: not implemented"; return nil, false }
