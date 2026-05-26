package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		userExampleRouter(group, handler.NewUserExampleHandler())
	})
}

func userExampleRouter(group *gin.RouterGroup, h handler.UserExampleHandler) {
	_ = "STUB: not implemented"
	return
}

// JWT authentication reference: https://go-sponge.com/component/transport/gin.html#jwt-authorization-middleware

// All the following routes use jwt authentication, you also can use middleware.Auth(middleware.WithExtraVerify(fn))
//g.Use(middleware.Auth())

// If jwt authentication is not required for all routes, authentication middleware can be added
// separately for only certain routes. In this case, g.Use(middleware.Auth()) above should not be used.

// [post] /api/v1/userExample
// [delete] /api/v1/userExample/:id
// [put] /api/v1/userExample/:id
// [get] /api/v1/userExample/:id
// [post] /api/v1/userExample/list
