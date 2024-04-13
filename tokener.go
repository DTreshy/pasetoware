package pasetoware

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type Tokener struct {
	TokenPrivateKey paseto.V4AsymmetricSecretKey
	TokenPublicKey  paseto.V4AsymmetricPublicKey
}

func NewTokener() *Tokener {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	return &Tokener{
		TokenPrivateKey: privateKey,
		TokenPublicKey:  publicKey,
	}
}

func (t *Tokener) GenerateToken(m map[string]string, expiration time.Duration) string {
	token := paseto.NewToken()
	currTime := time.Now()

	token.SetIssuedAt(currTime)
	token.SetNotBefore(currTime)
	token.SetExpiration(currTime.Add(expiration))

	for key, val := range m {
		token.SetString(key, val)
	}

	return token.V4Sign(t.TokenPrivateKey, nil)
}

func (t *Tokener) ParseToken(token string) (*paseto.Token, error) {
	parser := paseto.MakeParser([]paseto.Rule{paseto.ValidAt(time.Now())})

	parsedToken, err := parser.ParseV4Public(t.TokenPublicKey, token, nil)
	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}
