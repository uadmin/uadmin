package uadmin

import (
	"net"
	"net/http"
	"regexp"
	"strings"
)

// SQLInjection is the function to check for SQL injection attacks.
// Parameters:
//
//	-key: column_name, table name
//	-value: WHERE key(OP)value, SET key=value, VALUES (key,key...)
//
// return true for sql injection attempt and false for safe requests
func SQLInjection(r *http.Request, key, value string) bool {
	var err error

	user := GetUserFromRequest(r)
	if user == nil {
		user = &User{}
	}
	ip := GetRemoteIP(r)
	if ip, _, err = net.SplitHostPort(ip); err != nil {
		ip = GetRemoteIP(r)
	}
	errMsg := "SQL Injection attempt (%s '%s'). User:" + user.Username + " IP:" + ip
	if key != "" {
		// Case 1 - Comment injection
		if strings.Contains(key, "--") || strings.Contains(key, "#") {
			Trail(CRITICAL, errMsg, "comment injection", key)
			return true
		}
		// Case 2 - Comment injection
		if strings.Contains(key, "/*") || strings.Contains(key, "*/") {
			Trail(CRITICAL, errMsg, "comment injection", key)
			return true
		}
		// Case 3 - Stacking
		if strings.Contains(key, ";") {
			Trail(CRITICAL, errMsg, "stacking", key)
			return true
		}
		// Case 4 - HEX Injection
		if strings.Contains(key, "0x") {
			Trail(CRITICAL, errMsg, "hex injection", key)
			return true
		}
		// Case 5 - Concatenation
		if strings.Contains(key, "+") || strings.Contains(key, "||") {
			Trail(CRITICAL, errMsg, "concatenation", key)
			return true
		}
		// Case 6 - Functions
		if strings.Contains(key, "(") || strings.Contains(key, ")") {
			// Skip allows functions
			if match, _ := regexp.MatchString(`DATE\([a-z_]*\)`, key); match {
			} else if match, _ := regexp.MatchString(`YEAR\([a-z_]*\)`, key); match {
			} else if match, _ := regexp.MatchString(`MONTH\([a-z_]*\)`, key); match {
			} else if match, _ := regexp.MatchString(`DAY\([a-z_]*\)`, key); match {
			} else {
				Trail(CRITICAL, errMsg, "functions", key)
				return true
			}

		}
		// Case 7 - Space
		if strings.Contains(key, " ") {
			Trail(CRITICAL, errMsg, "space", key)
			return true
		}
		// Case 8 - Escaping
		if strings.Contains(key, "'") || strings.Contains(key, "`") || strings.Contains(key, "\"") {
			Trail(CRITICAL, errMsg, "escaping", key)
			return true
		}
		// Case 9 - Escaping
		if strings.Contains(value, "'") || strings.Contains(value, "`") {
			Trail(CRITICAL, errMsg, "escaping", value)
			return true
		}
	}
	// if value != ""
	// We are depending on gorm in here
	return false
}
