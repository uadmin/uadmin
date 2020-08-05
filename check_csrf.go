package uadmin

import (
	"net"
	"net/http"
)

/*
CheckCSRF checks if the request is a possible CSRF. CSRF or Cross-Site
Request Forgery is a type of attack here a logged in user clicks on a link
that is sent to a website where the user is already authenticated and has
instructions for the website to change some state. A possible attack could
delete user or system data, change it or add new data to the system. Anti-CSRF
measures are implemented in all state chaning APIs and UI handler.

The way uAdmin implements CSRF is by checking for a request parameter GET or
POST called `x-csrf-token`. The value of this parameter could be equal to the
session key. You can get the session key from the session cookie or if you are
using `uadmin.RenderHTML` or `uadmin.RenderHTMLMulti`, then you will find it in
the context as `{{CSRF}}`. If you submitting a form you can add this value to
a hidden input.

To implement anti CSRF protection in your own API:

	func MyAPI(w http.ResponseWriter, r *http.Request) {
		if CheckCSRF(r) {
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status": "error",
				"err_msg": "The request does not have x-csrf-token",
			})
		}

		// API code ...
	}

	http.HandleFunc("/myapi/", MyAPI)

If you you call this API:

	http://0.0.0.0:8080/myapi/

It will return an error message and the system will create a CRITICAL
level log with details about the possible attack. To make the request
work, `x-csrf-token` paramtere should be added.

	http://0.0.0.0:8080/myapi/?x-csrf-token=MY_SESSION_KEY

Where you replace `MY_SESSION_KEY` with the session key.
*/
func CheckCSRF(r *http.Request) bool {
	var err error
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
