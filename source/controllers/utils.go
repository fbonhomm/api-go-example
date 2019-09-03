/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package controllers

import (
    "errors"
    "github.com/gin-gonic/gin"
)

// GetToken
func GetToken(c *gin.Context) (info map[string]string, err error) {
    token, ok := c.Get("Token")

    if ok == false {
        err = errors.New("token is not found")
    } else {
        info = token.(map[string]string)
    }

    return info, err
}
