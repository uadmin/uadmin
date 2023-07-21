package uadmin

import (
	"context"
	"net/http"
	"strings"
)

func dAPIOpenIDLoginHandler(w http.ResponseWriter, r *http.Request, s *Session) {
	_ = s
	redirectURI := r.FormValue("redirect_uri")

	Trail(DEBUG, "HERE")

	if r.Method == "GET" {
		if session := IsAuthenticated(r); session != nil {
			Preload(session, "User")
			c := map[string]interface{}{
				"SiteName":         SiteName,
				"Language":         getLanguage(r),
				"RootURL":          RootURL,
				"Logo":             Logo,
				"user":             s.User,
				"OpenIDWebsiteURL": redirectURI,
			}
			RenderHTML(w, r, "./templates/uadmin/"+Theme+"/openid_concent.html", c)
			return
		}

		http.Redirect(w, r, RootURL+"login/?next="+RootURL+"api/d/auth/openidlogin?"+r.URL.Query().Encode(), 303)
		return
	}

	if s == nil {
		w.WriteHeader(http.StatusUnauthorized)
		ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Invalid credentials",
		})
		return
	}

	// Preload the user to get the group name
	Preload(&s.User)

	ctx := context.WithValue(r.Context(), CKey("aud"), getAUD(redirectURI))
	r = r.WithContext(ctx)
	jwt := createJWT(r, s)

	http.Redirect(w, r, redirectURI+"?access-token="+jwt, 303)

}

func getAUD(URL string) string {
	aud := ""

	if strings.HasPrefix(URL, "https://") {
		aud = "https://"
		URL = strings.TrimPrefix(URL, "https://")
	}

	if strings.HasPrefix(URL, "http://") {
		aud = "http://"
		URL = strings.TrimPrefix(URL, "http://")
	}

	if strings.Contains(URL, "/") {
		URL = URL[:strings.Index(URL, "/")]
		aud += URL
	}

	return aud
}
