package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Categorizando", nil)
}

func main() {
	http.HandleFunc("/api/categorize", catHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
