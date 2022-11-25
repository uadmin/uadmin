package uadmin

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

// SendEmail sends email using system configured variables
func SendEmail(to, cc, bcc []string, subject, body string, attachments ...string) (err error) {
	if EmailFrom == "" || EmailUsername == "" || EmailPassword == "" || EmailSMTPServer == "" || EmailSMTPServerPort == 0 {
		errMsg := "Email not sent because email global variables are not set"
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

	// prepare body by splitting it into lines of length 73 followed by =
	body = strings.ReplaceAll(body, "\n", "<br/>")
	body = strings.ReplaceAll(body, "=", "=3D")
	body = strings.Join(splitString(body, 73), "=\r\n")

	// Construct the email
	MIME := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"utf-8\";\r\nContent-Transfer-Encoding: quoted-printable\r\n"
	delimeter := fmt.Sprintf("--=_uAdmin_%s_%s.%s", GenerateBase32(3), GenerateBase32(9), GenerateBase32(10))
	if len(attachments) != 0 {
		MIME = "MIME-version: 1.0;\r\n"
		MIME += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\nContent-Transfer-Encoding: quoted-printable\r\n", delimeter)
		body = "\r\n--" + delimeter + "\r\n" + fmt.Sprintf("Content-Type: text/html; boundary=\"%s\"\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n", delimeter) + body

		for i := range attachments {
			// get filename
			filename := filepath.Base(attachments[i])

			// read the file
			rawFile, err := os.ReadFile(attachments[i])
			if err != nil {
				Trail(WARNING, "Unable to attach file %s. %s", attachments[i], err)
				continue
			}

			body += "\r\n\r\n--" + delimeter + "\r\n"
			body += fmt.Sprintf("Content-Type: "+getMimeFromFileName(filename)+"; boundary=\"%s\"\r\n", delimeter)
			body += "Content-Transfer-Encoding: base64\r\n"
			body += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"
			body += "Content-Transfer-Encoding: quoted-printable\r\n"
			body += "\r\n" + strings.Join(splitString(base64.StdEncoding.EncodeToString(rawFile), 73), "\r\n")
		}
	}

	msg := "From: " + EmailFrom + "\r\n"
	msg += "To: " + strings.Join(to, ",") + "\r\n"
	if len(cc) > 0 {
		msg += "CC: " + strings.Join(cc, ",") + "\r\n"
	}
	msg += "Date: " + time.Now().UTC().Format(time.RFC1123Z) + "\r\n"
	msg += "Message-ID: " + fmt.Sprintf("<%s-%s-%s-%s-%s@%s>", GenerateBase32(8), GenerateBase32(4), GenerateBase32(4), GenerateBase32(4), GenerateBase32(12), domain[0]) + "\r\n"
	msg += "Subject: " + subject + "\r\n"
	msg += MIME + "\r\n"
	msg += body
	msg += "\r\n\r\n"
	// Append CC and BCC
	if cc != nil {
		to = append(to, cc...)
	}
	if bcc != nil {
		to = append(to, bcc...)
	}

	go func() {
		err = smtp.SendMail(fmt.Sprintf("%s:%d", EmailSMTPServer, EmailSMTPServerPort),
			smtp.PlainAuth(EmailFrom, EmailUsername, EmailPassword, EmailSMTPServer),
			EmailFrom, to, []byte(msg))

		if err != nil {
			Trail(WARNING, "Email was not sent. %s", err)
		}
	}()

	return nil
}

func splitString(v string, maxLen int) []string {
	splits := []string{}

	var l, r int
	for l, r = 0, maxLen; r < len(v); l, r = r, r+maxLen {
		for !utf8.RuneStart(v[r]) {
			r--
		}
		splits = append(splits, v[l:r])
	}
	splits = append(splits, v[l:])
	return splits
}

func getMimeFromFileName(v string) string {
	ext := filepath.Ext(v)
	mType := mime.TypeByExtension(ext)
	if mType == "" {
		return "application/octet-stream"
	}
	return mType
}
