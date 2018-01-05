package email

import (
	"fmt"
	"testing"
)

func TestSendGridEmail(t *testing.T) {
	testSend := "robot@me.com"
	testEmail := "m_norris30@yahoo.com"
	testSubject := "hello"
	testMessage := "hello"
	testHTML := "<strong>Fair is Awesome!!!!</strong>"

	response, err := SendGridEmail(testSend, testEmail, testSubject, testMessage, testHTML)
	if err != nil {
		fmt.Sprintf("response: %s", response)
		t.Errorf("error: %e", err)
	}
}
