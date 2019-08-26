/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package fonctional

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
    "testing"
)

type Token struct {
    AccessToken     string `json:"access_token"`
    RefreshToken    string  `json:"refresh_token"`
}

var Client = &http.Client{}
var Tokens Token

// TestUserPost
func TestMain(m *testing.M) {
    var data = url.Values{"email": {"example@test.com"}, "password": {"12345678"}}

    req, _ := http.NewRequest("POST", "http://localhost:3000/auth", strings.NewReader(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, _ := Client.Do(req)

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    json.Unmarshal(body, &Tokens)

    os.Exit(m.Run())
}
