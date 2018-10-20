package uadmin

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// StartServer !
func StartServer(port int) {
	if !registered {
		Register()
	}

	if val := getBindIP(); val != "" {
		BindIP = val
	}
	if BindIP == "" {
		BindIP = "0.0.0.0"
	}
	Trail(OK, "Server Started: http://%s:%d", BindIP, port)
	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", BindIP, port), nil))

}

// StartSecureServer !
func StartSecureServer(port int, certFile, keyFile string) {
	if !registered {
		Register()
	}
	if val := getBindIP(); val != "" {
		BindIP = val
	}
	if BindIP == "" {
		BindIP = "0.0.0.0"
	}
	Trail(OK, "Server Started: https://%s:%d\n", BindIP, port)
	if val := getBindIP(); val != "" {
		BindIP = val
	}
	log.Println(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", BindIP, port), certFile, keyFile, nil))
}

func getBindIP() string {
	// Check if there is a bind ip file in the source code
	ex, _ := os.Executable()
	buf, err := ioutil.ReadFile(path.Join(filepath.Dir(ex), ".bindip"))
	if err != nil {
		return string(buf)
	}
	return ""
}
