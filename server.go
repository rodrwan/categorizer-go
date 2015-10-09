package main

import (
	"fmt"
	"net/http"
	"./extractor"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func catHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Categorizando")
}

func main() {
	path := "./data/data.json"
	data := extractor.LoadData(path)

	for _, d := range data["data"] {
		ds, _ := d.([]interface{})
	  fmt.Println(ds[0])
	}
	http.HandleFunc("/api/categorize", catHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
