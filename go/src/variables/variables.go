package main

import "fmt"

func makeIndexer() (func() int) {
	value := 0

	increment := func() int {
		value += 1
		return value
	}

	return increment
}

func main() {
	var x1 int
	var x2 int = 12
	var x3 = 13
	x4 := 14
	fmt.Println("Xs:", x1, x2, x3, x4)
	// Xs: 0 12 13 14

	var y1 string
	var y2 bool
	var y3 int
	var y4 float64
	fmt.Println("Ys:", y1, y2, y3, y4)
	// Ys: "" false 0 0.0

	indexer := makeIndexer()
	fmt.Println(
		"Indexer:",
		indexer(),
		indexer(),
		indexer(),
	)
	// Indexer: 1 2 3
}
