package foursquare

import (
	"errors"
	"strings"
	"unicode"
)

type FourSqaure struct {
	alphabet string
	key1     string
	key2     string
}

func New(key1, key2 string) (FourSqaure, error) {
	f := FourSqaure{alphabet: "abcdefghiklmnopqrstuvwxyz", key1: key1, key2: key2}

	if len(key1) < 1 || len(key2) < 1 {
		return f, errors.New("invalid key")
	}
	f.key1 = f.keygen(f.key1)
	f.key2 = f.keygen(f.key2)

	return f, nil
}

func (f FourSqaure) removeSpecialCharacters(text string) string {
	var res string

	for _, v := range text {
		if unicode.IsLetter(v) {
			res += string(v)
		}
	}

	return res
}

func (f FourSqaure) removeDuplicates(text string) string {
	var res string

	for _, v := range text {
		if !strings.ContainsRune(res, v) {
			res += string(v)
		}
	}

	return res
}

func (f FourSqaure) keygen(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "j", "i")
	text = f.removeSpecialCharacters(text)
	text = f.removeDuplicates(text)

	for _, v := range f.alphabet {
		if !strings.ContainsRune(text, v) {
			text += string(v)
		}
	}

	return text
}

func (f FourSqaure) findCoords(text string, r rune) [2]int {
	index := strings.IndexRune(text, r)
	return [2]int{index / 5, index % 5}
}

func (f FourSqaure) bigramEnc(bigram string) string {
	coords1 := f.findCoords(f.alphabet, rune(bigram[0]))
	coords2 := f.findCoords(f.alphabet, rune(bigram[1]))
	var res string
	res += string(f.key1[(coords1[0]*5)+(coords2[1])])
	res += string(f.key2[(coords2[0]*5)+(coords1[1])])
	return res
}

func (f FourSqaure) bigramDec(bigram string) string {
	coords1 := f.findCoords(f.key1, rune(bigram[0]))
	coords2 := f.findCoords(f.key2, rune(bigram[1]))
	var res string
	res += string(f.alphabet[(coords1[0]*5)+(coords2[1])])
	res += string(f.alphabet[(coords2[0]*5)+(coords1[1])])
	return res
}

func (f FourSqaure) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = strings.ToLower(plaintext)
	plaintext = f.removeSpecialCharacters(plaintext)
	n := len(plaintext)
	if n%2 != 0 {
		plaintext += string('x')
	}
	for i := 0; i < n; i = i + 2 {
		ciphertext += f.bigramEnc(plaintext[i : i+2])
	}

	return ciphertext
}

func (f FourSqaure) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = strings.ToLower(ciphertext)
	ciphertext = f.removeSpecialCharacters(ciphertext)
	n := len(ciphertext)
	if n%2 != 0 {
		ciphertext += string('x')
	}
	for i := 0; i < n; i = i + 2 {
		plaintext += f.bigramDec(ciphertext[i : i+2])
	}
	return plaintext
}

func Operation(message, key1, key2, operation string) (string, error) {
	f, err := New(key1, key2)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return f.Encrypt(message), nil
	}

	return f.Decrypt(message), nil
}
