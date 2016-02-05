package main

import "fmt"
import "strings"

func helloWorld() {
	fmt.Println("Hello World")
}

func helloTo(name string) {
	fmt.Println("Hello " + name)
}

func helloToMany(names ...string) {
	fmt.Println("Hello " + strings.Join(names, ", "))
}

func helloInt(value int) {
	fmt.Printf("Hello %v\n", value)
}

func helloMultiply(multiplier int, value float64) {
	fmt.Printf("Hello %.8f\n", float64(multiplier)*value)
}

func main() {
	helloWorld()
	// Hello World

	helloTo("Gaia")
	// Hello Gaia

	helloToMany("Cronus", "Rhea", "Helios", "Eos")
	// Hello Cronus, Rhea, Helios, Eos

	helloInt(12)
	// Hello 12

	helloMultiply(2, 3.75)
	// Hello 7.5
}
