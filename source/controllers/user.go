/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package controllers

import (
    "github.com/fbonhomm/api-go/source/config"
    "github.com/fbonhomm/api-go/source/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

// UserCreate
func UserCreate(c *gin.Context) {
    user := models.User{
        Name: c.PostForm("name"),
        Email: c.PostForm("email"),
        Password: c.PostForm("password"),
    }

    if err := config.Db.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{ "error": err.Error() })
    } else {
        c.JSON(http.StatusCreated, gin.H{ "item": user })
    }
}

// UserGetId
func UserGetId(c *gin.Context) {
    user := models.User{}

    if err := config.Db.First(&user, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
    } else {
        c.JSON(http.StatusOK, gin.H{ "item": user })
    }
}

// UserDeleteId
func UserDeleteId(c *gin.Context) {
    user := models.User{}

    if err := config.Db.First(&user, c.Param("id")).Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error() })
    } else {
        c.JSON(http.StatusOK, gin.H{ "item": user })
    }
}
