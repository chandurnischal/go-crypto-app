package simplesubstitution

import (
	"errors"
	"strings"
	"unicode"
)

type SimpleSubstitution struct {
	alphabet string
	key      string
}

func New(key string) (SimpleSubstitution, error) {
	s := SimpleSubstitution{alphabet: "abcdefghijklmnopqrstuvwxyz", key: key}
	s.key = s.keygen(s.key)
	if s.key == s.alphabet {
		return s, errors.New("invalid key")
	}
	return s, nil
}

func (s SimpleSubstitution) removeSpecialCharacters(text string) string {
	var res string

	for _, v := range text {
		if unicode.IsLetter(v) {
			res += string(v)
		}
	}

	return res
}

func (s SimpleSubstitution) removeDuplicates(text string) string {
	var res string

	for _, v := range text {
		if !strings.ContainsRune(res, v) {
			res += string(v)
		}
	}

	return res
}

func (s SimpleSubstitution) keygen(text string) string {
	text = strings.ToLower(text)
	text = s.removeSpecialCharacters(text)
	text = s.removeDuplicates(text)

	for _, v := range s.alphabet {
		if !strings.ContainsRune(text, v) {
			text += string(v)
		}
	}
	return text
}

func (s SimpleSubstitution) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = strings.ToLower(plaintext)

	for _, v := range plaintext {
		if unicode.IsLetter(v) {
			index := strings.IndexRune(s.alphabet, v)
			ciphertext += string(s.key[index])
		} else {
			ciphertext += string(v)
		}
	}
	return ciphertext
}

func (s SimpleSubstitution) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = strings.ToLower(ciphertext)

	for _, v := range ciphertext {
		if unicode.IsLetter(v) {
			index := strings.IndexRune(s.key, v)
			plaintext += string(s.alphabet[index])
		} else {
			plaintext += string(v)
		}
	}

	return plaintext
}

func Operation(message, key, operation string) (string, error) {
	s, err := New(key)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return s.Encrypt(message), nil
	}
	return s.Decrypt(message), nil
}
