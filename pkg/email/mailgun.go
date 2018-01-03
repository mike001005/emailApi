package email

import (
	"os"

	"gopkg.in/mailgun/mailgun-go.v1"
)

// SendSimpleMessage Sends an email using Mail gun api
func SendSimpleMessage(from string, toEmail string) (string, string, error) {
	mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
	msg := mailgun.NewMessage(from, "Message subject", "This is the message body", toEmail)
	response, id, err := mg.Send(msg)
	return response, id, err
}
