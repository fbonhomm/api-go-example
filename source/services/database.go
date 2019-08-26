/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package services

import (
    "fmt"
    "github.com/fbonhomm/api-go/source/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "os"
)

var Db *gorm.DB
var Err error

func Database() {
    Db, Err = gorm.Open(
        "postgres",
        fmt.Sprintf(
            "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
            os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
            os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
    )

    if Err != nil {
        log.Panic("Error connect database: ", Err)
    }

    Db.AutoMigrate(&models.User{})
}
