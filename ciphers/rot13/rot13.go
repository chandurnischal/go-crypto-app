package rot13

import "cryptoapp/ciphers/affine"

type ROT13 struct {
	affine.Affine
}

func New() (ROT13, error) {
	aff, err := affine.New(1, 13)

	return ROT13{aff}, err
}

func Operation(message, operation string) (string, error) {
	r, err := New()

	if err != nil {
		return "", err
	}

	if operation == "Encrypt" {
		return r.Encrypt(message), nil
	}
	return r.Decrypt(message), nil
}
