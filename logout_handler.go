package uadmin

import (
	"net/http"
)

// logoutHandler !
func logoutHandler(w http.ResponseWriter, r *http.Request, session *Session) {
	cookie, err := r.Cookie("session")

	if err == nil && cookie != nil {
		// Store Logout to the user log
		go func() {
			log := &Log{}
			log.SignIn(session.User.Username, log.Action.Logout(), r)
			log.Save()
		}()

		session.Logout()
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	http.Redirect(w, r, RootURL, 303)
}
