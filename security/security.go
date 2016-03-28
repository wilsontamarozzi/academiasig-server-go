package security

import (
	"crypto/md5"
    "encoding/hex"
	
	"AcademiaSIG-API/services"
	"github.com/goji/httpauth"
)

func GetAuthOpts() httpauth.AuthOptions {

	return httpauth.AuthOptions {
  		Realm: "DevCo",
    	AuthFunc: AuthenticationUser,
	}
}

func AuthenticationUser(user string, password string) bool {
    return services.AuthenticationUser(user, GetMD5Hash(password))
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}