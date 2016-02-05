function Vertex(x, y) {
  this.x = x;
  this.y = y;
}

Vertex.prototype.length = function() {
  return Math.sqrt(
      Math.pow(this.x, 2) +
      Math.pow(this.y, 2)
  );
};

Vertex.prototype.normalize = function() {
  var length = this.length();
  this.x /= length;
  this.y /= length;
};

Vertex.prototype.zero = function() {
  this.x = 0;
  this.y = 0;
};

function main () {
  var v = new Vertex(6, 7);
  console.log("Length:", v.length().toFixed(3));
  // Length: 9.220

  v.normalize();
  console.log("Normalized:", v.x, v.y);
  // Normalized: 0.651 0.759

  v.zero();
  console.log("Zeroed:", v.x, v.y);
  // Zeroed: 0 0
}
main();