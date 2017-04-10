# Google Cloud Functions Go: Bucket Example

## Usage

Build the plugin:

```
go build -buildmode=plugin -o functions.so examples/bucket/main.go
```

Create the Google Cloud Function source package:

```
google-cloud-functions-go \
  --entry-point F \
  --event-type bucket \
  --plugin-path function.so
```

```
wrote F-bucket-1491763297.zip
```

Use the Cloud Function UI to deploy the fuction from a zip file.
