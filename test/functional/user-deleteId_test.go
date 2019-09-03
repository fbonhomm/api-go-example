/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package functional

import (
    "fmt"
    "testing"
)

// TestDeleteId
func TestDeleteId(t *testing.T) {
    var body interface{}
    var err error

    if body, err = RequestApiJson("DELETE", BaseUrl + "/users/65", nil, true); err != nil {
        t.Fatal(err)
    }

    ret := body.(map[string]interface{})
    ret = ret["item"].(map[string]interface{})

    // {"item":{"ID":65,"CreatedAt":"2019-08-27T15:45:28.971196746+02:00","UpdatedAt":"2019-08-27T15:45:28.971196746+02:00","DeletedAt":null,"name":"test1","email":"example@test1.com"}}
    fmt.Println(ret["ID"])
}
