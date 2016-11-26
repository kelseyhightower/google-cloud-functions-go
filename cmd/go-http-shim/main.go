// Copyright 2016 Google Inc. All Rights Reserved.
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
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"strings"
)

type HTTPRequest struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"header"`
	Method     string            `json:"method"`
	RemoteAddr string            `json:"remote_addr"`
	URL        string            `json:"url"`
}

type HTTPResponse struct {
	Body       string            `json:"body"`
	Header     map[string]string `json:"header"`
	StatusCode int               `json:"status_code"`
}

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var httpRequest HTTPRequest
	err = json.Unmarshal(stdin, &httpRequest)
	if err != nil {
		log.Fatal(err)
	}

	r := httptest.NewRequest(httpRequest.Method, httpRequest.URL, bytes.NewBufferString(httpRequest.Body))
	for k, v := range httpRequest.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = httpRequest.RemoteAddr

	w := httptest.NewRecorder()

	F(w, r)

	resp := w.Result()

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}
	httpResponse := HTTPResponse{
		Body:       w.Body.String(),
		Header:     header,
		StatusCode: resp.StatusCode,
	}

	out, err := json.Marshal(&httpResponse)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(out)
}
