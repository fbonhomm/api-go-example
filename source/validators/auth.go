/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package validators

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
)

type authLogin struct {
    Email       string  `form:"email" binding:"required,email"`
    Password    string  `form:"password" binding:"required,min=8,max=50"`
}


// ValidateAuthLogin
func ValidateAuthLogin(c *gin.Context) {
    var v authLogin

    if err := c.ShouldBindWith(&v, binding.FormPost); err != nil {
        errorHandling(c, err.Error())
    }
}
