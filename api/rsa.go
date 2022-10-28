package api

import (
	"crypto/rand"
	"crypto/rsa"
	"fwork/exceptions"
)

func GeneratePrivateKey(keySize int) (*rsa.PrivateKey, error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceNotFoundCode)
		e.SetMessage(exceptions.ResourceNotFoundMessage)
		e.Include(exceptions.Data{
			Value: err.Error(),
		})

		return nil, e.Exception()
	}

	return rsaKey, nil
}
