package main

import "fmt"
import "net/http"
import "os"
import "path"
import "io/ioutil"

func handler(w http.ResponseWriter, r *http.Request) {
	index := r.URL.Path[1:]
	filePath := path.Join(
		os.Getenv("RESOURCES_PATH"),
		"best-selling-2015-" + string(index) + ".json",
	)

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(contents))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}