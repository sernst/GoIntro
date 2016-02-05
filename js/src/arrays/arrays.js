function modifyArray(a) {
  a[0] = "Larry";
  a[1] = "Curly";
  a[2] = "Moe";
}

function main() {
  var array = ["Athos", "Porthos", "Aramis"];
  modifyArray(array);
  console.log("Array:", array[0], array[1], array[2]);
  // Array: Larry Curly Moe

  var slice = [];
  slice.push("Athos");
  slice.push("Porthos");
  slice.push("Aramis");
  console.log("Slice:", slice[0], slice[1], slice[2]);
  // Slice: Athos Porthos Aramis

  slice = [
    "Red", "Orange", "Yellow", "Green",
    "Blue", "Indigo", "Violet"
  ];
  console.log("Slice:", slice.slice(2, 5));
  // Slice: [Yellow, Green, Blue]

  modifyArray(slice.slice(2, 5));
  console.log("Slice:", slice);
  // Slice: [Red, Orange, Yellow, Green, Blue, Indigo, Violet]
}
main();
