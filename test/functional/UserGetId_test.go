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
    "testing"
)

// TestUserPost
func TestGetIdPost(t *testing.T) {
    resp, err := http.Get("http://localhost:3000/users/10")

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
