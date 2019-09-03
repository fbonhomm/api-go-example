/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package functional

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"

    "github.com/fbonhomm/api-go/source/routers"
)

// RequestApiJson
func RequestApiJson(method, url string, data url.Values, auth bool) (map[string]interface{}, error) {
    var result map[string]interface{}
    var req *http.Request
    var res *httptest.ResponseRecorder
    var err error

    routerEngine := routers.RouterInitialize()

    if req, err = http.NewRequest(method, url, strings.NewReader(data.Encode())); err != nil {
        return result, err
    }

    if auth == true {
        req.Header.Add("Authorization", `Bearer ` + Tokens.AccessToken)
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    res = httptest.NewRecorder()
    routerEngine.ServeHTTP(res, req)

    if err = json.Unmarshal(res.Body.Bytes(), &result); err != nil {
        return result, err
    }

    return result, err
}
