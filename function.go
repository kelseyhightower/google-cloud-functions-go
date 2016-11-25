package main

import (
	"C"
	"io/ioutil"
	"net/http"
)

func F(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(d)
}
