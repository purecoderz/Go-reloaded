package packs

import (
	"slices"
	"strings"
)

func Quotes(words []string, index, start, end int) ([]string, int, int, int) {
	if start >= 0 {
		end = index
		words[start] += strings.Join(words[start+1:end], " ") + "'"
		words = slices.Delete(words, start+1, end+1)
		index = start
		start, end = -1, -1
	} else {
		start = index
	}

	return words, index, start, end
}

func Apostrophe(words []string, index int) ([]string, int) {
	if index > 0 {
		words[index-1] += words[index]
		words = slices.Delete(words, index, index+1)
		index--
	}
	return words,index
}
