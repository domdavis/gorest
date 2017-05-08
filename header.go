package gorest

import "net/http"

// Header provides a wrapper around an http.Header allowing for the setting
// of basic authentication details on a REST call.
type Header interface {
	Add(string, string)
	Set(string, string)
	Get(string) string
	Del(string)
	applyTo(*http.Request)
}

type header struct {
	username string
	password string
	http.Header
}

// BasicHeader returns a wrapper around an http.Header for use with REST calls.
func BasicHeader() Header {
	return &header{Header: http.Header{}}
}

// AuthHeader returns a wrapper around an http.Header and sets BasicAuth
// on REST calls.
func AuthHeader(username, password string) Header {
	return &header{username: username, password: password,
		Header: http.Header{}}
}

func (h *header) applyTo(req *http.Request) {
	req.Header = h.Header

	if h.username != "" && h.password != "" {
		req.SetBasicAuth(h.username, h.password)
	}
}
