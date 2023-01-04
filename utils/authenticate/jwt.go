package authenticate

import (
	"errors"
	"fmt"
	"time"

	"github.com/reddit/jwt-go"
)

type JWTAuthenticator struct {
	secretKey      string
	expirationTime time.Duration
}

func NewJWTAuthenticator(secretKey string, expirationTime time.Duration) (Authenticator, error) {
	return &JWTAuthenticator{
		secretKey,
		expirationTime,
	}, nil
}

func (a *JWTAuthenticator) Generate(payload *Payload) (*Token, error) {
	payload.AddExpired(a.expirationTime)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(a.secretKey))
	if err != nil {
		return nil, fmt.Errorf("")
	}

	return &Token{
		Token:     token,
		IssueAt:   payload.IssueAt,
		ExpiredAt: payload.ExpiredAt,
	}, nil

}

func (a *JWTAuthenticator) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("")
		}
		return []byte(a.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, fmt.Errorf("")) {
			return nil, fmt.Errorf("")
		}
		return nil, fmt.Errorf("")
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, fmt.Errorf("")
	}
	return payload, nil
}
