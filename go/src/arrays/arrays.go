package main

import "fmt"

func modifyArray(a [3]string) {
	a[0] = "Larry"
	a[1] = "Curly"
	a[2] = "Moe"
}

func modifySlice(a []string) {
	a[0] = "Larry"
	a[1] = "Curly"
	a[2] = "Moe"
}

func main() {
	array := [3]string{"Athos", "Porthos", "Aramis"}
	modifyArray(array)
	fmt.Println("Array:", array[0], array[1], array[2])
	// Array: Athos Porthos Aramis

	slice := []string{}
	slice = append(slice, "Athos")
	slice = append(slice, "Porthos")
	slice = append(slice, "Aramis")
	fmt.Println("Slice:", slice[0], slice[1], slice[2])
	// Slice: Athos Porthos Aramis

	modifySlice(slice)
	fmt.Println("Slice:", slice[0], slice[1], slice[2])
	// Slice: Larry Curly Moe

	slice = []string{
		"Red", "Orange", "Yellow", "Green",
		"Blue", "Indigo", "Violet",
	}
	fmt.Println("Slice:", slice[2:5])
	// Slice: [Yellow, Green, Blue]

	modifySlice(slice[2:5])
	fmt.Println("Slice:", slice)
	// Slice: [Red, Orange, Larry, Curly, Moe, Indigo, Violet]

	array = [3]string{"Athos", "Porthos", "Aramis"}
	modifySlice(array[:])
	fmt.Println("Sliced Array:", array[0], array[1], array[2])
	// Sliced Array: Larry Curly Moe
}
