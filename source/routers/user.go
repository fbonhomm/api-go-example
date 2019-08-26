/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package routers

import (
    "github.com/fbonhomm/api-go/source/controllers"
    "github.com/fbonhomm/api-go/source/middlewares"
    "github.com/fbonhomm/api-go/source/validators"
    "github.com/gin-gonic/gin"
)

// main
func User(router *gin.Engine) {
    route := router.Group("/users")

    route.POST("/", middlewares.Auth, validators.ValidateUserCreate, controllers.UserCreate)
    route.GET("/:id", middlewares.Auth, validators.ValidateUserGetId, controllers.UserGetId)
    route.DELETE("/:id", middlewares.Auth, validators.ValidateUserDeleteId, controllers.UserDeleteId)
}
