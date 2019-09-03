/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package functional

import (
    "log"
    "net/http"
    "net/url"
    "os"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "github.com/fbonhomm/api-go/source/services"
    "github.com/fbonhomm/api-go/test/fixture"
)

type Token struct {
    AccessToken     string `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
}

const RegexToken = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`

var BaseUrl string
var Client = &http.Client{}
var Tokens Token


func getToken() {
    var body map[string]interface{}
    var err error
    var data = url.Values{"email": {"example@test.com"}, "password": {"12345678"}}

    if body, err = RequestApiJson("POST", BaseUrl + "/auth", data, false); err != nil {
        log.Fatal(err)
    }

    Tokens.RefreshToken = body["refresh_token"].(string)
    Tokens.AccessToken = body["access_token"].(string)
}

func setupDB() {
    user := fixture.User()

    if err := services.Db.Create(&user).Error; err != nil {
        log.Fatal(err)
    }
}

func removeDB() {
    services.Db.Exec("DROP TABLE users;")
    services.Db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME") + ";")
}

// init
func init() {
    // Initialize environment variable contained in .env
    if err := godotenv.Load("../../.env.testing"); err != nil {
        log.Fatal("No .env.testing file found")
    }

    BaseUrl = os.Getenv("HOST") + ":" + os.Getenv("PORT")
}

// TestUserPost
func TestMain(m *testing.M) {
    var ret int

    // Set gin mode to test
    gin.SetMode(gin.TestMode)

    // Initialize Database
    services.Database()

    // Initialize Jwt
    services.Jwt()

    // Setup DB
    setupDB()

    // Get token for test request
    getToken()

    ret = m.Run()

    // Remove DB
    removeDB()

    os.Exit(ret)
}