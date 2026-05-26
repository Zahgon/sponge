package errcode

// HCode Generate an error code between 200000 and 300000 according to the number
//
// http service level error code, Err prefix, example.
//
// var (
// ErrUserCreate = NewError(HCode(1)+1, "failed to create user")		// 200101
// ErrUserDelete = NewError(HCode(1)+2, "failed to delete user")		// 200102
// ErrUserUpdate = NewError(HCode(1)+3, "failed to update user")		// 200103
// ErrUserGet    = NewError(HCode(1)+4, "failed to get user details")	// 200104
// )
func HCode(num int) int { _ = "STUB: not implemented"; return 0 }
