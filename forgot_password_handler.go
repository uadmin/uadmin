package uadmin

import (
	"fmt"
	"net/http"
	"strings"
)

// forgotPasswordHandler !
func forgotPasswordHandler(u *User, r *http.Request, link string, msg string) error {
	if u.Email == "" {
		return fmt.Errorf("unable to reset password, the user does not have an email")
	}
	if msg == "" {
		msg = `<p>Dear {NAME},</p>

		Have you forgotten your password to access {WEBSITE}. Don't worry we got your back. Please follow the link below to reset your password.
		
		If you want to reset your password, click this link:
		<a href="{URL}">{URL}</a>
		
		If you didn't request a password reset, you can ignore this message.
		
		Regards,
		{WEBSITE} Support
		`
	}

	link, err := u.GeneratePasswordResetLink(r, link)
	if err != nil {
		return err
	}

	msg = strings.ReplaceAll(msg, "{NAME}", u.String())
	msg = strings.ReplaceAll(msg, "{WEBSITE}", SiteName)
	msg = strings.ReplaceAll(msg, "{URL}", link)
	subject := "Password reset for " + SiteName

	err = SendEmail([]string{u.Email}, []string{}, []string{}, subject, msg)

	return err
}
