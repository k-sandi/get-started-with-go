/*
RETURNING MULTIPLE VALUES

o) Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return valuenya fi function.
*/

package main

import "fmt"

func getFullName() (string, string) {
	firstName := "Kurnia"
	lastName := "Sandi"
	return firstName, lastName
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(lastName)

	//ignore another variable or return value
	_, lastName := getFullName()
	// fmt.Println(firstName)
	fmt.Println(lastName)
}