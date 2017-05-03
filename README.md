# RESTful package for Go

The REST package provides RESTful semantics to communicating with as REST
endpoint using JSON. The package user deals with Go types representing the 
request and response models avoiding _stringly_ typed models and overuse of
`map[string]interface{}`.

## Installation

```bash
go get github.com/domdavis/rest
```

## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/domdavis/rest"
)

type Version struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func main() {
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
```

## License

Copyright (c) 2017 Dom Davis, distributed under the MIT license. See the 
[LICENSE](LICENSE) file for full details.
