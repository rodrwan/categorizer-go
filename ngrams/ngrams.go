package ngrams

import (
	"strings"
)

func MakeNgrams(s string, n int) []string {
	var ngrams []string
	splittedWords := strings.Split(s, " ")

	i := 0
  limit := len(splittedWords) - (n - 1)
	for i < limit {
			ngrams = append(ngrams, strings.Join(splittedWords[i : i+n], " "))
			i++
	}
	return ngrams
}

