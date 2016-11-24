const spawn = require('child_process').spawn;
var request = require('request');

exports.helloGET = function helloGET (req, res) {
  const child = spawn("./go-http-shim", [], {
    detached: true,
    stdio: 'ignore'
  });

  request({
    url: "http://unix:/tmp/go-http-shim.sock:/",
    method: req.method,
    headers: {
      "User-Agent": req.get("User-Agent"),
      "Host": req.hostname
    }
  },
  function (error, response, body) {
    if (error) {
      console.log(error);
    }
    res.status(response.statusCode);
    res.send(body);
  });
};
