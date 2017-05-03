package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Response allows easy access to the key parts of the REST response.
type Response interface {
	Location() string
	Body() []byte

	// Unmarshal unmarshals the response body into the provided type.
	Unmarshal(in interface{}) error

	// HTTPResponse returns the raw HTTP response returned from the request.
	// The body of this response will have already been read and closed. Use
	// Body or Unmarshal to read the response body rather than the underlying
	// HTTP response.
	HTTPResponse() *http.Response
}

type response struct {
	location string
	body     []byte
	response *http.Response
}

func NewResponse(res *http.Response) (Response, error) {
	defer res.Body.Close()
	r := &response{response: res}

	if url, err := res.Location(); err == nil && url != nil {
		r.location = url.String()
	}

	body, err := ioutil.ReadAll(res.Body)
	r.body = body

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return r, err
	default:
		return r, fmt.Errorf("Got status code: %d", res.StatusCode)
	}
}

func (r *response) Location() string {
	return r.location
}

func (r *response) Body() []byte {
	return r.body
}

func (r *response) Unmarshal(in interface{}) error {
	return json.Unmarshal(r.body, in)
}

func (r *response) HTTPResponse() *http.Response {
	return r.response
}
