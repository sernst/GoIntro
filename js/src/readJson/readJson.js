var path = require('path');
var fs = require('fs');

function main() {
  var filePath = path.join(
      process.env.RESOURCES_PATH,
      'best-selling-2015-1.json'
  );
  console.log('PATH:', filePath);
  // PATH: /a/path/to/resources/best-selling-2015-1.json

  try {
    var contents = fs.readFileSync(filePath, 'utf8');
  } catch (err) {
    console.log('Failed to read file', err);
    return;
  }

  console.log('CONTENTS:', contents.substr(0, 32));
  // CONTENTS: { "make": "Ford", "model": "Fo

  try {
    var data = JSON.parse(contents);
  } catch (err) {
    console.log('Failed to load JSON data', err);
    return;
  }
  console.log(data.make + ' ' + data.model + ': ' + data.consumerRating);
  // Ford F-Series: 9.2
}
main();