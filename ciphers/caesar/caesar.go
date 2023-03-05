package caesar

import "cryptoapp/ciphers/affine"

type Caesar struct {
	affine.Affine
}

func New(n int) (Caesar, error) {
	aff, err := affine.New(1, n%26)
	return Caesar{aff}, err
}

func Operation(offset int, message, operation string) (string, error) {
	c, err := New(offset)
	if err != nil {
		return "", err
	}
	if operation == "Encrypt" {
		return c.Encrypt(message), nil
	}
	return c.Decrypt(message), nil
}
