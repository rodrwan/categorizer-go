package extractor

import (
	"encoding/json"
  "io/ioutil"
)

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

