const spawnSync = require('child_process').spawnSync;

exports.helloGET = function helloGET(req, res) {
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

  result = spawnSync('./go-http-shim', [], {
    input: requestBody,
    stdio: 'pipe',
    env: {'GCF_HTTP_URL': 'http://example.com', 'GCF_HTTP_METHOD': req.method}
  });

  res.status(200);
  res.send(result.stdout);
};
