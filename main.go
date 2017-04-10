// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

var (
	entryPoint     string
	eventType      string
	outputFilename string
	pluginPath     string
	shimPath       string
)

func generateOutputFilename() string {
	return fmt.Sprintf("%s-%s-%d.zip", entryPoint, eventType, time.Now().Unix())
}

func main() {
	flag.StringVar(&entryPoint, "entry-point", "F", "the name of a Go function that will be executed when the Cloud Function is triggered.")
	flag.StringVar(&eventType, "event-type", "event", "The Cloud Function event type. (bucket, http, or topic)")
	flag.StringVar(&outputFilename, "o", generateOutputFilename(), "The output file name.")
	flag.StringVar(&pluginPath, "plugin-path", "functions.so", "The path to the Go plugin that exports the function to be executed.")
	flag.StringVar(&shimPath, "shim-path", "cloud-functions-go-shim", "The path to the cloud-functions-go-shim binary.")
	flag.Parse()

	// Set the log format to basic to omit timestamps.
	log.SetFlags(0)

	var err error

	shimPath, err = exec.LookPath(shimPath)
	if err != nil {
		log.Fatalf("unable to locate the cloud-functions-go-shim binary: %s", err)
	}

	zipfile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)

	// Add the index.js file to the zip archive.

	// Generate the node.js shim based on the event type.
	nodejsShim, err := archive.Create("index.js")
	if err != nil {
		log.Fatalf("unable to generate the node.js shim: %s", err)
	}

	var tmpl string

	switch eventType {
	case "event":
		tmpl = nodejsEventTemplate
	case "http":
		tmpl = nodejsHTTPTemplate
	}

	t, err := template.New("index.js").Parse(tmpl)
	if err != nil {
		log.Fatalf("unable to generated the node.js shim: %s", err)
	}

	data := struct {
		EntryPoint string
		EventType  string
		PluginPath string
	}{
		EntryPoint: entryPoint,
		EventType:  eventType,
		PluginPath: path.Base(pluginPath),
	}

	t.Execute(nodejsShim, data)

	// Add the Go plugin to the zip archive.
	functions, err := archive.Create(path.Base(pluginPath))
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	p, err := os.Open(pluginPath)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}
	defer p.Close()

	_, err = io.Copy(functions, p)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	// Add the cloud-functions-go-shim binary to the zip archive.
	shimFile, err := os.Open(shimPath)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}
	defer shimFile.Close()

	shimInfo, err := shimFile.Stat()
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	shimHeader, err := zip.FileInfoHeader(shimInfo)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}
	shimHeader.Name = "cloud-functions-go-shim"

	shim, err := archive.CreateHeader(shimHeader)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	_, err = io.Copy(shim, shimFile)
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	err = archive.Close()
	if err != nil {
		log.Fatalf("unable to create %s: %s", outputFilename, err)
	}

	fmt.Printf("wrote %s\n", outputFilename)

	os.Exit(0)
}

const nodejsEventTemplate = `const spawnSync = require('child_process').spawnSync;

exports.{{.EntryPoint}} = function {{.EntryPoint}}(event, callback) {
  var args = [
    '--entry-point', '{{.EntryPoint}}',
    '--event-type', '{{.EventType}}',
    '--plugin-path', '{{.PluginPath}}'
  ];

  result = spawnSync('./cloud-functions-go-shim', args, {
    input: JSON.stringify(event),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     console.log(result.stderr.toString());
     callback(new Error(result.stderr.toString()));
     return;
  } else {
    callback(null, result.stdout.toString());
  }
};
`

const nodejsHTTPTemplate = `const spawnSync = require('child_process').spawnSync;

exports.{{.EntryPoint}} = function {{.EntryPoint}}(req, res) {
  var requestBody;

  switch (req.get('content-type')) {
    case 'application/json':
      requestBody = JSON.stringify(req.body);
      break;
    case 'application/octet-stream':
      requestBody = req.body;
      break;
    case 'text/plain':
      requestBody = req.body;
      break;
  }

  var fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl;
  var httpRequest = {
    'body': requestBody,
    'header': req.headers,
    'method': req.method,
    'remote_addr': req.ip,
    'url': fullUrl
  };

  var args = [
    '--entry-point', '{{.EntryPoint}}',
    '--event-type', '{{.EventType}}',
    '--plugin-path', '{{.PluginPath}}'
  ];

  result = spawnSync('./cloud-functions-go-shim', args, {
    input: JSON.stringify(httpRequest),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     console.log(result.stderr.toString());
     res.status(500);
     return;
  }

  data = JSON.parse(result.stdout);
  res.status(data.status_code);
  res.send(data.body);
};
`
