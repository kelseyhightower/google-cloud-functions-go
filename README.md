# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions).

## Usage

Create a Go plugin holding the function to be executed:

```
go build -buildmode=plugin -o functions.so examples/topic/main.go
```

### Testing

Use the `cloud-functions-go-shim` to test your function:

```
cat examples/topic/event.json | \
  cloud-functions-go-shim -entry-point F -event-type topic -plugin-path functions.so 
```

> processing event: 12345678

### Packaging and deployment

The `cloud-functions-go` command is used to package your function along with the necessary shims for execution in the Cloud Functions environment. 

```
cloud-functions-go -entry-point F -event-type topic -plugin-path functions.so
```

> wrote F-event-1491796383.zip

The zip archive containes the following files:

* index.js - node.js shim that calls out to the `cloud-functions-go-shim` binary
* functions.so - the Go function to be executed
* cloud-functions-go-shim - the shim between node.js and Go 
```

Upload `F-event-1491796383.zip` then set the function to execute to `F` and the trigger to `Cloud Pub/Sub topic`.
