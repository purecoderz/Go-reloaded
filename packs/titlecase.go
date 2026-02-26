package packs

import (
	"unicode"
)

func ToTitleCase(param string) string {
	if param == "" {
		return ""
	}
	runes := []rune(param)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
