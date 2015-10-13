package model

import (
  "strings"
)

type NgramModel struct {
	bagOfNgrams map[string]int
}

func (nm *NgramModel) Init() {
  nm.bagOfNgrams = make(map[string]int)
}

func (nm *NgramModel) FitModel(ngram string) {
	if _, ok := nm.bagOfNgrams[strings.ToLower(ngram)]; ok {
		nm.bagOfNgrams[ngram] += 1
  } else {
	  nm.bagOfNgrams[ngram] = 1
	}
}

func (nm *NgramModel) GetNgram(ngram string) int{
  return nm.bagOfNgrams[strings.ToLower(ngram)]
}
