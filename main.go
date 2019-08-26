/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package main

import (
    "github.com/fbonhomm/api-go/source/routers"
    "github.com/fbonhomm/api-go/source/services"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "time"
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
    services.Database()

    // Initialize Jwt
    services.Jwt()

    routerEngine := gin.Default()

    routerEngine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    routers.Auth(routerEngine)
    routers.User(routerEngine)
    // for route := range routers {
    //     route(routerEngine)
    // }

    routerEngine.Run()
    services.Db.Close()
}
