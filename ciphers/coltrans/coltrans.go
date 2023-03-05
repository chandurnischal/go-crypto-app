package coltrans

import (
	"errors"
	"sort"
	"strings"
	"unicode"
)

type ColTrans struct {
	key, sortedKey string
	padder         byte
}

func New(keyword string, pad byte) (ColTrans, error) {
	c := ColTrans{key: keyword, padder: pad}
	c.key = c.removeDuplicateRunes()

	if len(keyword) < 1 {
		return c, errors.New("invalid key length")
	}

	if !unicode.IsLetter(rune(pad)) {
		return c, errors.New("invalid padding character")
	}

	c.sortedKey = c.sortString(c.key)
	return c, nil
}

func (c ColTrans) removeDuplicateRunes() string {
	var t string
	for _, v := range c.key {
		if !strings.ContainsRune(t, v) {
			t += string(v)
		}
	}
	return t
}

func (c ColTrans) removeSpecialCharacters(text string) string {
	var t string

	for _, v := range text {
		if unicode.IsLetter(v) {
			t += string(strings.ToLower(string(v)))
		}
	}
	return t
}

func (c ColTrans) sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func (c ColTrans) textPadder(text string) string {
	keylen := len(c.key)

	for len(text)%keylen != 0 {
		text += string(c.padder)
	}
	return text
}

func (c ColTrans) indicesOfRune(text string, r rune) []int {
	indices := []int{}

	for i, v := range text {
		if v == r {
			indices = append(indices, i)
		}
	}

	return indices
}

func (c ColTrans) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = c.removeSpecialCharacters(plaintext)
	plaintext = c.textPadder(plaintext)
	repeatString := strings.Repeat(c.key, len(plaintext)/len(c.key))

	for _, v := range c.sortedKey {
		indices := c.indicesOfRune(repeatString, v)
		for _, index := range indices {
			ciphertext += string(plaintext[index])
		}
	}
	return ciphertext
}

func (c ColTrans) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = c.removeSpecialCharacters(ciphertext)
	ciphertext = c.textPadder(ciphertext)
	n := len(ciphertext)
	var repeatString string
	for _, v := range c.sortedKey {
		repeatString += strings.Repeat(string(v), n/len(c.sortedKey))
	}

	decryptIndices := [][]int{}

	for _, v := range c.key {
		decryptIndices = append(decryptIndices, c.indicesOfRune(repeatString, v))
	}
	m := len(decryptIndices[0])
	n = len(decryptIndices)

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			plaintext += string(ciphertext[decryptIndices[i][j]])
		}
	}

	return plaintext
}

func Operation(pad byte, message, key, operation string) (string, error) {
	col, err := New(key, pad)

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return col.Encrypt(message), nil
	}

	return col.Decrypt(message), nil

}
