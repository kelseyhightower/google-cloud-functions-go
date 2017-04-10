# Google Cloud Functions Go 

This project contains a collection of tutorials and hacks for using Go with [Google Cloud Functions](https://cloud.google.com/functions).

## Usage

```
go build -buildmode=plugin -o functions.so examples/topic/main.go
```

### Testing your function

```
cat examples/topic/event.json | \
  cloud-functions-go-shim \
    -entry-point F \
    -event-type topic \
    -plugin-path functions.so 
```

> processing event: 12345678

At this point everything is working. Now we need to package our function and the shim for use with Google Cloud Functions.

## Google Cloud Functions

```
cloud-functions-go \
  -entry-point F \
  -event-type topic \
  -plugin-path functions.so
```

> wrote F-event-1491796383.zip

Upload `F-event-1491796383.zip` then set the function to execute to `F` and the trigger to `Cloud Pub/Sub topic`.
