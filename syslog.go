// +build !windows

package uadmin

import (
	"log/syslog"
)

var syslogMap = map[int]syslog.Priority{
	DEBUG:     syslog.LOG_DEBUG,
	WORKING:   syslog.LOG_DEBUG,
	INFO:      syslog.LOG_INFO,
	OK:        syslog.LOG_INFO,
	WARNING:   syslog.LOG_WARNING,
	ERROR:     syslog.LOG_ERR,
	CRITICAL:  syslog.LOG_CRIT,
	ALERT:     syslog.LOG_ALERT,
	EMERGENCY: syslog.LOG_EMERG,
}

func Syslogf(level int, msg string, a ...interface{}) {
	logger, err := syslog.NewLogger(syslog.LOG_LOCAL0|syslogMap[level], 0)
	if err != nil {
		Trail(ERROR, "Hanlder.NewLogger. %s", err.Error())
		return
	}
	if len(msg) != 0 && msg[len(msg)-1] != '\n' {
		msg += "\n"
	}
	logger.Printf(msg, a...)
}
