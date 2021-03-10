package middleware

import (
	"net/http"

	logManager "redis-cache-api/Helper/ApiHelper/LogManager"

	"github.com/google/uuid"

	helper "redis-cache-api/Helper"
	authen "redis-cache-api/Helper/ApiHelper/Authentication"
	constants "redis-cache-api/Helper/Constants"
)

// MiddlewareHandle func(h Handler) Handler
func MiddlewareHandle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID := uuid.New()
		defer func() {
			if err := recover(); err != nil {
				logError(sessionID.String(), err.(error).Error())
				helper.InitialResponseError(w, 500, constants.HTTPInternalServerErrorMessage, nil)
			}
		}()
		if authen.Authorized(w, r) {
			h.ServeHTTP(w, r)
		} else {
			helper.InitialResponseError(w, 401, constants.HTTPUnauthorized, nil)
		}
	})
}

// BeginInvoke func(url string, sessionID string)
func beginInvoke(url string, sessionID string) {
	logManager.LogInfo("session-id :" + sessionID + "| About to start " + url + " request")
}

// EndInvoke func(sessionID string)
func endInvoke(sessionID string) {
	logManager.LogInfo("session-id :" + sessionID + "| Request completed")
}

// logError func(sessionID string)
func logError(sessionID string, errorMsg string) {
	logManager.LogError("session-id :" + sessionID + "| The error message : " + errorMsg)
}
