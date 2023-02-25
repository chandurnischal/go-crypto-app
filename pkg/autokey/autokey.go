package autokey

import (
	"errors"
	"strings"
	"unicode"
)

type Autokey struct {
	alphabet, key string
}

func New(keyword string) (Autokey, error) {
	auto := Autokey{alphabet: "abcdefghijklmnopqrstuvwxyz", key: keyword}

	auto.key = auto.removeSpecialCharacters(auto.key)
	if len(auto.key) < 1 {
		return auto, errors.New("invalid key length")
	}

	return auto, nil
}

func (a Autokey) removeSpecialCharacters(text string) string {
	var t string

	for _, v := range text {
		if unicode.IsLetter(v) {
			t += string(strings.ToLower(string(v)))
		}
	}
	return t
}

func (a Autokey) alphabetShifter(n int) string {
	var t string

	for i := range a.alphabet {
		t += string(a.alphabet[(i+n)%26])
	}

	return t
}

func (a Autokey) encKeyGen(text string) string {

	n := len(text)
	i := 0
	for len(a.key) != n {
		a.key += string(text[i])
		i++
	}
	return a.key
}

func (a Autokey) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = a.removeSpecialCharacters(plaintext)
	a.key = a.encKeyGen(plaintext)

	for i := range a.key {
		shifter := strings.IndexByte(a.alphabet, a.key[i])
		index := strings.IndexByte(a.alphabet, plaintext[i])
		ciphertext += string(a.alphabetShifter(shifter)[index])
	}

	return ciphertext
}

func (a Autokey) Decrypt(ciphertext string) string {
	temp := a.key
	var plaintext string
	for i, v := range ciphertext {
		shifter := strings.IndexByte(a.alphabet, byte(a.key[i]))
		t := a.alphabetShifter(shifter)
		index := strings.IndexRune(t, v)
		temp += string(a.alphabet[index])
		plaintext += string(a.alphabet[index])

	}
	return plaintext
}

func Operation(message, key, operation string) (string, error) {
	auto, err := New(key)

	if operation == "Encrypt" {
		return auto.Encrypt(message), err
	}
	return auto.Decrypt(message), err
}
