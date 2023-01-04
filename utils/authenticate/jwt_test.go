package authenticate

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/duyledat197/go-gen-tools/utils"
	"github.com/reddit/jwt-go"
	"github.com/stretchr/testify/assert"
)

var info = &Payload{
	UserID:    faker.UUIDDigit(),
	UserName:  faker.Username(),
	Ip:        faker.IPv4(),
	UserAgent: faker.MacAddress(),
}

func Test_JWT(t *testing.T) {
	authenticator, err := NewJWTAuthenticator(utils.RandStringBytes(32), time.Minute)

	assert.NoError(t, err)

	token, err := authenticator.Generate(info)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.Token)
	assert.NotEmpty(t, token.ExpiredAt)

	payload, err := authenticator.Verify(token.Token)

	assert.NoError(t, err)
	assert.NotNil(t, payload)
	assert.Equal(t, payload.UserID, info.UserID)
	assert.Equal(t, payload.UserName, info.UserName)
	assert.Equal(t, payload.Ip, info.Ip)
	assert.Equal(t, payload.UserAgent, info.UserAgent)
	assert.WithinDuration(t, payload.ExpiredAt, token.ExpiredAt, time.Second)
	assert.WithinDuration(t, payload.IssueAt, token.IssueAt, time.Second)
}

func Test_JWT_ExpiredToken(t *testing.T) {
	authenticator, err := NewJWTAuthenticator(utils.RandStringBytes(32), -time.Minute)

	assert.NoError(t, err)

	token, err := authenticator.Generate(info)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.Token)
	assert.NotEmpty(t, token.ExpiredAt)

	payload, err := authenticator.Verify(token.Token)
	assert.Error(t, err)
	assert.Equal(t, err, nil)
	assert.Nil(t, payload)
}

func Test_JWT_InvalidToken(t *testing.T) {
	authenticator, err := NewJWTAuthenticator(utils.RandStringBytes(32), time.Minute)

	assert.NoError(t, err)

	token, err := authenticator.Generate(info)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.Token)
	assert.NotEmpty(t, token.ExpiredAt)

	payload := &Payload{}
	payload.AddExpired(time.Minute)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tknString, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	p, err := authenticator.Verify(tknString)
	assert.Error(t, err)
	assert.Equal(t, err, nil)
	assert.Nil(t, p)
}
