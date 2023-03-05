package playfair

import (
	"errors"
	"strings"
	"unicode"
)

type Playfair struct {
	alphabet, key string
}

func New(keyword string) (Playfair, error) {
	p := Playfair{alphabet: "abcdefghiklmnopqrstuvwxyz", key: keyword}

	if len(keyword) < 1 {
		return p, errors.New("invalid key length")
	}
	p.key = p.removeDuplicates(p.key)
	p.key = p.removeSpecialCharacters(p.key)
	p.key = p.keygen()

	return p, nil
}

func (p Playfair) removeSpecialCharacters(text string) string {
	var t string
	for _, v := range text {
		if unicode.IsLetter(v) {
			t += strings.ToLower(string(v))
		}
	}
	return t
}

func (p Playfair) removeDuplicates(text string) string {
	var t string

	for _, v := range text {
		if !strings.ContainsRune(t, v) {
			t += string(v)
		}
	}

	return t
}

func (p Playfair) keygen() string {
	for _, v := range p.alphabet {
		if !strings.ContainsRune(p.key, v) {
			p.key += string(v)
		}
	}
	return p.key
}

func (p Playfair) getCoords(b byte, text string) [2]int {
	index := strings.IndexByte(text, b)

	return [2]int{index / 5, index % 5}
}

func (p Playfair) getPrevious(n int) int {
	n = n - 1
	if n < 0 {
		return 4
	}
	return n
}

func (p Playfair) bigramDecrypt(bigram string) string {
	var result string

	coords1 := p.getCoords(bigram[0], p.key)
	coords2 := p.getCoords(bigram[1], p.key)

	if coords1[1] == coords2[1] {
		// same column
		result += string(p.key[(p.getPrevious(coords1[0])*5)+coords1[1]])
		result += string(p.key[(p.getPrevious(coords2[0])*5)+coords2[1]])
		return result
	}

	if coords1[0] == coords2[0] {
		// same row
		result += string(p.key[(coords1[0]*5)+p.getPrevious(coords1[1])])
		result += string(p.key[(coords2[0]*5)+p.getPrevious(coords2[1])])
		return result
	}
	// square
	result += string(p.key[(coords1[0]*5)+coords2[1]])
	result += string(p.key[(coords2[0]*5)+coords1[1]])
	return result
}

func (p Playfair) getNext(n int) int {
	n = n + 1
	if n > 4 {
		return 0
	}
	return n
}

func (p Playfair) bigramEncrypt(bigram string) string {
	var result string

	coords1 := p.getCoords(bigram[0], p.key)
	coords2 := p.getCoords(bigram[1], p.key)

	if coords1[1] == coords2[1] {
		// same column
		result += string(p.key[(p.getNext(coords1[0])*5)+coords1[1]])
		result += string(p.key[(p.getNext(coords2[0])*5)+coords2[1]])
		return result
	}

	if coords1[0] == coords2[0] {
		// same row
		result += string(p.key[(coords1[0]*5)+p.getNext(coords1[1])])
		result += string(p.key[(coords2[0]*5)+p.getNext(coords2[1])])
		return result
	}

	result += string(p.key[(coords1[0]*5)+coords2[1]])
	result += string(p.key[(coords2[0]*5)+coords1[1]])
	return result
}

func (p Playfair) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = p.removeSpecialCharacters(plaintext)
	plaintext = strings.ReplaceAll(plaintext, "j", "i")

	if len(plaintext)%2 != 0 {
		plaintext += string('x')
	}

	n := len(plaintext)
	for i := 0; i < n; i = i + 2 {
		ciphertext += p.bigramEncrypt(plaintext[i : i+2])
	}

	return ciphertext
}

func (p Playfair) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = p.removeSpecialCharacters(ciphertext)
	n := len(ciphertext)

	for i := 0; i < n; i = i + 2 {
		plaintext += p.bigramDecrypt(ciphertext[i : i+2])
	}

	return plaintext
}

func Operation(message, key, operation string) (string, error) {
	p, err := New(key)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return p.Encrypt(message), nil
	}

	return p.Decrypt(message), nil

}
