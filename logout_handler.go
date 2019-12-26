package uadmin

import (
	"net/http"
)

// logoutHandler !
func logoutHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	Logout(r)

	// Expire all cookies on logout
	for _, cookie := range r.Cookies() {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, RootURL, 303)
}
