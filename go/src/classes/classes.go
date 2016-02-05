package main

import "fmt"
import "math"

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) length() float64 {
	return math.Sqrt(
		math.Pow(v.X, 2.0) +
		math.Pow(v.Y, 2.0),
	)
}

func (v Vertex) normalize() {
	length := v.length()
	v.X /= length
	v.Y /= length
}

func (v *Vertex) zero() {
	v.X = 0.0
	v.Y = 0.0
}

func main() {
	v := Vertex{X: 6, Y: 7}
	fmt.Printf("Length: %.3f\n", v.length())
	// Length: 9.220

	v.normalize()
	fmt.Println("Normalized:", v.X, v.Y)
	// Normalized: 6 7

	v.zero()
	fmt.Println("Zeroed:", v.X, v.Y)
	// Zeroed: 0 0
}
