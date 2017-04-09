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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"plugin"
	"strings"

	"github.com/kelseyhightower/google-cloud-functions-go/event"
)

func eventHandler(f plugin.Symbol, data []byte) (string, error) {
	var e event.Event
	var message string

	err := json.Unmarshal(data, &e)
	if err != nil {
		return "", fmt.Errorf("unable to load the event: %s", err)
	}

	message, err = f.(func(event.Event) (string, error))(e)
	if err != nil {
		return "", err
	}

	return message, nil
}

func httpHandler(f plugin.Symbol, data []byte) (string, error) {
	var httpRequest event.HTTP

	err := json.Unmarshal(data, &httpRequest)
	if err != nil {
		return "", fmt.Errorf("unable to load the event: %s", err)
	}

	r := httptest.NewRequest(httpRequest.Method, httpRequest.URL, bytes.NewBufferString(httpRequest.Body))
	for k, v := range httpRequest.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = httpRequest.RemoteAddr

	w := httptest.NewRecorder()

	f.(func(http.ResponseWriter, *http.Request))(w, r)

	resp := w.Result()

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}

	out, err := json.Marshal(&event.HTTPResponse{
		Body:       w.Body.String(),
		Header:     header,
		StatusCode: resp.StatusCode,
	})

	if err != nil {
		return "", err
	}

	return string(out), nil
}
