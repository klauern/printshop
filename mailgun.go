package printshop

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/mailgun/mailgun-go.v1"
)

// MailClient provides all of the necessary bits to connect to the MailGun API.
type MailClient struct {
	domain       string
	apiKey       string
	publicApiKey string
}

// Sender is an interface for sending emails.
type Sender interface {
	Send(email *Email) (bool, error)
}

func (c *MailClient) Send(email *Email) (bool, error) {
	mg := mailgun.NewMailgun(c.domain, c.apiKey, c.publicApiKey)
	body, err := email.RenderBody()
	if err != nil {
		return false, errors.Wrap(err, "Unable to get Email Body")
	}
	msg := mg.NewMessage(email.meta.from, email.meta.subject, body)
	resp, id, err := mg.Send(msg)
	if err != nil {
		return false, errors.Wrap(err, "Unable to Send Message")
	}
	log.Printf("ID: %s Resp: %s\n", id, resp)
	return true, nil
}
