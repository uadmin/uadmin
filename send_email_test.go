package uadmin

import (
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

var receivedEmail string

// TestSendEmail is a unit testing function for SendEmail() function
func TestSendEmail(t *testing.T) {
	// This email should be sent
	SendEmail([]string{"user@example.com"}, []string{}, []string{}, "subject", "body")
	time.Sleep(time.Millisecond * 500)
	if receivedEmail == "" {
		t.Errorf("SendEmail didn't send an email")
	}
	if !strings.Contains(receivedEmail, "From: uadmin@example.com") {
		t.Errorf("SendEmail don't have a valid From")
	}
	if !strings.Contains(receivedEmail, "To: user@example.com") {
		t.Errorf("SendEmail don't have a valid To")
	}
	if !strings.Contains(receivedEmail, "Subject: subject") {
		t.Errorf("SendEmail don't have a valid Subject")
	}
	if !strings.Contains(receivedEmail, "body") {
		t.Errorf("SendEmail don't have a valid body")
	}
	receivedEmail = ""

	// Not try sending an email with missing settings
	temp := EmailUsername
	EmailUsername = ""
	err := SendEmail([]string{"user@example.com"}, []string{}, []string{}, "subject", "body")
	if err == nil {
		t.Errorf("SendEmail send an email with missing settings")
	}
	EmailUsername = temp
}

func startEmailServer() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost:2525")
	if err != nil {
		Trail(ERROR, "listening: %s", err)
		return
	}

	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			Trail(ERROR, "startEmailServer error accepting connection. %s", err)
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.

	conn.Write([]byte("220 smtp.example.com ESMTP Postfix (Ubuntu)\n"))

	_, err := conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}

	buf = make([]byte, 1024)
	conn.Write([]byte(`250-smtp.example.com
250-AUTH LOGIN PLAIN
250-PIPELINING
250-SIZE 102400000
250-VRFY
250-ETRN
250-ENHANCEDSTATUSCODES
250-8BITMIME
250 DSN
`))
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("235 Authentication succeeded\n"))

	buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("250 2.1.0 Ok\n"))

	buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("250 2.1.5 Ok\n"))

	buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("354 End data with <CR><LF>.<CR><LF>\n"))

	buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("250 2.0.0 Ok: queued as 16756A11026D\n"))
	receivedEmail = string(buf)

	buf = make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		Trail(ERROR, "reading: %s", err)
	}
	conn.Write([]byte("221 2.0.0 Bye\n"))

	// Close the connection when you're done with it.
	conn.Close()
}
