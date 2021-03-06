package gorest_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"encoding/json"

	"fmt"

	"github.com/domdavis/gorest"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEndpoint_Get(t *testing.T) {
	Convey("Calling an endpoint that supports GET", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "body")
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodGet)
		r, err := e.Get(h)

		Convey("Will return a populated response", func() {
			var body string

			So(r, ShouldNotBeNil)
			So(r.Body(), ShouldNotBeEmpty)
			r.Unmarshal(&body)
			So(body, ShouldEqual, "body")
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("Calling an endpoint that does not support GET ", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "body")
		defer s.Close()

		e := gorest.New(s.URL)
		r, err := e.Get(h)

		Convey("Will return a nil response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldStartWith, "Unsuported method 'GET' on")
		})
	})
}

func TestEndpoint_Put(t *testing.T) {
	Convey("Calling an endpoint that supports PUT", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "response")
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodPut)
		r, err := e.Put(h, "request")

		Convey("Will return a populated response", func() {
			var body string

			So(r, ShouldNotBeNil)
			So(r.Body(), ShouldNotBeEmpty)
			r.Unmarshal(&body)
			So(body, ShouldEqual, "response")
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("Calling an endpoint that does not support PUT", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "response")
		defer s.Close()

		e := gorest.New(s.URL)
		r, err := e.Put(h, "request")

		Convey("Will return a nil response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldStartWith, "Unsuported method 'PUT' on")
		})
	})
}

func TestEndpoint_Post(t *testing.T) {
	Convey("Calling an endpoint that supports POST", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "response")
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodPost)
		r, err := e.Post(h, "request")

		Convey("Will return a populated response", func() {
			var body string

			So(r, ShouldNotBeNil)
			So(r.Body(), ShouldNotBeEmpty)
			r.Unmarshal(&body)
			So(body, ShouldEqual, "response")
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("Calling an endpoint that does not support POST", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "response")
		defer s.Close()

		e := gorest.New(s.URL)
		r, err := e.Post(h, "request")

		Convey("Will return a nil response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldStartWith, "Unsuported method 'POST' on")
		})
	})
}

func TestEndpoint_Delete(t *testing.T) {
	Convey("Calling an endpoint that supports DELETE", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "body")
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodDelete)
		r, err := e.Delete(h)

		Convey("Will return a populated response", func() {
			var body string

			So(r, ShouldNotBeNil)
			So(r.Body(), ShouldNotBeEmpty)
			r.Unmarshal(&body)
			So(body, ShouldEqual, "body")
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("Calling an endpoint that does not support DELETE", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "body")
		defer s.Close()

		e := gorest.New(s.URL)
		r, err := e.Delete(h)

		Convey("Will return a nil response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldStartWith, "Unsuported method 'DELETE' on")
		})
	})
}

func TestEndpoint(t *testing.T) {
	Convey("Calling an endpoint with an invalid hostname", t, func() {

		e := gorest.New("http://\\", gorest.MethodGet)
		r, err := e.Get(gorest.BasicHeader())

		Convey("Will not return a response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual,
				`parse http://\: invalid character "\\" in host name`)
		})
	})

	Convey("Calling an endpoint with an invalid body", t, func() {
		h := gorest.BasicHeader()
		s := testEndpoint(200, "response")
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodPut)
		r, err := e.Put(h, make(chan string))

		Convey("Will not return a response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "json: unsupported type: chan string")
		})
	})

	Convey("Calling an endpoint with an invalid URL", t, func() {
		e := gorest.New("invalid", gorest.MethodGet)
		r, err := e.Get(gorest.BasicHeader())

		Convey("Will not return a response", func() {
			So(r, ShouldBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual,
				`Get invalid: unsupported protocol scheme ""`)
		})
	})
}

func TestSecureEndpoint(t *testing.T) {
	Convey("Calling a secure endpoint", t, func() {
		username, password := "username", "password"
		h := gorest.AuthHeader(username, password)
		s := testSecureEndpoint()
		defer s.Close()

		e := gorest.New(s.URL, gorest.MethodGet)
		r, err := e.Get(h)

		Convey("Will return a populated response", func() {
			So(r, ShouldNotBeNil)
			So(r.Body(), ShouldNotBeEmpty)
			So(string(r.Body()), ShouldStartWith, username)
			So(string(r.Body()), ShouldEndWith, password)
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}

func testEndpoint(code int, body interface{}) *httptest.Server {
	b, _ := json.MarshalIndent(body, "", "    ")
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Location", "http://localhost")
			w.WriteHeader(code)
			w.Write(b)
		}))

	return server
}

func testSecureEndpoint() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			username, password, _ := r.BasicAuth()
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Location", "http://localhost")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s:%s", username, password)
		}))

	return server
}
