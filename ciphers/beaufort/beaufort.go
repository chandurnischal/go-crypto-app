package beaufort

import (
	"errors"
	"strings"
	"unicode"
)

type Beaufort struct {
	alphabet, key string
}

func New(keyword string) (Beaufort, error) {
	b := Beaufort{alphabet: "abcdefghijklmnopqrstuvwxyz", key: keyword}
	b.key = b.removeSpecialCharacters(b.key)
	if len(b.key) < 1 {
		return b, errors.New("invalid key length")
	}
	return b, nil
}

func (b Beaufort) removeSpecialCharacters(text string) string {
	var t string

	for _, v := range text {
		if unicode.IsLetter(v) {
			t += string(strings.ToLower(string(v)))
		}
	}
	return t
}

func (b Beaufort) alphabetShifter(n int) string {
	var t string

	for i := range b.alphabet {
		t += string(b.alphabet[(i+n)%26])
	}
	return t
}

func (b Beaufort) keygen(text string) string {
	n := len(text)
	i := 0
	for len(b.key) != n {
		b.key += string(b.key[i])
		i++
	}
	return b.key
}

func (b Beaufort) EncAndDec(plaintext string) string {
	var ciphertext string
	plaintext = b.removeSpecialCharacters(plaintext)
	b.key = b.keygen(plaintext)

	for i, v := range plaintext {
		shifter := strings.IndexRune(b.alphabet, v)
		t := b.alphabetShifter(shifter)
		index := strings.IndexByte(t, b.key[i])
		ciphertext += string(b.alphabet[index])
	}

	return ciphertext
}

func Operation(message, key string) (string, error) {
	b, err := New(key)

	if err != nil {
		return "", err
	}

	return b.EncAndDec(message), nil
}
