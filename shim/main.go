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
	"flag"
	"io/ioutil"
	"log"
	"os"
	"plugin"
)

var (
	entryPoint   string
	eventType    string
	pluginPath   string
)

func main() {
	flag.StringVar(&entryPoint, "entry-point", "F", "the name of a Go function that will be executed when the Cloud Function is triggered.")
	flag.StringVar(&eventType, "event-type", "", "The Cloud Function event type: http or event")
	flag.StringVar(&pluginPath, "plugin-path", "functions.so", "The path to the Go plugin that exports the function to be executed.")
	flag.Parse()

	if eventType == "" {
		log.Fatal("Event type required; set using the --event-type flag.")
	}

	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalf("unable to load the Go plugin: %s", err)
	}

	f, err := p.Lookup(entryPoint)
	if err != nil {
		log.Fatalf("unable to load the entrypoint function: %s", err)
	}

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to load the event: %s", err)
	}

	var message string

	switch eventType {
	case "event":
		message, err = eventHandler(f, stdin)
	case "http":
		message, err = httpHandler(f, stdin)
	default:
		log.Fatalf("invalid event type: %s", eventType)
	}

	if err != nil {
		log.Fatal(err)
	}

	if message != "" {
		os.Stdout.WriteString(message)
	}

	os.Exit(0)
}
