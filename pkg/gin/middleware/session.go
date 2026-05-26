package middleware

import (
	"github.com/gin-gonic/gin"
)

// 1. Universal session middleware example refer to https://github.com/gin-contrib/sessions?tab=readme-ov-file#basic-examples

// -------------------------------------------------------------------------------------------

// 2. Special session for rails

// RailsCookieAuthMiddleware validates and decrypts a Rails encrypted cookie,
// attaches the session payload to context under key "rails_session".
func RailsCookieAuthMiddleware(secretKeyBase string, cookieName string) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}
