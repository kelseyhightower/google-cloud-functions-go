# Google Cloud Functions Go - HTTP Example 

## Usage

Build the plugin:

```
go build -buildmode=plugin -o functions.so main.go
```

### Testing your function

```
cat request.json | google-cloud-functions-go --entry-point F \
  --event-type http \
  --plugin-path functions.so
```
```
{"body":"{\"message\":\"Go Serverless!\"}\n","header":{"Content-Type":"text/plain; charset=utf-8"},"status_code":200}
```

At this point everything is working. Now we need to package our function and the shim for use with Google Cloud Functions.

## Google Cloud Functions

Package the `function` binary and the `index.js` shim:

```
GOOS=linux go build -o function .
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
