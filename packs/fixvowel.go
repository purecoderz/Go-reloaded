package packs

import (
	"strings"
	"unicode"
)

// func FixVowels(words []string) {
// 	vowel := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'h': true}
// 	for index, letterA := range words {
// 		if len(letterA) > 1 {
// 			continue
// 		}
// 		if strings.ToLower(letterA) != "a" {
// 			continue
// 		}
// 		if index < len(words)-1 {
// 			runes := []rune(words[index+1])
// 			if vowel[unicode.ToLower(runes[0])] {
// 				words[index] += "n"
// 			}
// 		}

// 	}

// }

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

		runes := []rune(words[index+1])

		if vowel[unicode.ToLower(runes[0])] {
			words[index] += "n"
		}

	}
}
