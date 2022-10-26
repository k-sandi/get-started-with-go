/*
NAMED RETURN VALUES

*/

package main 

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Kurnia"
	lastName = "Sandi"

	return
}

func main() {
	a, b := getCompleteName()
	fmt.Println(a)
	fmt.Println(b)
}