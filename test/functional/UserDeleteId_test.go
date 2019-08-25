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
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", "http://localhost:3000/users/14", nil)
    resp, err := client.Do(req)

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
