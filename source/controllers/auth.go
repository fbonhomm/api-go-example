/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package controllers

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/fbonhomm/api-go/source/config"
    "github.com/fbonhomm/api-go/source/models"
    "github.com/fbonhomm/api-go/source/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

// Login
func AuthLogin(c *gin.Context) {
    user := models.User{
        Email: c.PostForm("email"),
    }

    if err := config.Db.First(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if err := user.Compare(c.PostForm("password")); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    token := jwt.New(jwt.SigningMethodES256)
    refresh := jwt.New(jwt.SigningMethodES384)

    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["name"] = user.Name
    claims["email"] = user.Email
    claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

    accessToken, err := token.SignedString(services.PrivateKeyAccess)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    claims = refresh.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    refreshToken, err := refresh.SignedString(services.PrivateKeyRefresh)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "access_token": accessToken,
        "refresh_token": refreshToken,
    })
}
