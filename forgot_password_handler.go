package uadmin

import (
	"fmt"
	"net"
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

	// Check if the host name is in the allowed hosts list
	allowed := false
	var host string
	var allowedHost string
	var err error
	if host, _, err = net.SplitHostPort(GetHostName(r)); err != nil {
		host = r.Host
	}
	for _, v := range strings.Split(AllowedHosts, ",") {
		if allowedHost, _, err = net.SplitHostPort(v); err != nil {
			allowedHost = v
		}
		if allowedHost == host {
			allowed = true
			break
		}
	}
	host = GetHostName(r)
	if !allowed {
		Trail(CRITICAL, "Reset password request for host: (%s) which is not in AllowedHosts settings", host)
		return nil
	}

	schema := GetSchema(r)
	if link == "" {
		link = "{SCHEMA}://{HOST}" + RootURL + "resetpassword?u={USER_ID}&key={OTP}"
	}
	link = strings.ReplaceAll(link, "{SCHEMA}", schema)
	link = strings.ReplaceAll(link, "{HOST}", host)
	link = strings.ReplaceAll(link, "{USER_ID}", fmt.Sprint(u.ID))
	link = strings.ReplaceAll(link, "{OTP}", u.GetOTP())

	msg = strings.ReplaceAll(msg, "{NAME}", u.String())
	msg = strings.ReplaceAll(msg, "{WEBSITE}", SiteName)
	msg = strings.ReplaceAll(msg, "{URL}", link)
	subject := "Password reset for " + SiteName

	err = SendEmail([]string{u.Email}, []string{}, []string{}, subject, msg)

	return err
}
