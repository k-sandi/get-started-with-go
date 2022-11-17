/*
STRUCT METHOD
=> Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function.
=> Menambahkan function ke dalam struct seakan-akan struct memiliki function (padahal tidak)
=> Method adalah function.
*/

package main 

import "fmt"

type Customer struct {
	Name, Address string 
	Age int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name,"My name is", customer.Name)
}

func main() {
	var data Customer 
	data.Name = "Sandi"
	data.Address = "South Jakarta"
	data.Age = 28

	data.sayHi("Joko")

	// fmt.Println(data.Name)
	// fmt.Println(data.Address)
	// fmt.Println(data.Age)

	// // Struct Literal
	// joko := Customer {
	// 	Name: "Joko",
	// 	Address: "Cirebon",
	// 	Age: 35,
	// }

	// fmt.Println(joko)

	// budi := Customer {"Budi","Jakarta",35}
	// fmt.Println(budi)
}