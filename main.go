/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package main

import (
    "github.com/fbonhomm/api-go/source/config"
    "github.com/fbonhomm/api-go/source/routers"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

// init
func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }
}

// main
func main() {
    // Initialize Database
    db := config.Database()

    routerEngine := gin.Default()

    routers.Auth(routerEngine)
    routers.User(routerEngine)
    // for route := range routers {
    //     route(routerEngine)
    // }

    routerEngine.Run()
    db.Close()
}
