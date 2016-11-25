# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions).

## Requirements

 - Linux
 - Go 1.8

## Usage

Save the following source code to a file named `function.go`:

```
package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type response struct {
     Message string `json:"message"`
}

func F(w http.ResponseWriter, r *http.Request) {
    d, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(500)
        return
    }
    if err := json.NewEncoder(w).Encode(response{Message: string(d)}); err != nil {
        w.WriteHeader(500)
        return
    }
}
```

Build the `function.so` plugin:

```
go build -buildmode=plugin function.go
```

Build the `go-http-shim` binary:

```
go build -o go-http-shim cmd/go-http-shim/main.go
```

### Testing your function

In separate terminal start the `go-http-shim` server:

```
echo "Go Serverless!" | go-http-shim 
```
```
{"message":"Go Serverless!"}
```

At this point everything is working. Now we need to package our function and the shim for use with Google Cloud Functions.

## Google Cloud Functions

```
zip -r go-serverless.zip go-http-shim function.so index.js
```

```
updating: go-http-shim (deflated 68%)
updating: function.so (deflated 71%)
updating: index.js (deflated 46%)
```

Upload `go-serverless.zip` and set the function to execute to `helloGET`
