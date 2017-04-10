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

type ObjectChange struct {
	Data      StorageObject `json:"data"`
	EventId   string        `json:"eventId"`
	Timestamp string        `json:"timestamp"`
	EventType string        `json:"eventType"`
	Resource  string        `json:"resource"`
}

type StorageObject struct {
	Kind           string `json:"kind"`
	ResourceState  string `json:"resourceState"`
	Id             string `json:"id"`
	SelfLink       string `json:"selfLink"`
	Name           string `json:"name"`
	Bucket         string `json:"bucket"`
	Generation     string `json:"generation"`
	Metageneration string `json:"metageneration"`
	ContentType    string `json:"contentType"`
	TimeCreated    string `json:"timeCreated"`
	Updated        string `json:"updated"`
	TimeDeleted    string `json:"timeDeleted"`
	StorageClass   string `json:"storageClass"`
	Size           string `json:"size"`
	MD5Hash        string `json:"md5Hash"`
	MediaLink      string `json:"mediaLink"`
	CRC32c         string `json:"crc32c"`
}

type TopicPublish struct {
	Data      PubSubMessage `json:"data"`
	EventId   string        `json:"eventId"`
	Timestamp string        `json:"timestamp"`
	EventType string        `json:"eventType"`
	Resource  string        `json:"resource"`
}

type PubSubMessage struct {
	Attributes string `json:"attributes"`
	Data       string `json:"data"`
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
