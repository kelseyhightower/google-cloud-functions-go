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
