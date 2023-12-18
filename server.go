package uadmin

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/uadmin/uadmin/colors"
)

const welcomeMessage = "" +
	`         ___       __          _` + "\n" +
	colors.FGBlueB + `  __  __` + colors.FGNormal + `/   | ____/ /___ ___  (_)___` + "\n" +
	colors.FGBlueB + ` / / / /` + colors.FGNormal + ` /| |/ __  / __ '__ \/ / __ \` + "\n" +
	colors.FGBlueB + `/ /_/ /` + colors.FGNormal + ` ___ / /_/ / / / / / / / / / /` + "\n" +
	colors.FGBlueB + `\__,_/` + colors.FGNormal + `_/  |_\__,_/_/ /_/ /_/_/_/ /_/` + "\n\n"

// const w2 = `` +
// 	`        ______      __` + "\n" +
// 	`       /\  _  \    /\ \              __` + "\n" +
// 	colors.FGBlueB + ` __  __` + colors.FGNormal + `\ \ \L\ \   \_\ \    ___ ___ /\_\    ___` + "\n" +
// 	colors.FGBlueB + `/\ \/\ \` + colors.FGNormal + `\ \  __ \  /'_' \ /' __' __'\/\ \ /' _ '\` + "\n" +
// 	colors.FGBlueB + `\ \ \_\ \` + colors.FGNormal + `\ \ \/\ \/\ \L\ \/\ \/\ \/\ \ \ \/\ \/\ \` + "\n" +
// 	colors.FGBlueB + ` \ \____/` + colors.FGNormal + ` \ \_\ \_\ \___,_\ \_\ \_\ \_\ \_\ \_\ \_\` + "\n" +
// 	colors.FGBlueB + `  \/___/ ` + colors.FGNormal + `  \/_/\/_/\/__,_ /\/_/\/_/\/_/\/_/\/_/\/_/` + "\n"

// ServerReady is a variable that is set to true once the server is ready to use
var ServerReady = false

func StartServerWithMux() (string, string, *http.ServeMux) {
	mux := http.NewServeMux()

	// register static and add parameter
	if !strings.HasSuffix(RootURL, "/") {
		RootURL = RootURL + "/"
	}
	if !strings.HasPrefix(RootURL, "/") {
		RootURL = "/" + RootURL
	}

	if !registered {
		Register()
	}
	if !settingsSynched {
		syncSystemSettings()
	}
	if !handlersRegistered {
		RegisterHandlersWithMuxer(mux, RootURL)
	}
	if val := getBindIP(); val != "" {
		BindIP = val
	}
	if BindIP == "" {
		BindIP = "0.0.0.0"
	}
	// Synch model translation
	// Get Global Schema
	stat := map[string]int{}
	for _, v := range CustomTranslation {
		tempStat := syncCustomTranslation(v)
		for k, v := range tempStat {
			stat[k] += v
		}
	}
	for k := range Schema {
		tempStat := syncModelTranslation(Schema[k])
		for k, v := range tempStat {
			stat[k] += v
		}
	}
	for k, v := range stat {
		complete := float64(v) / float64(stat["en"])
		if complete != 1 {
			Trail(WARNING, "Translation of %s at %.0f%% [%d/%d]", k, complete*100, v, stat["en"])
		}
	}

	Trail(OK, "Server Started: http://%s:%d", BindIP, Port)
	Trail(NONE, welcomeMessage)
	dbOK = true
	ServerReady = true

	return RootURL, BindIP, mux
}

// StartServer !
func StartServer() {
	_, bindIp, mux := StartServerWithMux()
	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", bindIp, Port), mux))
}

// StartSecureServer !
func StartSecureServer(certFile, keyFile string) {
	_, bindIp, mux := StartServerWithMux()
	log.Println(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", bindIp, Port), certFile, keyFile, mux))
}

func getBindIP() string {
	// Check if there is a bind ip file in the source code
	ex, _ := os.Executable()
	buf, err := ioutil.ReadFile(path.Join(filepath.Dir(ex), ".bindip"))
	if err == nil {
		return strings.Replace(string(buf), "\n", "", -1)
	}
	return ""
}
