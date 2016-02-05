var http = require('http');
var path = require('path');
var fs = require('fs');

function handleRequest(request, response){
  function handleError() {
    response.statusCode = 404;
    response.end();
  }

  try {
    var index = request.url.substr(1);
  } catch (_) {
    return handleError();
  }

  var filePath = path.join(
      process.env.RESOURCES_PATH,
      "best-selling-2015-" + index + '.json'
  );

  fs.exists(filePath, function (exists) {
    if (!exists) {
      return handleError();
    }

    fs.readFile(filePath, 'utf8', function (error, contents) {
      if (error) {
        return handleError();
      }

      response.writeHead(200, {'Content-Type': 'application/json'});
      response.end(contents);
    });
  });
}

function main() {
  var server = http.createServer(handleRequest);
  server.listen(8080);
}
main();
