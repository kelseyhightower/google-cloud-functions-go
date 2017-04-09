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

package event

type Event struct {
	Data      interface{} `json:"data"`
	EventID   string      `json:"eventId"`
	Timestamp string      `json:"timestamp"`
	EventType string      `json:"eventType"`
	Resource  string      `json:"resource"`
}

type PubsubMessage struct {
	Attributes  map[string]string `json:"attributes"`
	Data        string            `json:"data"`
	MessageId   string            `json:"messageId"`
	PublishTime string            `json:"publishTime"`
}

type HTTP struct {
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
