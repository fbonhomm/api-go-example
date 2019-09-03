/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package fixture

import (
    "github.com/fbonhomm/api-go/source/models"
)

// User
func User() models.User {
    // gofakeit.Seed(time.Now().UnixNano())

    return models.User{
        Name: "test",
        Email: "email",
        Password: "12345678",
    }
}
