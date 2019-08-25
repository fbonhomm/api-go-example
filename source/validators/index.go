/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package validators

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func errorHandling(c *gin.Context, msg string) {
    c.JSON(http.StatusBadRequest, gin.H{ "error": msg })
    c.Abort()
}
