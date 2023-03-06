package authenticate

type Authenticator interface {
	Generate(payload *Payload) (*Token, error)
	Verify(token string) (*Payload, error)
}
