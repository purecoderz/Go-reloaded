package main

import (
	"reload/packs"
	"strings"
)

func FormatWords2(words []string) []string {
	for index := 0; index < len(words); index++ {
		switch words[index] {
		case "(cap)":
			words, index = packs.ChangeCase(index, words, packs.ToTitleCase)
		case "(cap,":
			words, index = packs.ChangeMultipleCase(index, words, packs.ToTitleCase)
		case "(up)":
			words, index = packs.ChangeCase(index, words, strings.ToUpper)
		case "(up,":
			words, index = packs.ChangeMultipleCase(index, words, strings.ToUpper)
		case "(low)":
			words, index = packs.ChangeCase(index, words, strings.ToLower)
		case "(low,":
			words, index = packs.ChangeMultipleCase(index, words, strings.ToLower)
		case "(hex)":
			words, index = packs.NumConversion(16, index, words)
		case "(bin)":
			words, index = packs.NumConversion(2, index, words)
		}

	}

	// fixed vowels
	packs.FixVowels2(words)

	start := -1
	end := -1
	for index := 0; index < len(words); index++ {
		switch words[index] {
		case "'s", "'ve", "'t", "'re":
			words, index = packs.Apostrophe(words, index)
		case "'":
			words, index, start, end = packs.Quotes(words, index, start, end)
		default:
			// If it's not a quote or apostrophe, snap standard punctuation!
			words, index = packs.CheckPunctuation(index, words)
		}
	}

	return words

}
