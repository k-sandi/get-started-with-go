/**
STRUCT
> Struct adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data dalam satu kesatuan.
> Struct biasanya representasi data dalam program aplikasi yang kita buat
> Data di struct disimpan di dalam field
> Sederhananya struct adlah kumpulan dari field
> Struct itu sama dengan class jika di OOP
*/

package main 

import "fmt"

type Customer struct {
	Name, Address string 
	Age int
}

/**
Membuat Data Struct
> Struct adalah template/prototype data
> Struct tidak bisa langsung digunakan
> Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

func main() {
	var data Customer 
	data.Name = "Sandi"
	data.Address = "South Jakarta"
	data.Age = 28

	fmt.Println(data.Name)
	fmt.Println(data.Address)
	fmt.Println(data.Age)

	// Struct Literal
	joko := Customer {
		Name: "Joko",
		Address: "Cirebon",
		Age: 35,
	}

	fmt.Println(joko)

	budi := Customer {"Budi","Jakarta",35}
	fmt.Println(budi)
	
}