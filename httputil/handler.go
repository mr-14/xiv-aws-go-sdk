package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Vars URL path variables
type Vars map[string]string

// Status Http Status
type Status uint

// Body response body
type Body interface{}

// Wrapper wraps handler function to provide ease of use for RESTful APIs
func Wrapper(handler func(vars Vars, req *http.Request) (Status, Body)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status, respObj := handler(mux.Vars(r), r)
		SetResponse(w, status, respObj)
	}
}

// SetResponse sets response
func SetResponse(w http.ResponseWriter, status Status, respBody Body) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

	if respBody != nil {
		json.NewEncoder(w).Encode(respBody)
	}
}
