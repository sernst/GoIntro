package main

import "os"
import "path"
import "fmt"
import "io/ioutil"
import "encoding/json"

type BestSellingCar struct {
	Make           string  `json:"make"`
	Model          string  `json:"model"`
	ConsumerRating float64 `json:"consumerRating"`
}

func main() {
	filePath := path.Join(
		os.Getenv("RESOURCES_PATH"),
		"best-selling-2015-1.json",
	)
	fmt.Println("PATH:", filePath)
	// PATH: /a/path/to/resources/best-selling-2015-1.json

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	fmt.Println("CONTENTS:", string(contents)[:32])
	// CONTENTS: { "make": "Ford", "model": "Fo

	var raw interface{}
	err = json.Unmarshal(contents, &raw)
	if err != nil {
		fmt.Println("Failed to load JSON data", err)
		return
	}

	data, ok := raw.(map[string]interface{})
	if !ok {
		fmt.Println("Failed to parse unexpected JSON structure")
		return
	}

	fmt.Printf(
		"%s %s: %v\n",
		data["make"],
		data["model"],
		data["consumerRating"],
	)
	// Ford F-Series: 9.2

	var bestSeller = BestSellingCar{}

	err = json.Unmarshal(contents, &bestSeller)
	if err != nil {
		fmt.Println("Failed to load JSON data into structure", err)
		return
	}

	fmt.Printf(
		"%s %s: %v\n",
		bestSeller.Make,
		bestSeller.Model,
		bestSeller.ConsumerRating,
	)
	// Ford F-Series: 9.2
}
