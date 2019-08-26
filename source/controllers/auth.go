/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package controllers

import (
    "github.com/fbonhomm/api-go/source/libs"
    "github.com/fbonhomm/api-go/source/models"
    "github.com/fbonhomm/api-go/source/services"
    "github.com/gin-gonic/gin"
    "net/http"
)

// AuthLogin
func AuthLogin(c *gin.Context) {
    user := models.User{
        Email: c.PostForm("email"),
    }

    if err := services.Db.First(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if err := user.Compare(c.PostForm("password")); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    accessToken, err := libs.GenerateAccessToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    refreshToken, err := libs.GenerateRefreshToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "access_token": accessToken,
        "refresh_token": refreshToken,
    })
}

// AuthRefresh
func AuthRefresh(c *gin.Context) {
    user := models.User{}

    token, _ := c.Get("Token")
    info := token.(map[string]string)

    if err := services.Db.First(&user, info["id"]).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    accessToken, err := libs.GenerateAccessToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    refreshToken, err := libs.GenerateRefreshToken(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "access_token": accessToken,
        "refresh_token": refreshToken,
    })
}
