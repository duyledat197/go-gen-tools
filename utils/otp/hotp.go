package otp

import (
	"encoding/json"
	"fmt"

	"github.com/jltorresm/otpgo"
	"github.com/jltorresm/otpgo/config"
)

type HOTPAuthenticator struct {
	hotp *otpgo.HOTP
}

func NewHOTPAuthenticator(counter, leeway uint64) (OTPAuthenticator, error) {
	return &HOTPAuthenticator{
		&otpgo.HOTP{
			Counter:   counter,
			Leeway:    leeway,
			Algorithm: config.HmacSHA256,
			Length:    config.Length6,
		},
	}, nil
}

func (a *HOTPAuthenticator) Generate(payload *Payload) (*Token, error) {
	a.hotp.Key = payload.Key
	otp, err := a.hotp.Generate()
	if err != nil {
		return nil, err
	}
	uri := a.hotp.KeyUri(payload.UserName, payload.Issuer)
	b, err := json.Marshal(a.hotp)
	if err != nil {
		return nil, err
	}
	return &Token{
		Token: otp,
		Uri:   uri.String(),
		Json:  string(b),
	}, nil
}

func (a *HOTPAuthenticator) Verify(payload *Payload, token string) error {
	a.hotp.Key = payload.Key
	isValid, err := a.hotp.Validate(token)

	if !isValid {
		return fmt.Errorf("otp is not valid")
	}

	if err != nil {
		return err
	}

	return nil
}
