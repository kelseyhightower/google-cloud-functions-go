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
	"path"
	"time"
)

var (
	entryPoint string
	eventType  string
	pluginPath string
)

func main() {
	flag.StringVar(&entryPoint, "entry-point", "F", "the name of a Go function that will be executed when the Google Cloud Function is triggered.")
	flag.StringVar(&eventType, "event-type", "", "The Google Cloud Function event type. [http|event]")
	flag.StringVar(&pluginPath, "plugin-path", "function.so", "The path to the Go plugin that exports the function to be executed.")
	flag.Parse()

	if eventType == "" {
		log.Fatal("Event type required; set using the --event-type flag.")
	}

	packageName := fmt.Sprintf("%s-%s-%d.zip", entryPoint, eventType, time.Now().Unix())

	zipfile, err := os.Create(packageName)
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)

	jsOutput, err := archive.Create("index.js")
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	var js string

	switch eventType {
	case "event":
		js = eventJavascriptShim
	case "http":
		js = httpJavascriptShim
	}

	t, err := template.New("js").Parse(js)
	if err != nil {
		log.Fatalf("unable to generated javascript shim: %s", err)
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

	t.Execute(jsOutput, data)

	functions, err := archive.Create(path.Base(pluginPath))
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	targetPlugin, err := os.Open(pluginPath)
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}
	defer targetPlugin.Close()

	_, err = io.Copy(functions, targetPlugin)
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	info, err := AssetInfo("google-cloud-functions-go")
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	shim, err := archive.CreateHeader(header)
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	shimData, err := Asset("google-cloud-functions-go")
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}
	shim.Write(shimData)

	err = archive.Close()
	if err != nil {
		log.Fatalf("unable to create %s: %s", packageName, err)
	}

	fmt.Printf("wrote %s", packageName)

	os.Exit(0)
}


const eventJavascriptShim = `const spawnSync = require('child_process').spawnSync;

exports.{{.EntryPoint}} = function {{.EntryPoint}}(event, callback) {
  var args = [
    '--entry-point', '{{.EntryPoint}}',
    '--event-type', '{{.EventType}}',
    '--plugin-path', '{{.PluginPath}}'
  ];

  result = spawnSync('./google-cloud-functions-go', args, {
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

const httpJavascriptShim = `const spawnSync = require('child_process').spawnSync;

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

  result = spawnSync('./google-cloud-functions-go', args, {
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
