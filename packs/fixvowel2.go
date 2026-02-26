package packs

import (
	"unicode"
)

func FixVowels2(words []string) {
	vowel := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'h': true}
	for index, letterA := range words {
		if letterA == "a" || letterA == "A" {
			if index < len(words)-1 {
				runes := []rune(words[index+1])
				if vowel[unicode.ToLower(runes[0])] {
					words[index] += "n"
				}
			}
		}

	}

}
