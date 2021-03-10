package models

// Result struct{StatusCode int "json:\"statusCode\""; Message string "json:\"message\""; Result interface{} "json:\"result\""}
type Result struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
