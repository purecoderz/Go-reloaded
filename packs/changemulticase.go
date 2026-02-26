package packs

import (
	"slices"
	"strconv"
	"strings"
)

func ChangeMultipleCase(index int, words []string, action func(string) string) ([]string, int) {
	// println("Am here")
	num, err := strconv.Atoi(strings.TrimSuffix(words[index+1], ")"))
	if err != nil {
		return words, index
	}

	// if num > index

	for num > 0 {
		words[index-num] = action(words[index-num])
		num--
	}
	
	words = slices.Delete(words, index, index+2)
	index -= 2
	return words, index
}
