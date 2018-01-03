package email

import (
	"fmt"
	"testing"
)

func TestSendSimpleMessage(t *testing.T) {

	testSend := "robot@me.com"
	testEmail := "m_norris30@yahoo.com"
	response, id, err := SendSimpleMessage(testSend, testEmail)
	if err != nil {
		fmt.Sprintf("response: %s", response)
		fmt.Sprintf("id: %s", id)
		t.Errorf("error: %e", err)
	}
}
