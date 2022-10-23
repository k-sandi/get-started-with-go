/*
SWITCH EXPRESSION
*/

package main 

import "fmt"

func main() {
	
	name := "Kurniawam"

	switch name {
	case "Sandi":
		fmt.Println("Hello Sandi")
	case "Joko":
		fmt.Println("Hi Joko")
	default:
		fmt.Println("Hi Boleh Kenalan?")
	}

	/*
	SWITCH WITH SORT STATEMENT
	*/

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang!")
	case false:
		fmt.Println("Nama sudah benar")
	}

	/*
	SWITCH WITHOUT CONDITION
	*/
	
	length2 := len(name)

	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang!")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang!")
	default:
		fmt.Println("Nama sudah benar kisanak")
	}

}