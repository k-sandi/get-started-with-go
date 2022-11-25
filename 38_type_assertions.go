/*
* TYPE ASSERTIONS
* adalah kemampuan mengubah tipe data menjadi tipe data yang diinginkan
* fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong. Secara spesifik mengubah interface kosong menjadi spesific balikan data type
* 
*/

package main 

import "fmt"

func random() interface{} {
	return "Ups"
}

// func main() {
// 	var result interface{} = random()
// 	var resultString string = result.(string) //this is type assertion to string
// 	fmt.Println(resultString)
// }

/* Note:
* Hati-hati terjadi panic (program akan terhenti) jika tipe data salah
* Untuk mengindarinya maka gunakan Type Assertions Switch
*/

func main() {
	var result interface{} = random()
	
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Unknown type")
	}
}


