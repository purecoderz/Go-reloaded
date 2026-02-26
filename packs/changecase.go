package packs

import "slices"

func ChangeCase(index int, words []string, action func(string) string) ([]string, int) {
	if index > 0 {
		words[index-1] = action(words[index-1])
		words = slices.Delete(words, index, index+1)
		index--
		return words, index
	}
	return words,index
}
