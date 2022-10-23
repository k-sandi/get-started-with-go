/*
if is PERCABANGAN
if-then-else
*/

package main 

import "fmt"

func main() {

	var name = "Jokosusanto"

	if name == "Sandi" {
		fmt.Println("Hello Sandi")
	} else if name == "Joko" {
		fmt.Println("Hi Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}


	/*
	IF with short statement
	1. "IF" support the short statement before condition statement.
	2. Make simple statement before to execute the conditional validation and checking.
	*/

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}