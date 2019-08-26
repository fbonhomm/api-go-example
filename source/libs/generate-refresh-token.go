/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package libs

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/fbonhomm/api-go/source/models"
    "github.com/fbonhomm/api-go/source/services"
    "time"
)

// GenerateRefreshToken
func GenerateRefreshToken(user models.User) (refreshToken string, err error) {
    token := jwt.New(jwt.SigningMethodES384)

    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    refreshToken, err = token.SignedString(services.PrivateKeyRefresh)

    return refreshToken, err
}
