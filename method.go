package rest

import "net/http"

type method string

// The set of valid HTTP methods an Endpoint will respond to.
const (
	MethodGet    = method(http.MethodGet)
	MethodPut    = method(http.MethodPut)
	MethodPost   = method(http.MethodPost)
	MethodDelete = method(http.MethodDelete)
)

// ToString converts a method to a string
func (m method) ToString() string {
	return string(m)
}
