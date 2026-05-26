package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// ---------------------------------- client option ----------------------------------

type authToken struct {
	AppID    string `json:"app_id"`
	AppKey   string `json:"app_key"`
	IsSecure bool   `json:"isSecure"`
}

// GetRequestMetadata get metadata
func (t *authToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	_ = "STUB: not implemented" //nolint
	return nil, nil
}

// RequireTransportSecurity is require transport secure
func (t *authToken) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"

	// ClientTokenOption client token
	return false
}

func ClientTokenOption(appID string, appKey string, isSecure bool) grpc.DialOption {
	_ = "STUB: not implemented"
	return *new(grpc.DialOption)
}

// ---------------------------------- server interceptor ----------------------------------

// CheckToken check app id and app key
// Example:
//
//	var f CheckToken=func(appID string, appKey string) error{
//		if appID != targetAppID || appKey != targetAppKey {
//			return status.Errorf(codes.Unauthenticated, "app id or app key checksum failure")
//		}
//		return nil
//	}
type CheckToken func(appID string, appKey string) error

// UnaryServerToken recovery unary token
func UnaryServerToken(f CheckToken) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// StreamServerToken recovery stream token
func StreamServerToken(f CheckToken) grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}
