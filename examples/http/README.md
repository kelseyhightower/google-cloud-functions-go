# Google Cloud Functions Go - HTTP Example 

## Usage

Build the Go plugin:

```
go build -buildmode=plugin -o functions.so examples/http/main.go
```

Test the Go plugin with some test data:

```
cat examples/http/request.json | \
  cloud-functions-go-shim \
    -entry-point F \
    -event-type http \
    -plugin-path functions.so
```

> {"body":"{\"message\":\"Go Serverless!\"}\n","header":{"Content-Type":"text/plain; charset=utf-8"},"status_code":200}

### Package

```
cloud-functions-go -o go-serverless.zip \
  -entry-point F \
  -event-type event \
  -plugin-path functions.so
```

Upload the `go-serverless.zip` package the function to execute to `F` and the trigger to `HTTP trigger`.

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
