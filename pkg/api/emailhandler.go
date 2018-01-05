package api

import (
	"encoding/json"
	"net/http"

	"github.com/mike001005/emailApi/pkg/email"

	"fmt"
)

// Emails Args
type Emails struct {
	Emails []struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
		HTML    string `json:"html"`
	} `json:"emails"`
}

// Send Api
func Send(w http.ResponseWriter, r *http.Request) {

	statusCode, err := handleSendMailGun(w, r)
	if statusCode >= 400 {
		statusCode, err = handleSendGrid(w, r)
		if statusCode >= 400 {
			http.Error(w, err.Error(), statusCode)
		}
	}

}

// handleSendMailGun MailGun Api
func handleSendMailGun(w http.ResponseWriter, r *http.Request) (int, error) {
	data := Emails{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		if err != nil {
			fmt.Errorf("%s", err)
		}
	}

	for i := 0; i < len(data.Emails); i++ {
		response, id, err := email.SendSimpleMessage(data.Emails[i].From, data.Emails[i].To, data.Emails[i].Subject, data.Emails[i].Body)
		if err != nil {
			http.Error(w, err.Error(), 550)
		}
		w.Write([]byte(response + id))
	}

	return 200, nil
}

// handleSendSendGrid SendGrid Api
func handleSendGrid(w http.ResponseWriter, r *http.Request) (int, error) {

	data := Emails{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		if err != nil {
			fmt.Errorf("%s", err)
		}
	}
	for i := 0; i < len(data.Emails); i++ {
		response, err := email.SendGridEmail(data.Emails[i].From, data.Emails[i].To, data.Emails[i].Subject, data.Emails[i].Body, data.Emails[i].HTML)
		if err != nil {
			http.Error(w, err.Error(), 550)
		}
		w.Write([]byte(fmt.Sprintf("%s", response)))

	}
	return 200, nil
}
