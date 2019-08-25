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
    "testing"
)

// TestLogin
func TestLogin(t *testing.T) {
    resp, err := http.PostForm("http://localhost:3000/auth", url.Values{
        "email": {"example@test.com"}, "password": {"12345678"}})

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
