package main

import "C"

import (
	"io"
	"log"
	"net/http"
)

func F(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	io.WriteString(w, r.Method)
	io.WriteString(w, `{"message": "Go Serverless!"}`)
}
