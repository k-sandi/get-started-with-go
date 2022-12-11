package main 

import (
	"fmt"
	"flag"
)

func main() {
	/*
	* use pointer.
	* identify the flag data type.
	* store the argument (1: main argumnen, 2: default argument, 3: description or helper).
	*/
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	// parsing the flag (very important. You must run the command first before you print the his values)
	flag.Parse()

	// use pointer for get a value
	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *user)
	fmt.Println("Password : ", *password)
	fmt.Println("Number : ", *number)

	// the purpose of flag that can help you to use the helper for data validation in your terminal base apps
}