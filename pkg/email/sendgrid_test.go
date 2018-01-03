package email

import (
	"fmt"
	"testing"
)

func TestSendGridEmailError(t *testing.T) {
	testSend := "robot@me.com"
	testEmail := "m_norris30@yahoo.com"

	response, err := SendGridEmail(testSend, testEmail)
	if err != nil {
		fmt.Sprintf("response: %s", response)
		t.Errorf("error: %e", err)
	}
}


