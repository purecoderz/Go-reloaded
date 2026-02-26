package packs

import (
	"slices"
	"strconv"
)

func NumConversion(num, index int, words []string) ([]string, int) {
	if index > 0 {
		res, err := strconv.ParseInt(words[index-1], num, 64)
		if err != nil {
			return words, index
		}
		txt := strconv.Itoa(int(res))
		words[index-1] = txt
		words = slices.Delete(words, index, index+1)
		index--
	}

	return words, index

}
