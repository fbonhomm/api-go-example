/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package unit

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMain
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
