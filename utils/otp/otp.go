package otp

type OTPAuthenticator interface {
	Generate(payload *Payload) (*Token, error)
	Verify(payload *Payload, token string) error
}

type Token struct {
	Token string
	Uri   string
	Json  string
}

type Payload struct {
	UserID   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Key      string `bson:"key,omitempty" json:"key,omitempty"`
	Issuer   string `bson:"issuer,omitempty" json:"issuer,omitempty"`
}
