package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin/helper"
)

// Help is the command line help for the cli tool
const Help = `Usage: uadmin COMMAND [-e email] [-d domain]
This tools allows you to publish your project online

Commands:
  publish         This publishes your project online
  prepare         Generates folders and prepares static and templates
  version         Shows the version of uAdmin

Arguments:
  -e, --email     Your email. This is required for you to be able to maintain your project.
  -d, --domain    You can choose your domain name which will customize your URL

Get full documentation online:
https://uadmin-docs.readthedocs.io/en/latest/
`

// BaseServer being used for publishing apps to hosting server
const BaseServer = "https://publish.uadmin.io/"

func generateBase36(length int) string {
	rand.Seed(time.Now().UnixNano())
	a := "abcdefghijklmnopqrstuvwxyz0123456789"
	val := ""
	for i := 0; i < length; i++ {
		val += string(a[rand.Intn(len(a))])
	}

	return val
}

var command string
var email string
var domain string
var port string

func init() {
	flag.StringVar(&email, "email", "", "Your email address")
	flag.StringVar(&email, "e", "", "Your email address")
	flag.StringVar(&domain, "domain", "", "You can choose your domain name which will customize your URL")
	flag.StringVar(&domain, "d", "", "You can choose your domain name which will customize your URL")
}

func main() {
	args := os.Args

	// Check if there are no args
	if len(args) < 2 {
		fmt.Println(Help)
		return
	}

	// Check if the first arg is not a command
	if strings.HasSuffix(args[1], "-") {
		fmt.Println("ERROR: Invalid sytax. Please provide a command")
		fmt.Println(Help)
		return
	}
	command := args[1]

	flag.Parse()

	// prepapre command
	if command == "prepare" {
		var dst string
		var src string
		// First ge the path
		ex, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// Generate folders
		folderList := []string{"models", "api", "views", "media"}
		for _, v := range folderList {
			dst = filepath.Join(ex, v)
			if _, err = os.Stat(dst); os.IsNotExist(err) {
				err = os.MkdirAll(dst, os.FileMode(0744))
				if err != nil {
					uadmin.Trail(uadmin.WARNING, "Unable to create \"%s\" folder: %s", v, err)
				} else {
					uadmin.Trail(uadmin.OK, "Created: %s", dst)
				}
			}
		}

		// Copy folders
		folderList = []string{"static", "templates"}
		goPath := os.Getenv("GOPATH")
		if goPath == "" {
			if runtime.GOOS == "windows" {
				goPath = filepath.Join(os.Getenv("USERPROFILE"), "go")
			} else {
				goPath = filepath.Join(os.Getenv("HOME"), "go")
			}
			uadmin.Trail(uadmin.INFO, "Your GOPATH environment variable is not set. Using the default path: %s", goPath)
		}
		uadminPath := filepath.Join(goPath, "src", "github.com", "uadmin", "uadmin")
		for _, v := range folderList {
			msg := "Updated"
			if _, err = os.Stat(filepath.Join(ex, v)); os.IsNotExist(err) {
				msg = "Created"
			}
			dst = filepath.Join(ex, v)
			src = filepath.Join(uadminPath, v)
			err := Copy(src, dst)
			if err != nil {
				uadmin.Trail(uadmin.WARNING, "Unable to copy \"%s\" folder: %s", v, err)
			} else {
				uadmin.Trail(uadmin.OK, msg+": %s", dst)
			}
		}
		return
	} else if command == "version" {
		uadmin.Trail(uadmin.INFO, uadmin.Version)
		return
	} else if command == "publish" {
		// Get user info
		var buf []byte
		uadminProfile := map[string]string{}
		u, _ := user.Current()
		uadminPath := filepath.Join(u.HomeDir, ".uadmin")
		uadminFile, err := os.Open(uadminPath)
		if err != nil {
			uadminFile, err = os.Create(uadminPath)
			if err != nil {
				uadmin.Trail(uadmin.ERROR, "unable to create your uadmin profile: %s", err)
				return
			}
			buf = []byte("{}")
			uadminFile.Write(buf)
		} else {
			buf, err = ioutil.ReadAll(uadminFile)
			if err != nil {
				uadmin.Trail(uadmin.ERROR, "unable to access your uadmin profile: %s", err)
				return
			}
		}
		stat, _ := uadminFile.Stat()
		uadminFileMode := stat.Mode()
		uadminFile.Close()

		err = json.Unmarshal(buf, &uadminProfile)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Unable to read your uadmin profile")
		}

		// Collect Project info
		projProfile := map[string]string{}
		projPath := "./.uproj"
		projFile, err := os.Open(projPath)
		if err != nil {
			projFile, err = os.Create(projPath)
			if err != nil {
				uadmin.Trail(uadmin.ERROR, "unable to create your project profile: %s", err)
				return
			}
			buf = []byte("{}")
			projFile.Write(buf)
		} else {
			buf, err = ioutil.ReadAll(projFile)
			if err != nil {
				uadmin.Trail(uadmin.ERROR, "unable to access your project profile: %s", err)
				return
			}
		}
		stat, _ = projFile.Stat()
		projFileMode := stat.Mode()
		projFile.Close()

		err = json.Unmarshal(buf, &projProfile)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Unable to read your project profile")
		}

		// Collect Information that's not already in uadmin profile
		if email == "" {
			if _, ok := uadminProfile["email"]; !ok {
				for {
					fmt.Print("Enter your email: ")
					fmt.Scanf("%s", &email)
					if helper.ValidateEmail(email) {
						break
					}
					fmt.Println("Invalid email address.")
				}
				uadminProfile["email"] = email
				buf, _ = json.Marshal(uadminProfile)
				ioutil.WriteFile(uadminPath, buf, uadminFileMode)
			} else {
				email = uadminProfile["email"]
			}
		}
		if domain == "" {
			if _, ok := projProfile["domain"]; !ok {
				fmt.Println("Your project will be published to https://my-proj.uadmin.io")
				for {
					fmt.Print("Enter the name of your sub-domain (my-proj) [auto]: ")
					fmt.Scanf("%s", &domain)
					if helper.ValidateSubdomain(domain) || domain == "" {
						break
					}
					fmt.Println("Invalid sub domain")
				}
				projProfile["domain"] = domain
				buf, _ = json.Marshal(projProfile)
				ioutil.WriteFile(projPath, buf, projFileMode)
			} else {
				domain = projProfile["domain"]
			}
		}
		if port == "" {
			if _, ok := projProfile["port"]; !ok {
				fmt.Println("Did you change the default port from 8080?")
				fmt.Println("This is the port you have in uadmin.Port = 8080")
				var tempPort uint64
				for {
					fmt.Print("Enter the port that your server run on [8080]: ")
					fmt.Scanf("%s", &port)

					if tempPort, err = strconv.ParseUint(port, 10, 32); (err == nil) && (tempPort <= 65535) && (tempPort >= 1024) {
						break
					}
					if port == "" {
						port = "8080"
						break
					}
					fmt.Println("Invalid port. You should use a port between 1024-65535")
				}
				projProfile["port"] = port
				buf, _ = json.Marshal(projProfile)
				ioutil.WriteFile(projPath, buf, projFileMode)
			} else {
				port = projProfile["port"]
			}
		}
		token, _ := uadminProfile["token"]
		uid, _ := projProfile["uid"]
		ex, err := os.Getwd()
		if err != nil {
			return
		}
		GOPATH := os.Getenv("GOPATH")
		if GOPATH == "" {
			uadmin.Trail(uadmin.ERROR, "Your GOPATH environment variable is not set")
			return
		}
		myPath := strings.TrimPrefix(ex, GOPATH)
		myPath = strings.TrimPrefix(myPath, string(os.PathSeparator))
		myPath = strings.TrimSuffix(myPath, string(os.PathSeparator))

		nameParts := strings.Split(myPath, string(os.PathSeparator))

		// Compress files
		Compress(ex)
		defer os.Remove("./publish.tmp")

		// Send publish request
		extraParams := map[string]string{
			"email":  email,
			"token":  token,
			"uid":    uid,
			"path":   myPath,
			"domain": domain,
			"port":   port,
			"name":   nameParts[len(nameParts)-1],
		}

		client := &http.Client{}
		client.Timeout = time.Minute * 10
		var resp *http.Response
		var request *http.Request
		uadmin.Trail(uadmin.WORKING, "Uploading your application")
		for i := 0; i < 5; i++ {
			request, err = newfileUploadRequest(BaseServer+"api/push/", extraParams, "file", ex+"/publish.tmp")
			if err != nil {
				uadmin.Trail(uadmin.ERROR, "Unable to prepare you publish request. %s", err)
				return
			}

			resp, err = client.Do(request)
			if err == nil {
				break
			}
			uadmin.Trail(uadmin.ERROR, "Unable to connect to uadmin server. %s", err)
			return
		}
		uadmin.Trail(uadmin.OK, "Your application has been uploaded")

		uadmin.Trail(uadmin.WORKING, "Installing your application")
		// Parse The response
		buf, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			uadmin.Trail(uadmin.ERROR, "Unable to read response back from server.")
			return
		}
		uadmin.Trail(uadmin.OK, "Application installed successfully")

		// Store back any values if they exist
		respObj := map[string]string{}
		json.Unmarshal(buf, &respObj)
		if respToken, ok := respObj["token"]; ok {
			if respToken != uadminProfile["token"] {
				uadminProfile["token"] = respToken
				buf, _ := json.Marshal(uadminProfile)
				ioutil.WriteFile(uadminPath, buf, uadminFileMode)
			}
		}
		if respDomain, ok := respObj["domain"]; ok {
			// Get the subdomain from the URL
			respDomain = strings.Split(strings.Split(respDomain, "://")[1], ".")[0]
			if respDomain != projProfile["domain"] {
				projProfile["domain"] = respDomain
				buf, _ := json.Marshal(projProfile)
				ioutil.WriteFile(projPath, buf, projFileMode)
			}
		}
		if respDomain, ok := respObj["port"]; ok {
			if respDomain != projProfile["port"] {
				projProfile["port"] = respDomain
				buf, _ := json.Marshal(projProfile)
				ioutil.WriteFile(projPath, buf, projFileMode)
			}
		}
		if respUID, ok := respObj["uid"]; ok {
			if respUID != projProfile["uid"] {
				projProfile["uid"] = respUID
				buf, _ := json.Marshal(projProfile)
				ioutil.WriteFile(projPath, buf, projFileMode)
			}
		}

		if respObj["status"] == "ok" {
			uadmin.Trail(uadmin.OK, "Your Project has been published to %s", respObj["domain"])
		} else {
			uadmin.Trail(uadmin.ERROR, "Unable to publish your project %s", respObj["err_msg"])
		}
	}
}
