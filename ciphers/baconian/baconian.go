package baconian

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Baconian struct {
	alphabet string
}

func New() (Baconian, error) {
	bac := Baconian{alphabet: "abcdefghiklmnopqrstuvwxyz"}
	return bac, nil
}

func (b Baconian) binToDec(bin string) int {
	n := len(bin)
	var res int

	for i := n - 1; i >= 0; i-- {
		dig, _ := strconv.Atoi(string(bin[i]))
		res += dig * int(math.Pow(2, float64(n-i-1)))
	}

	return res
}

func (b Baconian) Encrypt(plaintext string) string {
	var ciphertext string

	for _, v := range plaintext {
		if unicode.IsLetter(v) {
			index := strings.Index(b.alphabet, strings.ToLower(string(v)))
			ciphertext += fmt.Sprintf("%05b", index)
		}
	}

	return ciphertext
}

func (b Baconian) Decrypt(ciphertext string) string {
	var plaintext string

	if len(ciphertext)%5 != 0 {
		return ciphertext
	}
	n := len(ciphertext)

	for i := 0; i < n; i = i + 5 {
		plaintext += string(b.alphabet[b.binToDec(ciphertext[i:i+5])])
	}

	return plaintext
}

func Operation(message, operation string) (string, error) {
	bac, err := New()
	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return bac.Encrypt(message), nil
	}
	return bac.Decrypt(message), nil
}
