function helloWorld() {
  console.log("Hello World");
}

function helloTo(name) {
  console.log("Hello " + name);
}

function helloToMany() {
  var args = Array.prototype.slice.call(arguments);
  console.log("Hello " + args.join(", "));
}

function helloInt(value) {
  console.log("Hello " + value);
}

function helloMultiply(multiplier, value) {
  console.log("Hello " + (multiplier*value).toFixed(8));
}

function main() {
  helloWorld();
  // Hello World

  helloTo("Gaia");
  // Hello Gaia

  helloToMany("Cronus", "Rhea", "Helios", "Eos");
  // Hello Cronus, Rhea, Helios, Eos

  helloInt(12);
  // Hello 12

  helloMultiply(2, 3.75);
  // Hello 7.5
}
main();
