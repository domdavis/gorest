package gorest_test

import (
	"fmt"

	"github.com/domdavis/gorest"
)

type Version struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func ExampleNew() {
	var v Version

	h := gorest.BasicHeader()
	h.Set("Accept", "application/json")

	e := gorest.New("https://www.foaas.com/version", gorest.MethodGet)
	r, err := e.Get(h)

	if err != nil {
		panic(err)
	}

	r.Unmarshal(&v)
	fmt.Println(v.Subtitle)
	// Output: FOAAS
}
