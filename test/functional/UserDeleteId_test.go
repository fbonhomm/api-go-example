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

// TestDeleteId
func TestDeleteId(t *testing.T) {
    req, err := http.NewRequest("DELETE", "http://localhost:3000/users/14", nil)
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
