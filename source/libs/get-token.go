/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package libs

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

// GetToken
func GetToken(c *gin.Context) string {
    var reqToken = c.Request.Header.Get("Authorization")
    var splitToken = strings.Split(reqToken, "Bearer")

    if len(splitToken) != 2 {
        c.JSON(http.StatusUnauthorized, gin.H{ "error": "Token not conform." })
        c.Abort()
    }

    return strings.TrimSpace(splitToken[1])
}
