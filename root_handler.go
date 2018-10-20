package uadmin

import (
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It Works!, Kinda?"))
}
