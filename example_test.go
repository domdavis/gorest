package rest_test

import (
	"fmt"
	"net/http"

	"github.com/domdavis/rest"
)

type Version struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func ExampleNew() {
	var v Version

	h := http.Header{}
	h.Add("Accept", "application/json")

	e := rest.New("https://www.foaas.com/version", rest.MethodGet)
	r, err := e.Get(h)

	if err != nil {
		panic(err)
	}

	r.Unmarshal(&v)
	fmt.Println(v.Subtitle)
	// Output: FOAAS
}
