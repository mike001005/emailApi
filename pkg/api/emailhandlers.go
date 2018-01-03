package api

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/mike001005/emailApi/pkg/email"

)

// Emails are the arguments for Emails
type Emails struct {
	Senders   []string `json:"senders"`
	Receivers []string `json:"receivers"`
}

// HandleMailGun MailGun APi
func HandleMailGun(w http.ResponseWriter, r *http.Request) {

	data := Emails{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for z := 0; z < len(data.Senders); z++ {
		for i := 0; i < len(data.Receivers); i++ {
			fmt.Println("Sending email with mailgun...")
			fmt.Sprintf("%s: %s", data.Senders[z], data.Receivers[i])
			response, id, err := email.SendSimpleMessage(data.Senders[z], data.Receivers[i])
			fmt.Println(id)
			if err != nil {
				fmt.Printf("mail gun failed:%s", err)
				fmt.Println("sending email with sendgrid..")
				sgResponse, err := email.SendGridEmail(data.Senders[z], data.Receivers[i])
				if err != nil {
					fmt.Printf("send grid failed: %e, exiting program", err)
				} else {
					fmt.Println(sgResponse.StatusCode)
					fmt.Println(sgResponse.Body)
					fmt.Println(sgResponse.Headers)
					fmt.Println("Succsess with SendGrid.")

				}
			}
			fmt.Printf("email was sent. %s:%s \n", id, response)
		}
	}
}

// HandleSendGrid SendGrid API
func HandleSendGrid(w http.ResponseWriter, r *http.Request) {

	data := Emails{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for z := 0; z < len(data.Senders); z++ {
		for i := 0; i < len(data.Receivers); i++ {
			fmt.Println("sending email with sendgrid..")
			sgResponse, err := email.SendGridEmail(data.Senders[z], data.Receivers[i])
			if err != nil || sgResponse.StatusCode > 400 {
				fmt.Printf("send grid failed: %e, Sending email with Send Grid", err)
				fmt.Println("Sending email with mailgun...")
				response, id, err := email.SendSimpleMessage(data.Senders[z], data.Receivers[i])
				if err != nil {
					fmt.Printf("email was sent. %s:%s \n", id, response)
				}
			} else {
				fmt.Println(sgResponse.StatusCode)
				fmt.Println(sgResponse.Body)
				fmt.Println(sgResponse.Headers)
				fmt.Println("Succsess with SendGrid.")

			}
		}
	}
}
