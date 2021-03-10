package models

import (
	"net/http"
)

// Route struct{Name string; Method string; Pattern string; HandlerFunc HandlerFunc}
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
