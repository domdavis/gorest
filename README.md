[![license](https://img.shields.io/github/license/domdavis/rest.svg)](https://github.com/domdavis/rest/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/domdavis/rest/all.svg)](https://github.com/domdavis/rest/releases)
[![Codeship](https://img.shields.io/codeship/4f4ea6a0-126f-0135-3216-027699a88aa9/master.svg)](https://app.codeship.com/projects/217087)
[![](https://img.shields.io/github/issues-raw/domdavis/rest.svg)](https://github.com/domdavis/rest/issues)

# RESTful package for Go

The GoREST package provides RESTful semantics to communicating with as REST
endpoint using JSON. The package user deals with Go types representing the 
request and response models avoiding _stringly_ typed models and overuse of
`map[string]interface{}`.

## Installation

```bash
go get github.com/domdavis/gorest
```

## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/domdavis/gorest"
)

type Version struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func main() {
	var v Version

	h := http.Header{}
	h.Add("Accept", "application/json")

	e := gorest.New("https://www.foaas.com/version", gorest.MethodGet)
	r, err := e.Get(h)

	if err != nil {
		panic(err)
	}

	r.Unmarshal(&v)
	fmt.Println(v.Subtitle)
	// Output: FOAAS
}
```

## License

Copyright (c) 2017 Dom Davis, distributed under the MIT license. See the 
[LICENSE](LICENSE) file for full details.
