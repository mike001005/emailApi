package email

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

// InterfaceMailGun mailGun interface
type InterfaceMailGun interface {
	SendSimpleMessage(from string, toEmail string) (string, string, error)
}

// SendSimpleMessage Sends an email using Mail gun api
func SendSimpleMessage(from string, toEmail string) (string, string, error) {
	mg := mailgun.NewMailgun(MgDomain, MgAPIKey, MgPublicAPIKey)
	msg := mailgun.NewMessage(from, "Message subject", "This is the message body", toEmail)
	response, id, err := mg.Send(msg)
	return response, id, err
}
