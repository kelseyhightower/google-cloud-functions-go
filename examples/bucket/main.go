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
	"log"

	"github.com/kelseyhightower/google-cloud-functions-go/event"
)

func F(e event.ObjectChange) (string, error) {
	log.SetFlags(0)

	log.Printf("processing object: %s", e.Data.Id)
	log.Printf("bucket: %s", e.Data.Bucket)
	log.Printf("filename: %s", e.Data.Name)
	log.Printf("resource state: %s", e.Data.ResourceState)

	return "", nil
}
