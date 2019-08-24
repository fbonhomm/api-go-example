/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package routers

import (
    "github.com/fbonhomm/api-go/source/controllers"
    "github.com/fbonhomm/api-go/source/validators"
    "github.com/gin-gonic/gin"
)

// main
func User(router *gin.Engine) {
    route := router.Group("/users")
    {
        route.POST("/", validators.ValidateUserCreate, controllers.UserCreate)
        route.GET("/:id", validators.ValidateUserGetId, controllers.UserGetId)
        // route.GET("/", fetchAllTodo)
        // route.PUT("/:id", updateTodo)
        // route.DELETE("/:id", deleteTodo)
    }
}
