# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions).

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

Build the `function` binary:

```
cp cmd/go-http-shim/main.go . 
go build -o function .
```

### Testing your function

```
cat request.json | function
```
```
{"body":"{\"message\":\"Go Serverless!\"}\n","header":{"Content-Type":"text/plain; charset=utf-8"},"status_code":200}
```

At this point everything is working. Now we need to package our function and the shim for use with Google Cloud Functions.

## Google Cloud Functions

Package the `function` binary and the `index.js` shim:

```
zip -r go-serverless.zip function index.js
```
```
updating: function (deflated 65%)
updating: index.js (deflated 53%)
```

Upload `go-serverless.zip` and set the function to execute to `helloGET`

Once the function is deployed [invoke it with an HTTP trigger](https://cloud.google.com/functions/docs/calling/http)

Example:

```
curl -X POST https://us-central1-hightowerlabs.cloudfunctions.net/go-serverless \
  -H "Content-Type: text/plain" \
  --data 'Go Serverless!'
```
```
{"message":"Go Serverless!"}
```
