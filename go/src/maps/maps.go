package main

import "fmt"

func modifyValues(m map[string]interface{}) {
	m["a"] = 2
	m["b"] = false
	m["c"] = "goodbye"
}

type Example struct {
	a int
	b bool
	c string
}

func modifyStruct(s Example) {
	s.a = 2
	s.b = false
	s.c = "goodbye"
}

func main() {
	m1 := map[string]int{
		"a": 12,
		"b": 22,
		"c": 8,
	}
	fmt.Println(
		"The answer to life, the universe, and everything is",
		m1["a"] + m1["b"] + m1["c"],
	)
	// The answer to life, the universe, and everything is 42

	m2 := map[string]interface{}{
		"a": 1,
		"b": true,
		"c": "hello",
	}
	modifyValues(m2)
	fmt.Println("Map:", m2["a"], m2["b"], m2["c"])
	// Map: 2 false goodbye

	s := Example{
		a:1,
		b:true,
		c:"hello",
	}
	modifyStruct(s)
	fmt.Println("Struct:", s.a, s.b, s.c)
	// Struct: 1 true hello
}
