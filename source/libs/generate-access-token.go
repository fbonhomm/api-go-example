/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package libs

import (
    "time"

    "github.com/dgrijalva/jwt-go"

    "github.com/fbonhomm/api-go/source/models"
    "github.com/fbonhomm/api-go/source/services"
)

// GenerateAccessToken
func GenerateAccessToken(user models.User) (accessToken string, err error) {
    token := jwt.New(jwt.SigningMethodES256)

    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["name"] = user.Name
    claims["email"] = user.Email
    claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

    accessToken, err = token.SignedString(services.PrivateKeyAccess)

    return accessToken, err
}
