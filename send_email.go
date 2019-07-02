package uadmin

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

// SendEmail sends email using system configured variables
func SendEmail(to, cc, bcc []string, subject, body string) (err error) {
	if EmailFrom == "" || EmailUsername == "" || EmailPassword == "" || EmailSMTPServer == "" || EmailSMTPServerPort == 0 {
		errMsg := "Email not sent because email global variables are not set."
		Trail(WARNING, errMsg)
		return fmt.Errorf(errMsg)
	}

	// Get the domain name of sender
	domain := strings.Split(EmailFrom, "@")
	if len(domain) < 2 {
		return
	}
	domain[0] = strings.TrimSpace(domain[0])
	domain[0] = strings.TrimSuffix(domain[0], ">")

	// Construct the email
	MIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := "From: " + EmailFrom + "\r\n"
	msg += "To: " + strings.Join(to, ",") + "\r\n"
	if len(cc) > 0 {
		msg += "CC: " + strings.Join(cc, ",") + "\r\n"
	}
	msg += "Date: " + time.Now().UTC().Format(time.RFC1123Z) + "\r\n"
	msg += "Message-ID: " + fmt.Sprintf("<%s-%s-%s-%s-%s@%s>", GenerateBase32(8), GenerateBase32(4), GenerateBase32(4), GenerateBase32(4), GenerateBase32(12), domain[0]) + "\r\n"
	msg += "Subject: " + subject + "\r\n"
	msg += MIME + "\r\n"
	msg += strings.Replace(body, "\n", "<br/>", -1)
	msg += "\r\n"
	// Append CC and BCC
	to = append(to, cc...)
	to = append(to, bcc...)

	go func() {
		err = smtp.SendMail(fmt.Sprintf("%s:%d", EmailSMTPServer, EmailSMTPServerPort),
			smtp.PlainAuth("", EmailUsername, EmailPassword, EmailSMTPServer),
			EmailFrom, to, []byte(msg))

		if err != nil {
			Trail(WARNING, "Email was not sent. %s", err)
		}
	}()

	return nil
}
