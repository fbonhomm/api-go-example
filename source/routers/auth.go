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
func Auth(router *gin.Engine) {
    route := router.Group("/auth")

    route.POST("/", validators.ValidateAuthLogin, controllers.AuthLogin)
    route.POST("/refresh", middlewares.AuthRefresh, validators.ValidateAuthRefresh, controllers.AuthRefresh)
}
