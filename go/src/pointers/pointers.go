package main

import "fmt"
import "math"

type Vertex struct {
	X float64
	Y float64
}

func modifyX(v Vertex) {
	v.X = -100
}

func modifyY(v *Vertex) {
	v.Y = -100
}

func makeVertex(angleDegrees float64) *Vertex {
	angleRads := angleDegrees*math.Pi/180.0
	return &Vertex{
		X: math.Cos(angleRads),
		Y: math.Sin(angleRads),
	}
}

func modifyArray(a [3]string) {
	a[0] = "Larry"
	a[1] = "Curly"
	a[2] = "Moe"
}

func modifyArrayPointer(a *[3]string) {
	a[0] = "Larry"
	a[1] = "Curly"
	a[2] = "Moe"
}

func main() {
	v := Vertex{X:6, Y:7}
	modifyX(v)
	fmt.Println("Modified X:", v.X, v.Y)
	// Modified X: 6 7

	modifyY(&v)
	fmt.Println("Modified Y:", v.X, v.Y)
	// Modified Y: 6 -100

	v = *makeVertex(90)
	fmt.Println("MakeVertex:", v.X, v.Y)
	// MakeVertex: 0 1

	v2 := makeVertex(45)
	fmt.Println("MakeVertex:", v2.X, v2.Y)
	// MakeVertex: 0.707 0.707

	v = Vertex{1, 2}
	v3 := v
	v.X += 9
	fmt.Println("V3:", v3.X, v3.Y)
	// V3: 1 2

	array := [3]string{"Athos", "Porthos", "Aramis"}
	modifyArray(array)
	fmt.Println("Array:", array[0], array[1], array[2])
	// Array: Athos Porthos Aramis

	modifyArrayPointer(&array)
	fmt.Println("Array Pointer:", array[0], array[1], array[2])
	// Array Pointer: Larry Curly Moe
}
