package uadmin

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/uadmin/uadmin/colors"
)

const w2 = "" +
	`         ___       __          _` + "\n" +
	colors.FG_BLUE_B + `  __  __` + colors.FG_NORMAL + `/   | ____/ /___ ___  (_)___` + "\n" +
	colors.FG_BLUE_B + ` / / / /` + colors.FG_NORMAL + ` /| |/ __  / __ '__ \/ / __ \` + "\n" +
	colors.FG_BLUE_B + `/ /_/ /` + colors.FG_NORMAL + ` ___ / /_/ / / / / / / / / / /` + "\n" +
	colors.FG_BLUE_B + `\__,_/` + colors.FG_NORMAL + `_/  |_\__,_/_/ /_/ /_/_/_/ /_/` + "\n"

const welcomeMessage = `` +
	`        ______      __` + "\n" +
	`       /\  _  \    /\ \              __` + "\n" +
	colors.FG_BLUE_B + ` __  __` + colors.FG_NORMAL + `\ \ \L\ \   \_\ \    ___ ___ /\_\    ___` + "\n" +
	colors.FG_BLUE_B + `/\ \/\ \` + colors.FG_NORMAL + `\ \  __ \  /'_' \ /' __' __'\/\ \ /' _ '\` + "\n" +
	colors.FG_BLUE_B + `\ \ \_\ \` + colors.FG_NORMAL + `\ \ \/\ \/\ \L\ \/\ \/\ \/\ \ \ \/\ \/\ \` + "\n" +
	colors.FG_BLUE_B + ` \ \____/` + colors.FG_NORMAL + ` \ \_\ \_\ \___,_\ \_\ \_\ \_\ \_\ \_\ \_\` + "\n" +
	colors.FG_BLUE_B + `  \/___/ ` + colors.FG_NORMAL + `  \/_/\/_/\/__,_ /\/_/\/_/\/_/\/_/\/_/\/_/` + "\n" +
	``

// StartServer !
func StartServer() {
	if !registered {
		Register()
	}

	if val := getBindIP(); val != "" {
		BindIP = val
	}
	if BindIP == "" {
		BindIP = "0.0.0.0"
	}
	Trail(OK, "Server Started: http://%s:%d", BindIP, Port)
	fmt.Println(welcomeMessage)
	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", BindIP, Port), nil))

}

// StartSecureServer !
func StartSecureServer(certFile, keyFile string) {
	if !registered {
		Register()
	}
	if val := getBindIP(); val != "" {
		BindIP = val
	}
	if BindIP == "" {
		BindIP = "0.0.0.0"
	}
	Trail(OK, "Server Started: https://%s:%d\n", BindIP, Port)
	fmt.Println(welcomeMessage)
	log.Println(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", BindIP, Port), certFile, keyFile, nil))
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
