/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package libs

import (
    "errors"
    "github.com/gin-gonic/gin"
    "strings"
)

// GetToken
func GetToken(c *gin.Context) (string, error) {
    var reqToken = c.Request.Header.Get("Authorization")
    var splitToken = strings.Split(reqToken, "Bearer")

    if len(splitToken) != 2 {
        return "", errors.New("Token not provided.")
    }

    return strings.TrimSpace(splitToken[1]), nil
}
