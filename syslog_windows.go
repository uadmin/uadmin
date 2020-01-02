// +build windows

package uadmin

import (
	"fmt"
	"os"
)

func Syslogf(level int, msg string, a ...interface{}) error {
	f, err := os.OpenFile("syslog.log", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	f.WriteString(levelMap[level] + fmt.Sprintf(msg, a...))
	f.Close()
	return nil
}
