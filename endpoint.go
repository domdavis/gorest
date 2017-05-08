package gorest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Endpoint represents a single REST endpoint on a given URL and responding to
// certain methods.
type Endpoint interface {
	// Get makes an HTTP GET request to the endpoint URL.
	Get(http.Header) (Response, error)

	// Delete makes an HTTP DELETE request to the endpoint URL.
	Delete(http.Header) (Response, error)

	// Put makes an HTTP PUT request using the JSON representation of the
	// provided type as the request body.
	Put(http.Header, interface{}) (Response, error)

	// Post makes an HTTP POST request using the JSON representation of the
	// provided type as the request body.
	Post(http.Header, interface{}) (Response, error)
}

type endpoint struct {
	url     string
	methods map[string]struct{}
}

// New returns a representation of a REST endpoint that will respond to the
// provided HTTP methods (one or more from GET, PUT, POST and DELETE).
func New(url string, methods ...method) Endpoint {
	allowed := make(map[string]struct{}, len(methods))

	for _, m := range methods {
		allowed[m.ToString()] = struct{}{}
	}

	return &endpoint{url: url, methods: allowed}
}

func (e *endpoint) Get(header http.Header) (Response, error) {
	return e.call(http.MethodGet, header, nil)
}

func (e *endpoint) Delete(header http.Header) (Response, error) {
	return e.call(http.MethodDelete, header, nil)
}

func (e *endpoint) Put(header http.Header, body interface{}) (Response, error) {
	return e.call(http.MethodPut, header, body)
}

func (e *endpoint) Post(header http.Header,
	body interface{}) (Response, error) {
	return e.call(http.MethodPost, header, body)
}

func (e *endpoint) call(m string, header http.Header,
	body interface{}) (Response, error) {

	if _, ok := e.methods[m]; !ok {
		return nil, fmt.Errorf("Unsuported method '%s' on %s", m, e.url)
	}

	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(b)
	req, err := http.NewRequest(m, e.url, reader)

	if err != nil {
		return nil, err
	}

	req.Header = header
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return NewResponse(res)
}
