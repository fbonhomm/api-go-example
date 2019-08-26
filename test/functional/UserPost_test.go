/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package fonctional

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"
    "testing"
)

// TestUserPost
func TestUserPost(t *testing.T) {
    var data = url.Values{"name": {"test"}, "email": {"example@test.com"}, "password": {"12345678"}}

    req, err := http.NewRequest("POST", "http://localhost:3000/users", strings.NewReader(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Authorization", `Bearer ` + Tokens.AccessToken)
    resp, err := Client.Do(req)

    if err != nil {
        log.Panic(err)
    }

    body, err := ioutil.ReadAll(resp.Body)

    if nil != err {
        log.Panic(err)
    }

    fmt.Println(string(body[:]))

    resp.Body.Close()
}
