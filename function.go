package main

import (
    "io"
    "net/http"
)

func F(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, `{"message": "Go Serverless!"}`)
}
