package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpKurnia NoKTP = "123456789"
	var marriedStatus Married = true
	fmt.Println(noKtpKurnia)
	fmt.Println(marriedStatus)
}