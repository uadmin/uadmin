package uadmin

import (
	"net"
	"net/http"
)

// CheckCSRF checks if the request is a possible CSRF
func CheckCSRF(r *http.Request) bool {
	var err error
	Trail(DEBUG, "Checking CSRF token: %s, s: %s", r.FormValue("x-csrf-token"), getSession(r))
	if r.FormValue("x-csrf-token") != "" && r.FormValue("x-csrf-token") == getSession(r) {
		return false
	}
	user := GetUserFromRequest(r)
	if user == nil {
		user = &User{}
	}
	ip := r.RemoteAddr
	if ip, _, err = net.SplitHostPort(ip); err != nil {
		ip = r.RemoteAddr
	}

	Trail(CRITICAL, "Request failed Anti-CSRF protection from user:%s IP:%s", user.Username, ip)
	return true
}
