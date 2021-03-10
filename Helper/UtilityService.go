package helper

import (
	"encoding/json"
	"net/http"
	helperModel "redis-cache-api/Helper/Models"
	"time"
)

// InitialResponseSuccess func(w http.ResponseWriter, result interface{})
func InitialResponseSuccess(w http.ResponseWriter, result interface{}, statusCode int, msg ...string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	message := "Complete"
	if len(msg) != 0 {
		message = msg[0]
	}
	json.NewEncoder(w).Encode(helperModel.Result{
		StatusCode: statusCode,
		Message:    message,
		Result:     result,
	})
}

// InitialResponseError func(w http.ResponseWriter, statusCode int, message string, result interface{})
func InitialResponseError(w http.ResponseWriter, statusCode int, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(helperModel.Result{
		StatusCode: statusCode,
		Message:    message,
		Result:     result,
	})
}

// ConvertStringToTime func(value string) Time
func ConvertStringToTime(value string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	result, _ := time.Parse(layout, value)
	return result
}
