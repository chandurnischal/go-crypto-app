package affine

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Affine struct {
	alphabet           string
	multiplier, offset int
}

func New(multiplier, offset int) (Affine, error) {
	a := Affine{
		alphabet:   "abcdefghijklmnopqrstuvwxyz",
		multiplier: multiplier,
		offset:     offset,
	}
	if multiplier < 1 {
		return a, errors.New("invalid multiplier")
	}
	if offset < 0 {
		return a, errors.New("invalid offset")
	}
	return a, nil
}

func (a *Affine) gcd(p, q int) int {
	for q != 0 {
		r := p % q
		p = q
		q = r
	}
	return p
}

func (a *Affine) relativePrime(p, q int) bool {
	return a.gcd(p, q) == 1
}

func (a *Affine) keygen() string {
	var key string
	if !a.relativePrime(a.multiplier, len(a.alphabet)) {
		a.multiplier = 1
	}
	a.offset %= len(a.alphabet)
	for i := range a.alphabet {
		key += string(a.alphabet[((a.multiplier*i)+a.offset)%len(a.alphabet)])
	}
	return key
}

func (a *Affine) Encrypt(plaintext string) string {
	var ciphertext string
	plaintext = strings.ToLower(plaintext)
	key := a.keygen()
	for _, v := range plaintext {
		if unicode.IsLetter(v) {
			index := strings.IndexRune(a.alphabet, v)
			ciphertext += string(key[index])
		} else {
			ciphertext += string(v)
		}
	}
	return ciphertext
}

func (a *Affine) Decrypt(ciphertext string) string {
	var plaintext string
	ciphertext = strings.ToLower(ciphertext)
	key := a.keygen()

	for _, v := range ciphertext {
		if unicode.IsLetter(v) {
			index := strings.IndexRune(key, v)
			plaintext += string(a.alphabet[index])
		} else {
			plaintext += string(v)
		}
	}

	return plaintext
}

func Operation(m, o, message, operation string) (string, error) {
	multiplier, err := strconv.Atoi(m)

	if err != nil {
		return "wrong multiplier", errors.New("invalid multiplier")
	}
	offset, err := strconv.Atoi(o)

	if err != nil {
		return "wrong offset", errors.New("invalid offset")
	}

	aff, err := New(multiplier, offset)
	if operation == "Encrypt" {
		return aff.Encrypt(message), err
	}
	return aff.Decrypt(message), err
}
