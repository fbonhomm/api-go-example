/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/fbonhomm/api-go/source/controllers"
	"github.com/fbonhomm/api-go/source/middlewares"
	"github.com/fbonhomm/api-go/source/validators"
)

// main
func User(router *gin.Engine) {
	route := router.Group("/users")

	route.POST("", middlewares.Auth, validators.ValidateUserCreate, controllers.UserCreate)
	route.GET("/:id", middlewares.Auth, validators.ValidateUserGetID, controllers.UserGetID)
	route.DELETE("/:id", middlewares.Auth, validators.ValidateUserDeleteID, controllers.UserDeleteID)
}
