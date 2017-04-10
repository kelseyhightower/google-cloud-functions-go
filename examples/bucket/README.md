# Google Cloud Functions Go: Bucket Example

## Usage

Build the plugin:

```
go build -buildmode=plugin -o functions.so examples/bucket/main.go
```

Test the plugin:

```
cat examples/bucket/event.json | \
  cloud-functions-go-shim \
    -entry-point F \
    -event-type bucket \
    -plugin-path functions.so
```

Create the Google Cloud Function source package:

```
cloud-functions-go \
  --entry-point F \
  --event-type bucket \
  --plugin-path functions.so
```

```
wrote F-bucket-1491804101.zip
```

Use the Cloud Function UI to deploy the fuction from a zip file.
