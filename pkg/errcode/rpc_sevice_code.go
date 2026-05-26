package errcode

import "google.golang.org/grpc/codes"

// RCode Generate an error code between 400000 and 500000 according to the number
//
// rpc service level error code, status prefix, example.
//
//	var (
//		StatusUserCreate = NewRPCStatus(RCode(1)+1, "failed to create user")		// 400101
//		StatusUserDelete = NewRPCStatus(RCode(1)+2, "failed to delete user")		// 400102
//		StatusUserUpdate = NewRPCStatus(RCode(1)+3, "failed to update user")		// 400103
//		StatusUserGet    = NewRPCStatus(RCode(1)+4, "failed to get user details")	// 400104
//	)
func RCode(num int) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }
