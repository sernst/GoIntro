function makeIndexer() {
  var value = 0;

  function increment() {
    value += 1;
    return value;
  }

  return increment;
}

function main() {
  var x1;
  var x2 = 12;
  console.log("Xs:", x1, x2);
  // Xs: undefined 12

  var y1, y2, y3, y4;
  console.log("Ys:", y1, y2, y3, y4);
  // Ys: undefined undefined undefined undefined

  var indexer = makeIndexer();
  console.log(
      "Indexer:",
      indexer(),
      indexer(),
      indexer()
  );
  // Indexer: 1 2 3
}
main();
