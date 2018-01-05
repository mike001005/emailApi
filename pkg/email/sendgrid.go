package email

import (
	"github.com/sendgrid/rest"

	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridEmail Sends emails using the sendGrid  email api
// "<strong>and easy to do anywhere, even with Go</strong>"
func SendGridEmail(fromAddress string, toAddress string, mailSubject string, mailMessage string, html string) (*rest.Response, error) {
	from := mail.NewEmail("Example User", fromAddress)
	subject := mailSubject
	to := mail.NewEmail("Example User", toAddress)
	plainTextContent := mailMessage
	htmlContent := html
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	return response, err
}
