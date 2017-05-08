package gorest_test

import (
	"testing"

	"github.com/domdavis/gorest"
	. "github.com/smartystreets/goconvey/convey"

	"net/http/httptest"

	"net/http"
)

func TestNew(t *testing.T) {
	Convey("A valid HTTP response", t, func() {
		r, err := gorest.NewResponse(httptest.NewRecorder().Result())

		Convey("Will return a proteus response", func() {
			So(r, ShouldNotBeNil)
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("An invalid response", t, func() {
		recorder := httptest.NewRecorder()
		recorder.WriteHeader(http.StatusInternalServerError)
		r, err := gorest.NewResponse(recorder.Result())

		Convey("Will return a default proteus response", func() {
			So(r, ShouldNotBeNil)
		})

		Convey("Will return an error", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Got status code: 500")
		})
	})

}

func TestResponse_Location(t *testing.T) {
	Convey("Asking a response for its location", t, func() {
		location := "http://location.com"
		res := httptest.NewRecorder().Result()
		res.Header.Add("Location", location)

		r, err := gorest.NewResponse(res)
		So(r, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("Will return the location set in the header", func() {
			So(r.Location(), ShouldEqual, location)
		})
	})
}

func TestResponse_Body(t *testing.T) {
	Convey("Asking a response for its body", t, func() {
		body := []byte("body")
		recorder := httptest.NewRecorder()
		recorder.Write(body)

		r, err := gorest.NewResponse(recorder.Result())
		So(r, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("Will return the response body", func() {
			So(r.Body(), ShouldResemble, body)
		})
	})
}

func TestResponse_Unmarshal(t *testing.T) {
	Convey("Asking a response to unmarshal its response body", t, func() {
		var o []int
		recorder := httptest.NewRecorder()
		recorder.WriteString("[1, 2]")
		r, err := gorest.NewResponse(recorder.Result())

		So(r, ShouldNotBeNil)
		So(err, ShouldBeNil)

		err = r.Unmarshal(&o)

		Convey("Will return the unmarshalled response body", func() {
			So(o, ShouldNotBeNil)
			So(o, ShouldHaveLength, 2)
			So(o, ShouldResemble, []int{1, 2})
		})

		Convey("Will not return an error", func() {
			So(err, ShouldBeNil)
		})
	})
}

func TestResponse_HTTPResponse(t *testing.T) {
	Convey("Asking a response for its raw HTTP response", t, func() {
		recorder := httptest.NewRecorder()
		recorder.Write([]byte("body"))

		r, err := gorest.NewResponse(recorder.Result())
		So(r, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("Will return the raw response", func() {
			So(r.HTTPResponse(), ShouldNotBeNil)
		})
	})
}
