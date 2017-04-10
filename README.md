# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions). 

## Disclaimer
This is not an official Google product (experimental or otherwise), it is just code that happens to be owned by Google.

## How it Works

Google Cloud Functions only supports node.js so shims must be used to wrap calls to Go code. The `cloud-functions-go-shim` binary bridges node.js and Go functions. Each Go function to be executed must be exported from a [Go plugin](https://golang.org/pkg/plugin/).

> The use of Go plugins limit the runtime environment to Linux.

## Install

[Download](https://github.com/kelseyhightower/google-cloud-functions-go/releases) the `cloud-functions-go` and `cloud-functions-go-shim` binaries and put them in your path.

## Usage

### Build

Create a Go plugin holding the function to be executed:

```
go build -buildmode=plugin -o functions.so examples/topic/main.go
```

### Test

Use the `cloud-functions-go-shim` to test your function:

```
cat examples/topic/event.json | \
  cloud-functions-go-shim -entry-point F -event-type topic -plugin-path functions.so 
```

> Testing only works on Linux; a current limitation of Go plugins.

### Package

The `cloud-functions-go` command is used to package your function along with the necessary shims for execution in the Cloud Functions environment. 

```
cloud-functions-go -entry-point F -event-type topic -plugin-path functions.so
```

Output

```
wrote F-topic-1491796383.zip
```

The zip archive containes the following files:

* index.js - node.js shim that calls out to the `cloud-functions-go-shim` binary
* functions.so - the Go function to be executed
* cloud-functions-go-shim - the shim between node.js and Go 

### Deploy

Upload the zip archive and set the function to `F` and the trigger to `topic`.
