package polybius

import (
	"errors"
	"strings"
	"unicode"
)

type Polybius struct {
	alphabet string
	key      string
	chars    string
}

func New(key, chars string) (Polybius, error) {
	p := Polybius{alphabet: "abcdefghiklmnopqrstuvwxyz", key: key, chars: chars}
	p.key = strings.ToLower(p.key)
	p.key = p.removeSpecialCharacters(p.key)
	p.key = p.removeDuplicates(p.key)
	p.key = p.keygen(p.key)
	p.chars = p.removeDuplicates(p.chars)

	if len(p.key) == 0 {
		return p, errors.New("invalid key")
	}

	if len(p.chars) != 5 {
		return p, errors.New("invalid ciphertext characters")
	}

	return p, nil
}

func (p Polybius) removeSpecialCharacters(text string) string {
	var res string

	for _, v := range text {
		if unicode.IsLetter(v) {
			res += string(v)
		}
	}
	return res
}

func (p Polybius) removeDuplicates(text string) string {
	var res string

	for _, v := range text {
		if !strings.ContainsRune(res, v) {
			res += string(v)
		}
	}

	return res
}

func (p Polybius) keygen(text string) string {

	for _, v := range p.alphabet {
		if !strings.ContainsRune(text, v) {
			text += string(v)
		}
	}

	return text
}

func (p Polybius) findCoords(text string, r rune) [2]int {
	index := strings.IndexRune(text, r)
	return [2]int{index / 5, index % 5}
}

func (p Polybius) charEnc(r rune) string {
	coords := p.findCoords(p.key, r)
	return string(p.chars[coords[0]]) + string(p.chars[coords[1]])
}

func (p Polybius) charDec(bigram string) string {
	x := strings.IndexByte(p.chars, bigram[0])
	y := strings.IndexByte(p.chars, bigram[1])
	return string(p.key[(x*5)+y])
}

func (p Polybius) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = strings.ToLower(plaintext)
	plaintext = p.removeSpecialCharacters(plaintext)

	for _, v := range plaintext {
		ciphertext += p.charEnc(v)
	}

	return ciphertext
}

func (p Polybius) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = p.removeSpecialCharacters(ciphertext)

	if len(ciphertext)%2 != 0 {
		ciphertext = ciphertext[:len(ciphertext)-1]
	}
	n := len(ciphertext)

	for i := 0; i < n; i = i + 2 {
		plaintext += p.charDec(ciphertext[i : i+2])
	}

	return plaintext
}

func Operation(message, key, chars, operation string) (string, error) {
	p, err := New(key, chars)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return p.Encrypt(message), nil
	}

	return p.Decrypt(message), nil
}
