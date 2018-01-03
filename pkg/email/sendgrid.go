package email

import (
	"github.com/sendgrid/rest"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

// SendGridEmail Sends emails using the sendGrid  email api
func SendGridEmail(fromAddress string, toAddress string) (*rest.Response, error) {
	from := mail.NewEmail("Example User", fromAddress)
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", toAddress)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("sendGridAPIKey"))
	response, err := client.Send(message)

	return response, err
}
