package main

import (
	"encoding/json"
	"fmt"
//	"strings"
  "io/ioutil"
)

/**
I should create an structure to share the new map with transactions

*/
type Sample struct {
  desc string
	sCat int
}

func check(e error) {
  if e != nil {
	  panic(e)
  }
}

func LoadData (path string) (data map[string][]interface{}) {
	samples, err := ioutil.ReadFile(path)
	check(err)
	if err := json.Unmarshal(samples, &data); err != nil {
	  panic(err)
	}
	return
}

func main() {
  path := "./data/data.json"
	data := LoadData(path)

	for _, d := range data["data"] {
		ds, _ := d.([]interface{})
		fmt.Println(ds[0])
		// fmt.Println(strings.Split(strings.ToLower(d[0]), " "))
	}

	// fmt.Println(data["data"][0])
	// fmt.Println(data["data"][1])
	// fmt.Println(data["data"][2])
}
