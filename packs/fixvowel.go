package packs

import (
	"strings"
	"unicode"
)

func FixVowels(words []string, index int) {
	vowel := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'h': true}

	if index < len(words)-1 {
		n := 1
		for index+n < len(words)-1 {
			if strings.HasPrefix(words[index+n], "(") || strings.HasSuffix(words[index+n], ")") {
				n++
			} else {
				break
			}
		}
		runes := []rune(words[index+n])

		if vowel[unicode.ToLower(runes[0])] {
			words[index] += "n"
		}
	}

}
