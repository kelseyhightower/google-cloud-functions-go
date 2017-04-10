# Google Cloud Functions Go: Topic Example

## Usage

Build the Go plugin:

```
go build -buildmode=plugin -o functions.so main.go
```

Test the function:

```
cat event.json | cloud-functions-go-shim -entry-point F -event-type topic -plugin-path functions.so
```

Create the Cloud Function zip archive:

```
cloud-functions-go -entry-point F -event-type topic -plugin-path functions.so
```

```
wrote F-topic-1491763297.zip
```

Use the Cloud Functions UI to deploy the fuction from a `ZIP upload`, set the fuction to execute to `F`, and the trigger to `Cloud Pub/Sub topic`.
