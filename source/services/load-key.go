/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * Licence: MIT
 */

package services

import (
    "crypto/ecdsa"
    "github.com/dgrijalva/jwt-go"
    "io/ioutil"
)

const ROOT = "./source/services"

var PublicKeyAccess *ecdsa.PublicKey
var PrivateKeyAccess *ecdsa.PrivateKey
var PublicKeyRefresh *ecdsa.PublicKey
var PrivateKeyRefresh *ecdsa.PrivateKey
var err error

func init() {
    var tmp []byte

    tmp, _ = ioutil.ReadFile(ROOT + "/jwt/access.public.pem")
    PublicKeyAccess, _ = jwt.ParseECPublicKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(ROOT + "/jwt/access.private.pem")
    PrivateKeyAccess, _ = jwt.ParseECPrivateKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(ROOT + "/jwt/refresh.public.pem")
    PublicKeyRefresh, _ = jwt.ParseECPublicKeyFromPEM(tmp)

    tmp, _ = ioutil.ReadFile(ROOT + "/jwt/refresh.private.pem")
    PrivateKeyRefresh, _ = jwt.ParseECPrivateKeyFromPEM(tmp)
}
