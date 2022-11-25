/*
* NIL
* Go untuk nil akan dibuatkan secara otomatis default valuenya.
* Nilai Nil pada Go hanya bisa digunakan dibeberapa tipe data, seperti interface, function, map, slice, pointer dan chanel
*/

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	//percobaan nil di awal
	// var person map[string] string = nil
	// fmt.Println(person)

	//setelah implementasi function NewMap
	// person := NewMap("Eko")
	// fmt.Println(person)

	//jika tanpa nil
	var person map[string]string = NewMap("Sandi")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}