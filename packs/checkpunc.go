package packs

import "slices"

func CheckPunctuation(index int, word []string) ([]string, int) {
	punc := map[rune]bool{'.': true, ',': true, '!': true, '?': true, ':': true, ';': true}
	runes := []rune(word[index])
	// ,
	// harod
	// ?.,jide

	if len(runes) == 0 || !punc[runes[0]] || index == 0 {
		return word, index
	}


	for i, char := range runes {
		if !punc[char] {
			word[index-1] += string(runes[:i])
			word[index] = string(runes[i:])
			return word, index
		}
	}

	word[index-1] += word[index]
	word = slices.Delete(word, index, index+1)
	index--
	return word, index
}
