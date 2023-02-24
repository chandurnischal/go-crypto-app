package atbash

import (
	"strings"
	"unicode"
)

type Atbash struct {
	alphabet string
}

func New() Atbash {
	return Atbash{alphabet: "abcdefghijklmnopqrstuvwxyz"}
}

func (a *Atbash) EncAndDec(text string) string {
	var res string
	text = strings.ToLower(text)
	n := len(a.alphabet)
	for _, v := range text {
		if unicode.IsLetter(v) {
			index := strings.IndexRune(a.alphabet, v)
			res += string(a.alphabet[n-index-1])
		} else {
			res += string(v)
		}
	}
	return res
}

func Operation(text string) (string, error) {
	at := New()
	return at.EncAndDec(text), nil
}
