// Package auth provides JWT authentication middleware for gin.
package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/errcode"
	"github.com/go-dev-frame/sponge/pkg/jwt"
)

type SigningMethodHMAC = jwt.SigningMethodHMAC
type Claims = jwt.Claims

var (
	HS256 = jwt.HS256
	HS384 = jwt.HS384
	HS512 = jwt.HS512
)

var (
	customSigningKey    []byte
	customSigningMethod *jwt.SigningMethodHMAC
	customExpire        time.Duration
	customIssuer        string

	errOption = errors.New("jwt option is nil, please initialize first, call middleware.InitAuth()")
)

type initAuthOptions struct {
	issuer        string
	signingMethod *SigningMethodHMAC
}

func defaultInitAuthOptions() *initAuthOptions { _ = "STUB: not implemented"; return nil }

// InitAuthOption set the jwt initAuthOptions.
type InitAuthOption func(*initAuthOptions)

func (o *initAuthOptions) apply(opts ...InitAuthOption) { _ = "STUB: not implemented"; return }

// WithInitAuthSigningMethod set signing method value
func WithInitAuthSigningMethod(sm *jwt.SigningMethodHMAC) InitAuthOption {
	_ = "STUB: not implemented"
	return *new(InitAuthOption)
}

// WithInitAuthIssuer set issuer value
func WithInitAuthIssuer(issuer string) InitAuthOption {
	_ = "STUB: not implemented"
	return *new(InitAuthOption)
}

// InitAuth initializes jwt options.
func InitAuth(signingKey []byte, expire time.Duration, opts ...InitAuthOption) {
	_ = "STUB: not implemented"
	return
}

// GenerateTokenOption set the jwt options.
type GenerateTokenOption func(*generateTokenOptions)

type generateTokenOptions struct {
	fields map[string]interface{}
}

func (o *generateTokenOptions) apply(opts ...GenerateTokenOption) {
	_ = "STUB: not implemented"
	return
}

// WithGenerateTokenFields set custom fields value
func WithGenerateTokenFields(fields map[string]interface{}) GenerateTokenOption {
	_ = "STUB: not implemented"
	return *new(GenerateTokenOption)
}

// GenerateToken generates a jwt token with the given uid and options.
func GenerateToken(uid string, opts ...GenerateTokenOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseToken parses the given token and returns the claims.
func ParseToken(token string) (*jwt.Claims, error) { _ = "STUB: not implemented"; return nil, nil }

// RefreshToken create a new token with the given claims.
func RefreshToken(claims *jwt.Claims) (string, error) { _ = "STUB: not implemented"; return "", nil }

// -------------------------------------------------------------------------------------------

// HeaderAuthorizationKey http header authorization key, value is "Bearer token"
const HeaderAuthorizationKey = "Authorization"

// ExtraVerifyFn extra verify function
type ExtraVerifyFn = func(claims *jwt.Claims, c *gin.Context) error

// AuthOption set the auth options.
type AuthOption func(*authOptions)

type authOptions struct {
	isReturnErrReason bool
	extraVerifyFn     ExtraVerifyFn
}

func defaultAuthOptions() *authOptions { _ = "STUB: not implemented"; return nil }

func (o *authOptions) apply(opts ...AuthOption) { _ = "STUB: not implemented"; return }

// WithReturnErrReason set return error reason
func WithReturnErrReason() AuthOption { _ = "STUB: not implemented"; return *new(AuthOption) }

// WithExtraVerify set extra verify function
func WithExtraVerify(fn ExtraVerifyFn) AuthOption {
	_ = "STUB: not implemented"
	return *new(AuthOption)
}

func responseUnauthorized(isReturnErrReason bool, errMsg string) *errcode.Error {
	_ = "STUB: not implemented"
	return nil
}

// Auth authorization middleware, support custom extra verify.
func Auth(opts ...AuthOption) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// remove Bearer prefix

// extra verify function

// set claims to context

// GetClaims get jwt claims from gin context.
func GetClaims(c *gin.Context) (*jwt.Claims, bool) { _ = "STUB: not implemented"; return nil, false }
