package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SigningMethodHMAC = jwt.SigningMethodHMAC

var (
	HS256 = jwt.SigningMethodHS256
	HS384 = jwt.SigningMethodHS384
	HS512 = jwt.SigningMethodHS512
)

var (
	defaultSigningKey    = []byte("CaqGzKLUsmWWbWI6F5EZbLwHsQeJ5RLyYTwBqa3mDKY6") // default key
	defaultSigningMethod = HS256                                                  // default HS256
	defaultExpire        = 24 * time.Hour                                         // default expiration one day
)

var (
	ErrTokenExpired = jwt.ErrTokenExpired
	//errInvalid      = errors.New("token is invalid")
	errClaims   = errors.New("claims is not match")
	errNotMatch = errors.New(" access token and refresh token is not match")
)

// ------------------------------------------------------------------------------------------

type registeredClaimsOptions struct {
	registeredClaims jwt.RegisteredClaims
}

func defaultRegisteredClaimsOptions(expire time.Duration, id string) *registeredClaimsOptions {
	_ = "STUB: not implemented"
	return nil
}

// RegisteredClaimsOption set the registered claims options.
type RegisteredClaimsOption func(*registeredClaimsOptions)

func (o *registeredClaimsOptions) apply(opts ...RegisteredClaimsOption) {
	_ = "STUB: not implemented"
	return
}

// WithIssuer set issuer (iss) value
func WithIssuer(issuer string) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithSubject set subject (sub) value
func WithSubject(subject string) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithAudience set audience (aud) value
func WithAudience(audience ...string) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithExpires set expires (exp) value
func WithExpires(d time.Duration) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithDeadline set expires (exp) value
func WithDeadline(expiresAt time.Time) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithNotBefore set not before (nbf) value
func WithNotBefore(notBefore time.Time) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithIssuedAt set issued at (iat) value
func WithIssuedAt(issuedAt time.Time) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// WithJwtID set jwt id (jti) value
func WithJwtID(id string) RegisteredClaimsOption {
	_ = "STUB: not implemented"
	return *new(RegisteredClaimsOption)
}

// -------------------------------------------------------------------------------

type generateTokenOptions struct {
	signKey    []byte
	signMethod jwt.SigningMethod

	fields map[string]interface{} // custom fields

	tokenClaimsOptions *registeredClaimsOptions
}

func defaultGenerateTokenOptions() *generateTokenOptions { _ = "STUB: not implemented"; return nil }

// GenerateTokenOption set the jwt options.
type GenerateTokenOption func(*generateTokenOptions)

func (o *generateTokenOptions) apply(opts ...GenerateTokenOption) {
	_ = "STUB: not implemented"
	return
}

// WithGenerateTokenSignMethod set sign method value
func WithGenerateTokenSignMethod(sm jwt.SigningMethod) GenerateTokenOption {
	_ = "STUB: not implemented"
	return *new(GenerateTokenOption)
}

// WithGenerateTokenSignKey set sign key value
func WithGenerateTokenSignKey(key []byte) GenerateTokenOption {
	_ = "STUB: not implemented"
	return *new(GenerateTokenOption)
}

// WithGenerateTokenFields set custom fields value
func WithGenerateTokenFields(fields map[string]interface{}) GenerateTokenOption {
	_ = "STUB: not implemented"
	return *new(GenerateTokenOption)
}

// WithGenerateTokenClaims set token claims value
func WithGenerateTokenClaims(opts ...RegisteredClaimsOption) GenerateTokenOption {
	_ = "STUB: not implemented"
	return *new(GenerateTokenOption)
}

// ------------------------------------------------------------------------------------

type validateTokenOptions struct {
	signKey []byte
}

func defaultValidateTokenOptions() *validateTokenOptions { _ = "STUB: not implemented"; return nil }

// ValidateTokenOption set parse token options.
type ValidateTokenOption func(*validateTokenOptions)

func (o *validateTokenOptions) apply(opts ...ValidateTokenOption) {
	_ = "STUB: not implemented"
	return
}

// WithValidateTokenSignKey set sign key value
func WithValidateTokenSignKey(key []byte) ValidateTokenOption {
	_ = "STUB: not implemented"
	return *new(ValidateTokenOption)
}

// ------------------------------------------------------------------------------

type refreshTokenOptions struct {
	signKey []byte
	expire  time.Duration
}

func defaultRefreshTokenOptions() *refreshTokenOptions { _ = "STUB: not implemented"; return nil }

// RefreshTokenOption set refresh token options.
type RefreshTokenOption func(*refreshTokenOptions)

func (o *refreshTokenOptions) apply(opts ...RefreshTokenOption) { _ = "STUB: not implemented"; return }

// WithRefreshTokenSignKey set sign key value
func WithRefreshTokenSignKey(key []byte) RefreshTokenOption {
	_ = "STUB: not implemented"
	return *new(RefreshTokenOption)
}

// WithRefreshTokenExpire set expire value
func WithRefreshTokenExpire(expire time.Duration) RefreshTokenOption {
	_ = "STUB: not implemented"
	return *new(RefreshTokenOption)
}

// ------------------------------------------------------------------------------------------

type generateTwoTokensOptions struct {
	signMethod jwt.SigningMethod
	signKey    []byte

	fields map[string]interface{} // custom fields

	accessTokenClaimsOptions  *registeredClaimsOptions
	refreshTokenClaimsOptions *registeredClaimsOptions
}

func defaultGenerateTwoTokensOptions() *generateTwoTokensOptions {
	_ = "STUB: not implemented"
	return nil
}

// 30 minutes
// 30 days

// GenerateTwoTokensOption set the jwt options.
type GenerateTwoTokensOption func(*generateTwoTokensOptions)

func (o *generateTwoTokensOptions) apply(opts ...GenerateTwoTokensOption) {
	_ = "STUB: not implemented"
	return
}

// WithGenerateTwoTokensSignMethod set sign method value
func WithGenerateTwoTokensSignMethod(sm jwt.SigningMethod) GenerateTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(GenerateTwoTokensOption)
}

// WithGenerateTwoTokensSignKey set sign key value
func WithGenerateTwoTokensSignKey(key []byte) GenerateTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(GenerateTwoTokensOption)
}

// WithGenerateTwoTokensFields set custom fields value
func WithGenerateTwoTokensFields(fields map[string]interface{}) GenerateTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(GenerateTwoTokensOption)
}

// WithGenerateTwoTokensAccessTokenClaims set Access token claims value
func WithGenerateTwoTokensAccessTokenClaims(opts ...RegisteredClaimsOption) GenerateTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(GenerateTwoTokensOption)
}

// WithGenerateTwoTokensRefreshTokenClaims set refresh token claims value
func WithGenerateTwoTokensRefreshTokenClaims(opts ...RegisteredClaimsOption) GenerateTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(GenerateTwoTokensOption)
}

// -------------------------------------------------------------------------------------

type refreshTwoTokensOptions struct {
	signKey            []byte
	accessTokenExpire  time.Duration
	refreshTokenExpire time.Duration
}

func defaultRefreshTwoTokensOptions() *refreshTwoTokensOptions {
	_ = "STUB: not implemented"
	return nil
}

// 30 minutes
// 30 days

// RefreshTwoTokensOption set refresh token options.
type RefreshTwoTokensOption func(*refreshTwoTokensOptions)

func (o *refreshTwoTokensOptions) apply(opts ...RefreshTwoTokensOption) {
	_ = "STUB: not implemented"
	return
}

// WithRefreshTwoTokensSignKey set sign key value
func WithRefreshTwoTokensSignKey(key []byte) RefreshTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(RefreshTwoTokensOption)
}

// WithRefreshTwoTokensRefreshTokenExpires set refresh token expire value
func WithRefreshTwoTokensRefreshTokenExpires(d time.Duration) RefreshTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(RefreshTwoTokensOption)
}

// WithRefreshTwoTokensAccessTokenExpires set access token expire value
func WithRefreshTwoTokensAccessTokenExpires(d time.Duration) RefreshTwoTokensOption {
	_ = "STUB: not implemented"
	return *new(RefreshTwoTokensOption)
}

func getAlg(alg string) (jwt.SigningMethod, error) {
	_ = "STUB: not implemented"
	return *new(jwt.SigningMethod), nil
}
