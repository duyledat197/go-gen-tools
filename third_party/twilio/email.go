package twilio

import (
	"fmt"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailClient class
type EmailClient struct {
	client *sendgrid.Client
	from   *mail.Email
}

// NewEmailClient constructor
func NewEmailClient(sendGridAPIKey string, from *mail.Email) *EmailClient {
	return &EmailClient{
		client: sendgrid.NewSendClient(sendGridAPIKey),
		from:   from,
	}
}

// Send Email function
func (s *EmailClient) Send(to *mail.Email, templateID, subject string, personalizations ...*mail.Personalization) (resp *rest.Response, err error) {
	message := mail.NewV3MailInit(s.from, subject, to)
	message.SetTemplateID(templateID)
	message.AddPersonalizations(personalizations...)
	resp, err = s.client.Send(message)
	if err != nil {
		err = fmt.Errorf("client.Send: %w", err)
	}
	return
}
