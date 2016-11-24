const spawn = require('child_process').spawn;
var request = require('request');
var bodyParser = require('body-parser');

exports.helloGET = function helloGET(req, res) {
  const child = spawn('./go-http-shim', [], {
    detached: true,
    stdio: 'ignore'
  });

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

  request({
    url: 'http://unix:/tmp/go-http-shim.sock:/',
    method: req.method,
    headers: {
      'User-Agent': req.get('User-Agent'),
      'Host': req.hostname
    },
    body: requestBody 
  },
  function (error, response, body) {
    if (error) {
      console.log(error);
    }
    res.status(response.statusCode);
    res.send(body);
  });
};
