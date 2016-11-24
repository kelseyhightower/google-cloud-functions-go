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
	"log"
	"net"
	"net/http"
	"plugin"
)

const socket = "/tmp/go-http-shim.sock"

func main() {
	p, err := plugin.Open("function.so")
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		log.Fatal(err)
	}

	m := http.NewServeMux()
	m.HandleFunc("/", f.(func(http.ResponseWriter, *http.Request)))

	l, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.Serve(l, m); err != nil {
		log.Fatal(err)
	}
}
