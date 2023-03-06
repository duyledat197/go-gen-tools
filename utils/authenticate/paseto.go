package authenticate

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoAuthenticator struct {
	paseto         *paseto.V2
	symmetricKey   []byte
	expirationTime time.Duration
}

func NewPasetoAuthenticator(symmetricKey string, expirationTime time.Duration) (Authenticator, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("")
	}
	return &PasetoAuthenticator{
		paseto:         paseto.NewV2(),
		symmetricKey:   []byte(symmetricKey),
		expirationTime: expirationTime,
	}, nil
}

func (a *PasetoAuthenticator) Generate(payload *Payload) (*Token, error) {
	payload.AddExpired(a.expirationTime)
	token, err := a.paseto.Encrypt(a.symmetricKey, payload, nil)

	if err != nil {
		return nil, fmt.Errorf("")
	}

	return &Token{
		Token:     token,
		ExpiredAt: payload.ExpiredAt,
		IssueAt:   payload.IssueAt,
	}, nil
}

func (a *PasetoAuthenticator) Verify(token string) (*Payload, error) {
	payload := &Payload{}

	if err := a.paseto.Decrypt(token, a.symmetricKey, payload, nil); err != nil {
		return nil, fmt.Errorf("")
	}

	if err := payload.Valid(); err != nil {
		return nil, err
	}

	return payload, nil
}
