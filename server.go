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
	colors.FGBlueB + `\__,_/` + colors.FGNormal + `_/  |_\__,_/_/ /_/ /_/_/_/ /_/` + "\n"

const w2 = `` +
	`        ______      __` + "\n" +
	`       /\  _  \    /\ \              __` + "\n" +
	colors.FGBlueB + ` __  __` + colors.FGNormal + `\ \ \L\ \   \_\ \    ___ ___ /\_\    ___` + "\n" +
	colors.FGBlueB + `/\ \/\ \` + colors.FGNormal + `\ \  __ \  /'_' \ /' __' __'\/\ \ /' _ '\` + "\n" +
	colors.FGBlueB + `\ \ \_\ \` + colors.FGNormal + `\ \ \/\ \/\ \L\ \/\ \/\ \/\ \ \ \/\ \/\ \` + "\n" +
	colors.FGBlueB + ` \ \____/` + colors.FGNormal + ` \ \_\ \_\ \___,_\ \_\ \_\ \_\ \_\ \_\ \_\` + "\n" +
	colors.FGBlueB + `  \/___/ ` + colors.FGNormal + `  \/_/\/_/\/__,_ /\/_/\/_/\/_/\/_/\/_/\/_/` + "\n" +
	``

var ServerReady = false

// StartServer !
func StartServer() {
	if !registered {
		Register()
	}
	if !settingsSynched {
		syncSystemSettings()
	}
	if !handlersRegistered {
		registerHandlers()
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
	fmt.Println(welcomeMessage)
	dbOK = true
	ServerReady = true
	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", BindIP, Port), nil))
}

// StartSecureServer !
func StartSecureServer(certFile, keyFile string) {
	if !registered {
		Register()
	}
	if !settingsSynched {
		syncSystemSettings()
	}
	if !handlersRegistered {
		registerHandlers()
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

	Trail(OK, "Server Started: https://%s:%d\n", BindIP, Port)
	fmt.Println(welcomeMessage)
	dbOK = true
	ServerReady = true
	log.Println(http.ListenAndServeTLS(fmt.Sprintf("%s:%d", BindIP, Port), certFile, keyFile, nil))
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
