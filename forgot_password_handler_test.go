package uadmin

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestForgotPasswordHandler is a unit testing function for forgotPasswordHandler() function
func TestForgotPasswordHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	user := User{}
	Get(&user, "id = ?", 1)

	err := forgotPasswordHandler(&user, r)
	if err == nil {
		t.Errorf("forgotPasswordHandler didn't return an error on a user with no email")
	}

	user.Email = "user@example.com"
	err = forgotPasswordHandler(&user, r)
	if err != nil {
		t.Errorf("forgotPasswordHandler returned an error. %s", err)
	}
	time.Sleep(time.Millisecond * 500)
	if receivedEmail == "" {
		t.Errorf("forgotPasswordHandler didn't send an email")
	}
	if !strings.Contains(receivedEmail, "From: uadmin@example.com") {
		t.Errorf("SendEmail don't have a valid From")
	}
	if !strings.Contains(receivedEmail, "To: user@example.com") {
		t.Errorf("SendEmail don't have a valid To")
	}
}
