package uadmin

import (
	"net/http"
	"strings"
)

func CORSHandler(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		if len(AllowedCORSOrigins) == 0 {
			if r.Header.Get("Origin") != "" {
				w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Add("Access-Control-Allow-Credentials", "true")
			} else {
				w.Header().Add("Access-Control-Allow-Origin", "*")
			}
		} else {
			allowedOrigin := false
			for i := range AllowedCORSOrigins {
				if strings.EqualFold(r.Header.Get("Origin"), AllowedCORSOrigins[i]) {
					allowedOrigin = true
					break
				}
			}
			if allowedOrigin {
				w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Add("Access-Control-Allow-Credentials", "true")
			} else {
				w.Header().Add("Access-Control-Allow-Origin", strings.Join(AllowedCORSOrigins, "|"))
				w.Header().Add("Access-Control-Allow-Credentials", "true")
			}

		}

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			// Allow all known methods
			w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD")

			// Allow requested headers
			w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))

			return
		}

		// run the handler
		f(w, r)
	}
}
