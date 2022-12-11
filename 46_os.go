package main 

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])

	// Get Hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	/* for detail abou os package visit: golang.org/pkg */

	// Get env 
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	/* go run file.go kurnia sandi */
}