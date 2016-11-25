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

  var fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl;
  var httpRequest = {
    'body': requestBody,
    'header': req.headers,
    'method': req.method,
    'remote_addr': req.ip,
    'url': fullUrl
  };

  result = spawnSync('./go-http-shim', [], {
    input: JSON.stringify(httpRequest),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     console.log(result.stderr.toString());
     res.status(500);
     res.send(result.stderr.toString());
     return;
  }

  data = JSON.parse(result.stdout);
  res.status(data.status_code);
  res.send(data.body);
};
