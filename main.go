package uadmin

import (
	"flag"
	"fmt"
)

func main() {
	var email = flag.String("email", "", "Your email")
	flag.Parse()
	fmt.Printf("Welcome to uAdmin CLI\n")
	if *email == "" {
		fmt.Printf("Email: ")
		fmt.Scanf("%s", &email)
	}
}
