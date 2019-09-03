/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package services

import (
    "crypto/ecdsa"
    "io/ioutil"
    "os"

    "github.com/dgrijalva/jwt-go"
)


var PublicKeyAccess *ecdsa.PublicKey
var PrivateKeyAccess *ecdsa.PrivateKey
var PublicKeyRefresh *ecdsa.PublicKey
var PrivateKeyRefresh *ecdsa.PrivateKey
var err error

func Jwt() {
    var tmp []byte
    var root = os.Getenv("ROOT") + "/source/services"

    tmp, err = ioutil.ReadFile(root + "/jwt/access.public.pem")
    PublicKeyAccess, _ = jwt.ParseECPublicKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(root + "/jwt/access.private.pem")
    PrivateKeyAccess, _ = jwt.ParseECPrivateKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(root + "/jwt/refresh.public.pem")
    PublicKeyRefresh, _ = jwt.ParseECPublicKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(root + "/jwt/refresh.private.pem")
    PrivateKeyRefresh, _ = jwt.ParseECPrivateKeyFromPEM(tmp)
}
