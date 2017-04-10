# Google Cloud Functions Go: HTTP Example 

## Usage

Build the Go plugin:

```
go build -buildmode=plugin -o functions.so main.go
```

Test the function:

```
cat request.json | cloud-functions-go-shim -entry-point F -event-type http -plugin-path functions.so
```

> {"body":"{\"message\":\"Go Serverless!\"}\n","header":{"Content-Type":"text/plain; charset=utf-8"},"status_code":200}

Create the Cloud Function zip archive:

```
cloud-functions-go -entry-point F -event-type http -plugin-path functions.so
```

Use the Cloud Functions UI to deploy the fuction from a `ZIP upload`, set the fuction to execute to `F`, and the trigger to `HTTP trigger`.

Once the function is deployed [invoke it with an HTTP trigger](https://cloud.google.com/functions/docs/calling/http)

Example:

```
curl -X POST https://us-central1-cloud-functions-go.cloudfunctions.net/http-test   \
  -H "Content-Type: text/plain" \
  --data 'Go Serverless!'
```
```
{"message":"Go Serverless!"}
```
