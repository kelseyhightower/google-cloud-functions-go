# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions).

## Usage

```
go build -buildmode=plugin -o functions.so examples/pubsub/main.go
```

### Testing your function

```
cat examples/pubsub/event.json | \
  cloud-functions-go-shim \
    -entry-point F \
    -event-type event \
    -plugin-path functions.so 
```

> processing event: 12345678

At this point everything is working. Now we need to package our function and the shim for use with Google Cloud Functions.

## Google Cloud Functions

```
cloud-functions-go -o go-serverless.zip \
  -entry-point F \
  -event-type event \
  -plugin-path functions.so
```

> wrote go-serverless.zip

Upload `go-serverless.zip` then set the function to execute to `F` and the trigger to `Cloud Pub/Sub topic`.
