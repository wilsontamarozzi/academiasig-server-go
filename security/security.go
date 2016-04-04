package security

import (
	"crypto/md5"
    "encoding/hex"
	
	"academiasig-api/services"
	"github.com/goji/httpauth"
)

func GetAuthOpts() httpauth.AuthOptions {

	return httpauth.AuthOptions {
  		Realm: "DevCo",
    	AuthFunc: AuthenticationUser,
	}
}

func AuthenticationUser(user string, password string) bool {
    if len(services.AuthenticationUser(user, password)) == 1 {
        return true
    } 

    return false
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}