package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Help is the command line help for the cli tool
const Help = `Usage: uadmin COMMAND [--src]
This tools helps you prepare a folder for a new project or update static files and templates

Commands:
  prepare         Generates folders and prepares static and templates
  version         Shows the version of uAdmin

Arguments:
  --src           If you want to copy static files and templates from src folder

Get full documentation online:
https://uadmin-docs.readthedocs.io/en/latest/
`

// // BaseServer being used for publishing apps to hosting server
// const BaseServer = "https://publish.uadmin.io/"

// func generateBase36(length int) string {
// 	rand.Seed(time.Now().UnixNano())
// 	a := "abcdefghijklmnopqrstuvwxyz0123456789"
// 	val := ""
// 	for i := 0; i < length; i++ {
// 		val += string(a[rand.Intn(len(a))])
// 	}

// 	return val
// }

// var command string
// var email string
// var domain string
// var port string

// func init() {
// 	flag.StringVar(&email, "email", "", "Your email address")
// 	flag.StringVar(&email, "e", "", "Your email address")
// 	flag.StringVar(&domain, "domain", "", "You can choose your domain name which will customize your URL")
// 	flag.StringVar(&domain, "d", "", "You can choose your domain name which will customize your URL")
// }

func main() {
	var rootCmd = &cobra.Command{
		Use:   "uadmin",
		Short: "uadmin commandline tool",
		Long:  `This tools helps you prepare a folder for a new project or update static files and templates.`,
		// Run: func(cmd *cobra.Command, args []string) {
		//   // Do Stuff Here
		// },
	}
	rootCmd.AddCommand(cmdPrepare)
	rootCmd.AddCommand(cmdVersion)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
