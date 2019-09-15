/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package unit

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

// CreateContext
func CreateContext(data url.Values) (res *httptest.ResponseRecorder, c *gin.Context) {
	var req http.Request
	res = httptest.NewRecorder()

	c, _ = gin.CreateTestContext(res)

	req.PostForm = data
	c.Request = &req

	return res, c
}
