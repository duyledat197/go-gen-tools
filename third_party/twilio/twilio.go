package twilio

import (
	"os"
)

// Twilio SDK object
type Twilio struct {
	phoneNumber    string
	subaccountSid  string
	sendGridAPIKey string
}

// NewTwilio constructor
func NewTwilio() *Twilio {
	return &Twilio{
		phoneNumber:    os.Getenv("TWILIO_PHONE_NUMBER"),
		subaccountSid:  os.Getenv("TWILIO_SUBACCOUNT_SID"),
		sendGridAPIKey: os.Getenv("SENDGRID_API_KEY"),
	}
}

// Messaging get client
func (s *Twilio) Messaging() *MessagingClient {
	return NewMessagingClient(s.phoneNumber, s.subaccountSid)
}
