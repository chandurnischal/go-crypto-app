package utils

import (
	"strings"
	"unicode"
)

func RemoveSpecialCharacters(text string) string {
	var t string

	for _, v := range text {
		if unicode.IsLetter(v) {
			t += string(strings.ToLower(string(v)))
		}
	}
	return t
}
