package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-dev-frame/sponge/pkg/jwt"
)

// ---------------------------------- client ----------------------------------

// SetJwtTokenToCtx set the token (excluding prefix Bearer) to the context in grpc client side
// Example:
//
// authorization := "Bearer jwt-token"
//
//	ctx := SetJwtTokenToCtx(ctx, jwt-token)
//	cli.GetByID(ctx, req)
func SetJwtTokenToCtx(ctx context.Context, token string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetAuthToCtx set the authorization (including prefix Bearer) to the context in grpc client side
// Example:
//
//	ctx := SetAuthToCtx(ctx, authorization)
//	cli.GetByID(ctx, req)
func SetAuthToCtx(ctx context.Context, authorization string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ---------------------------------- server interceptor ----------------------------------

var (
	headerAuthorize = "authorization"

	// auth Scheme
	authScheme = "Bearer"

	// authentication information in ctx key name
	authCtxClaimsName = "tokenInfo"

	// collection of skip authentication methods
	authIgnoreMethods = map[string]struct{}{}
)

// GetAuthorization combining tokens into authentication information
func GetAuthorization(token string) string { _ = "STUB: not implemented"; return "" }

// GetAuthCtxKey get the name of Claims
func GetAuthCtxKey() string { _ = "STUB: not implemented"; return "" }

// ExtraVerifyFn extra verify function
type ExtraVerifyFn = func(ctx context.Context, claims *jwt.Claims) error

// AuthOption setting the Authentication Field
type AuthOption func(*authOptions)

// authOptions settings
type authOptions struct {
	authScheme    string
	ctxClaimsName string
	ignoreMethods map[string]struct{}

	signKey       []byte // sign key for jwt
	extraVerifyFn ExtraVerifyFn
}

func defaultAuthOptions() *authOptions { _ = "STUB: not implemented"; return nil }

// ways to ignore forensics

func (o *authOptions) apply(opts ...AuthOption) { _ = "STUB: not implemented"; return }

// WithAuthScheme set the message prefix for authentication
func WithAuthScheme(scheme string) AuthOption { _ = "STUB: not implemented"; return *new(AuthOption) }

// WithAuthClaimsName set the key name of the information in ctx for authentication
func WithAuthClaimsName(claimsName string) AuthOption {
	_ = "STUB: not implemented"
	return *new(AuthOption)
}

// WithAuthIgnoreMethods ways to ignore forensics
// fullMethodName format: /packageName.serviceName/methodName,
// example /api.userExample.v1.userExampleService/GetByID
func WithAuthIgnoreMethods(fullMethodNames ...string) AuthOption {
	_ = "STUB: not implemented"
	return *new(AuthOption)
}

// WithSignKey set jwt sign key
func WithSignKey(key []byte) AuthOption { _ = "STUB: not implemented"; return *new(AuthOption) }

// WithExtraVerify set extra verify function
func WithExtraVerify(fn ExtraVerifyFn) AuthOption {
	_ = "STUB: not implemented"
	return *new(AuthOption)
}

// -------------------------------------------------------------------------------------------

// verify authorization from context, support default and custom verify processing
func jwtVerify(ctx context.Context, opt *authOptions) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// key is authScheme

//nolint

// GetJwtClaims get the jwt default claims from context, contains fixed fields uid and name
func GetJwtClaims(ctx context.Context) (*jwt.Claims, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// UnaryServerJwtAuth jwt unary interceptor
func UnaryServerJwtAuth(opts ...AuthOption) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerJwtAuth jwt stream interceptor
func StreamServerJwtAuth(opts ...AuthOption) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
