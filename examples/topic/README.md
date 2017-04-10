# Event Function

## Usage

Build the plugin:

```
go build -buildmode=plugin -o function.so main.go
```

Create the Google Cloud Function source package:

```
google-cloud-functions-go --package \
  --entry-point F \
  --event-type event \
  --plugin-path function.so
```

```
wrote F-event-1491763297.zip
```

Use the Cloud Function UI to deploy the fuction from a zip file.
