package main

import (
	"fmt"
	"strings"
	"net/http"
	"./extractor"
  "./ngrams"
	"./model"
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
  var newModel model.NgramModel
  newModel.Init()

	for _, d := range data["data"] {
    ds, _ := d.([]interface{})
		sGram := ngrams.MakeNgrams(strings.ToLower(ds[0].(string)), 3)
		for _, gram := range sGram {
			fmt.Println(gram)
		  newModel.FitModel(gram)
		}
		fmt.Println("")
	}
	http.HandleFunc("/api/categorize", catHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
