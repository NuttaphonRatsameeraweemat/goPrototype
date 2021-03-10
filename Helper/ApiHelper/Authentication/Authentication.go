package authentication

import (
	"encoding/base64"
	"net/http"
	"sort"
	"strings"

	env "redis-cache-api/EnvironmentVariable"
)

// Authorized func(w http.ResponseWriter, r *http.Request)
func Authorized(w http.ResponseWriter, r *http.Request) bool {
	result := false
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	username, password, canDecrypt := decryptValue(r)
	if canDecrypt {
		authenUserList, authenPasswordList := strings.Split(env.GetBasicAuthUser(), "|"), strings.Split(env.GetBasicAuthPassword(), "|")
		index := sort.SearchStrings(authenUserList, username)
		if index < len(authenUserList) &&
			(username == authenUserList[index] && password == authenPasswordList[index]) {
			result = true
		}
	}
	return result
}

// decryptValue func(r *Request) (string, string, bool)
func decryptValue(r *http.Request) (string, string, bool) {
	canDecrypt := false
	username, password := "", ""
	authorizationHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(authorizationHeader) == 2 {
		byteValue, err := base64.StdEncoding.DecodeString(authorizationHeader[1])
		if err == nil {
			userInfo := strings.SplitN(string(byteValue), ":", 2)
			if len(userInfo) == 2 {
				username, password, canDecrypt = userInfo[0], userInfo[1], true
			}
		}
	}

	return username, password, canDecrypt
}
