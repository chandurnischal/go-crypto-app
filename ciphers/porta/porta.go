package porta

import (
	"errors"
	"strings"
	"unicode"
)

type Porta struct {
	half1, half2, key string
}

func New(key string) (Porta, error) {
	p := Porta{half1: "abcdefghijklm", half2: "nopqrstuvwxyz", key: key}
	if len(key) < 1 {
		return p, errors.New("invalid key length")
	}

	p.key = p.removeSpecialCharacters(p.key)

	return p, nil
}

func (p Porta) removeSpecialCharacters(text string) string {
	var temp string
	for _, t := range text {
		if unicode.IsLetter(t) {
			temp += strings.ToLower(string(t))
		}
	}
	return temp
}

func (p Porta) keygen(text string) string {
	var temp string
	n := len(text)
	m := len(p.key)
	var i int

	for len(temp) != n {
		temp += string(p.key[i%m])
		i++
	}

	return temp
}

func (p Porta) alphabetShifter(r rune) string {
	alphabet := p.half1 + p.half2
	shifter := strings.IndexRune(alphabet, r)
	if shifter%2 != 0 {
		shifter = (shifter - 1) / 2
	} else {
		shifter = shifter / 2
	}
	var temp string
	for i := range p.half2 {
		temp += string(p.half2[(i+shifter)%13])
	}
	return temp
}

func (p Porta) EncAndDec(plaintext string) string {
	var ciphertext string
	plaintext = p.removeSpecialCharacters(plaintext)
	key := p.keygen(plaintext)
	for i, k := range key {
		alphabet := p.half1 + p.alphabetShifter(k)
		shiftedAlphabet := p.alphabetShifter(k) + p.half1
		index := strings.IndexByte(alphabet, plaintext[i])
		ciphertext += string(shiftedAlphabet[index])
	}
	return ciphertext
}

func Operation(message, keyword string) (string, error) {
	p, err := New(keyword)

	if err != nil {
		return "", err
	}

	return p.EncAndDec(message), nil
}
