// go-http-shim loads external Go plugins that meet the following
// specification:
//
//   - plugins must be named "function.so"
//   - plugins must export the following function:
//       F(w http.ResponseWriter, r *http.Request)
//
// External plugin can be built with the following command:
//
//    go build -buildmode=plugin function.go
//

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"plugin"
)

func main() {
	p, err := plugin.Open("function.so")
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest(os.Getenv("GCF_HTTP_METHOD"), os.Getenv("GCF_HTTP_URL"), bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	f.(func(http.ResponseWriter, *http.Request))(w, r)

	io.WriteString(os.Stdout, w.Body.String())
}
