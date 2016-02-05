function modifyValues(o) {
  o.a = 2;
  o.b = false;
  o.c = 'goodbye';
}

function Example(a, b, c) {
  this.a = a;
  this.b = b;
  this.c = c;
}

function main() {
  var m1 = {
    a: 12,
    b: 22,
    c: 8
  };
  console.log(
      "The answer to life, the universe, and everything is",
      m1["a"] + m1["b"] + m1["c"]
  );
  // The answer to life, the universe, and everything is 42

  var m2 = {
    a: 1,
    b: true,
    c: "hello"
  };
  modifyValues(m2);
  console.log("Object:", m2.a, m2.b, m2.c);
  // Object: 2 false goodbye

  var s = new Example(1, true, "hello");
  modifyValues(s);
  console.log("Example:", s.a, s.b, s.c);
  // Example: 2 false goodbye
}
main();
