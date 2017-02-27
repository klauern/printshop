package printshop

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type MailClient struct {
	domain       string
	apiKey       string
	publicApiKey string
}

type Sender interface {
	Send(email *Email, client *MailClient) (bool, error)
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
