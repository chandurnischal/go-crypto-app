package vigenere

import (
	"errors"
	"strings"
	"unicode"
)

type Vigenere struct {
	alphabet, key string
}

func New(keyword string) (Vigenere, error) {
	v := Vigenere{alphabet: "abcdefghijklmnopqrstuvwxyz", key: keyword}
	v.key = v.removeSpecialCharacters(v.key)
	if len(v.key) < 1 {
		return v, errors.New("invalid key length")
	}
	return v, nil
}

func (v Vigenere) removeSpecialCharacters(text string) string {
	var t string

	for _, v := range text {
		if unicode.IsLetter(v) {
			t += strings.ToLower(string(v))
		}
	}

	return t
}

func (v Vigenere) alphabetShifter(n int) string {
	var t string
	m := len(v.alphabet)
	for i := range v.alphabet {
		t += string(v.alphabet[(i+n)%m])
	}

	return t
}

func (v Vigenere) keygen(text string) string {
	var temp string
	n := len(text)
	m := len(v.key)
	i := 0

	for len(temp) != n {
		temp += string(v.key[i%m])
		i++
	}
	return temp
}

func (v Vigenere) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = v.removeSpecialCharacters(plaintext)
	tempKey := v.keygen(plaintext)

	for i, p := range plaintext {
		shifter := strings.IndexRune(v.alphabet, p)
		shiftedAlpha := v.alphabetShifter(shifter)
		index := strings.IndexByte(v.alphabet, tempKey[i])
		ciphertext += string(shiftedAlpha[index])
	}
	return ciphertext
}

func (v Vigenere) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = v.removeSpecialCharacters(ciphertext)
	tempKey := v.keygen(ciphertext)

	for i, c := range ciphertext {
		shifter := strings.IndexByte(v.alphabet, tempKey[i])
		shiftedAlpha := v.alphabetShifter((shifter))
		index := strings.IndexRune(shiftedAlpha, c)
		plaintext += string(v.alphabet[index])

	}

	return plaintext
}

func Operation(message, key, operation string) (string, error) {
	v, err := New(key)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return v.Encrypt(message), nil
	}
	return v.Decrypt(message), nil

}
