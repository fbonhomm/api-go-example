/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package middlewares

import (
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/fbonhomm/api-go/source/libs"
    "github.com/fbonhomm/api-go/source/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func Auth(c *gin.Context) {
    tokenString := libs.GetToken(c)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return services.PublicKeyAccess, nil
    })

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{ "error": err })
        c.Abort()
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        c.Set("Token", claims)
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{ "error": err })
        c.Abort()
    }
}
