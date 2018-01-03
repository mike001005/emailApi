package email

import (
	"gopkg.in/mailgun/mailgun-go.v1"
	"os"
)

// SendSimpleMessage Sends an email using Mail gun api
func SendSimpleMessage(from string, toEmail string) (string, string, error) {
	mg := mailgun.NewMailgun(os.Getenv("MgDomain"), os.Getenv("MgAPIKey"), os.Getenv("MgPublicAPIKey"))
	msg := mailgun.NewMessage(from, "Message subject", "This is the message body", toEmail)
	response, id, err := mg.Send(msg)
	return response, id, err
}
