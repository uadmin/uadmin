package uadmin

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func containsDotDot(v string) bool {
	if !strings.Contains(v, "..") {
		return false
	}
	for _, ent := range strings.FieldsFunc(v, func(r rune) bool { return r == '/' || r == '\\' }) {
		if ent == ".." {
			return true
		}
	}
	return false
}

// StaticHandler is a function that serves static files
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if containsDotDot(r.URL.Path) {
		w.WriteHeader(404)
		return
	}
	var modTime time.Time
	ab := false
	var midnightDelta int
	for k := range staticABTests {
		if k == r.URL.Path && len(staticABTests[k]) != 0 {
			index := getABT(r) % len(staticABTests[k])
			r.URL.Path = staticABTests[k][index].v

			// Calculate number of seconds until midnight if no calculated yet
			if midnightDelta == 0 {
				midnight := time.Now()
				midnight = time.Date(midnight.Year(), midnight.Month(), midnight.Day(), 0, 0, 0, 0, midnight.Location())
				midnightDelta = int(midnight.Sub(time.Now()).Seconds())
			}
			// Add a header to expire the satic content at midnigh
			w.Header().Add("Cache-Control", "private, max-age="+fmt.Sprint(midnightDelta))
			modTime = time.Now()
			ab = true

			go func() {
				abTestsMutex.Lock()
				t := staticABTests[k]
				t[index].imp++
				staticABTests[k] = t
				abTestsMutex.Unlock()
			}()
			break
		}
	}

	f, err := os.Open("." + r.URL.Path)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	if !ab {
		stat, err := os.Stat("." + r.URL.Path)
		if err != nil || stat.IsDir() {
			w.WriteHeader(404)
			return
		}
		modTime = stat.ModTime()
		w.Header().Add("Cache-Control", "private, max-age=3600")
	}

	http.ServeContent(w, r, "."+r.URL.Path, modTime, f)
}
