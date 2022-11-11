package uadmin

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"

	"github.com/uadmin/uadmin/colors"
)

// Reporting Levels
const (
	NONE      = 10
	DEBUG     = 0
	WORKING   = 1
	INFO      = 2
	OK        = 3
	WARNING   = 4
	ERROR     = 5
	CRITICAL  = 6
	ALERT     = 7
	EMERGENCY = 8
)

var trailTag = map[int]string{
	NONE:      colors.None,
	DEBUG:     colors.Debug,
	WORKING:   colors.Working,
	INFO:      colors.Info,
	OK:        colors.OK,
	WARNING:   colors.Warning,
	ERROR:     colors.Error,
	CRITICAL:  colors.Critical,
	ALERT:     colors.Alert,
	EMERGENCY: colors.Emergency,
}

var levelMap = map[int]string{
	NONE:      "",
	DEBUG:     "[  DEBUG ]   ",
	WORKING:   "[ WORKING]   ",
	INFO:      "[  INFO  ]   ",
	OK:        "[   OK   ]   ",
	WARNING:   "[ WARNING]   ",
	ERROR:     "[  ERROR ]   ",
	CRITICAL:  "[CRITICAL]   ",
	ALERT:     "[  ALERT ]   ",
	EMERGENCY: "[  EMERG ]   ",
}

var trailBytes = []byte{}
var trailChan = []chan string{} //make(chan string)

// Trail prints to the log
func Trail(level int, msg interface{}, i ...interface{}) {
	if level >= ReportingLevel {
		message := fmt.Sprint(msg)
		if level != WORKING && !strings.HasSuffix(message, "\n") {
			message += "\n"
		} else if level == WORKING && !strings.HasPrefix(message, "\r") {
			message = message + "\r"
		}
		if ReportTimeStamp {
			log.Printf(trailTag[level]+message, i...)
		} else {
			fmt.Printf(trailTag[level]+message, i...)
		}

		// Run error handler if it exists
		if ErrorHandleFunc != nil {
			stack := string(debug.Stack())
			stackList := strings.Split(stack, "\n")
			stack = strings.Join(stackList[5:], "\n")
			go ErrorHandleFunc(level, fmt.Sprintf(fmt.Sprint(msg), i...), stack)
		}

		// Log to syslog
		if LogTrail && level >= TrailLoggingLevel && level != WORKING {
			// Send log to syslog
			Syslogf(level, message, i...)
		}

		// Add tail bytes
		if level != WORKING {
			trailBytes = append(trailBytes, []byte(fmt.Sprintf(trailTag[level]+message, i...))...)
			go func(message string) {
				for i := len(trailChan) - 1; i >= 0; i-- {
					select {
					case trailChan[i] <- message:
						// fmt.Println("sent " + message)
					default:
						trailChan = append(trailChan[:i], trailChan[i+1:]...)
						fmt.Println("deleted")
					}
				}
			}(fmt.Sprintf(trailTag[level]+message, i...))

			if len(trailBytes) > TrailCacheSize {
				trailBytes = trailBytes[len(trailBytes)-TrailCacheSize:]
			}
		}
	}
}

func RegisterTrailChan(c chan string) {
	trailChan = append(trailChan, c)
}
