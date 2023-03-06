package otp

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jltorresm/otpgo"
	"github.com/jltorresm/otpgo/config"
)

type TOTPAuthenticator struct {
	totp *otpgo.TOTP
}

func NewTOTPAuthenticator(period time.Duration) (OTPAuthenticator, error) {
	return &TOTPAuthenticator{
		&otpgo.TOTP{
			Period:    int(period.Seconds()),
			Algorithm: config.HmacSHA256,
			Length:    config.Length6,
		},
	}, nil
}

func (a *TOTPAuthenticator) Generate(payload *Payload) (*Token, error) {
	a.totp.Key = payload.Key
	otp, err := a.totp.Generate()
	if err != nil {
		return nil, err
	}
	uri := a.totp.KeyUri(payload.UserName, payload.Issuer)
	b, err := json.Marshal(a.totp)
	if err != nil {
		return nil, err
	}
	return &Token{
		Token: otp,
		Uri:   uri.String(),
		Json:  string(b),
	}, nil
}

func (a *TOTPAuthenticator) Verify(payload *Payload, token string) error {
	a.totp.Key = payload.Key
	isValid, err := a.totp.Validate(token)

	if !isValid {
		return fmt.Errorf("otp is not valid")
	}

	if err != nil {
		return err
	}

	return nil
}
