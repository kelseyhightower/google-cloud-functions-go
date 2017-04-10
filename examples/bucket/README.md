# Google Cloud Functions Go: Bucket Example

## Usage

Build the Go plugin:

```
go build -buildmode=plugin -o functions.so main.go
```

Test the function:

```
cat event.json | cloud-functions-go-shim -entry-point F -event-type bucket -plugin-path functions.so
```

Create the Cloud Function zip archive:

```
cloud-functions-go -entry-point F -event-type bucket -plugin-path functions.so
```

```
wrote F-bucket-1491804101.zip
```

Use the Cloud Functions UI to deploy the fuction from a `ZIP upload`, set the fuction to execute to `F`, and the trigger to `Cloud Storage bucket`.
