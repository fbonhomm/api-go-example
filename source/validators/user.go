/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package validators

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "net/http"
)

// https://godoc.org/gopkg.in/go-playground/validator.v8

type userCreate struct {
    Name        string  `form:"name" binding:"required,min=2,max=50"`
    Email       string  `form:"email" binding:"required,email"`
    Password    string  `form:"password" binding:"required,min=8,max=50"`
}

type userGetId struct {
    Id  string  `uri:"id" binding:"required,numeric,min=1,max=5"`
}

func errorHandling(c *gin.Context, msg string) {
    c.JSON(http.StatusBadRequest, gin.H{ "error": msg })
    c.Abort()
}


// ValidateUserCreate
func ValidateUserCreate(c *gin.Context) {
    var v userCreate

    if err := c.ShouldBindWith(&v, binding.FormPost); err != nil {
        errorHandling(c, err.Error())
    }
}

// ValidateUserGetId
func ValidateUserGetId(c *gin.Context) {
    var v userGetId

    if err := c.ShouldBindUri(&v); err != nil {
        errorHandling(c, err.Error())
    }
}
