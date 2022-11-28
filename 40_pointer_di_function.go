/*
POINTER DI FUNCTION
- Saat kit membuat parameter di function, secara default adalah pass by value, artinya data akan dicopy lalu dikirim ke function tersebut.
- Oleh karena itu, jika mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah.
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut.
- Untuk melakukan ini, kita bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator "*" di parameternya.
*/

package main 

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	/* Kode Pass by Value */
	// address1 := Address{"Sumbawa", "Nusa Tenggara Barat", "Indoneisa"}
	// address2 := address1 

	// address2.City = "KSB"

	// fmt.Println(address1)
	// fmt.Println(address2)

	/* Kode Pass by Reference with Pointer */

	/* Operatir "&" */
	/* bisa pakai cara ini: */
	// address1 := Address{"Sumbawa", "Nusa Tenggara Barat", "Indoneisa"}
	// address2 := &address1 //cukup tambahkan operator "&" di awal variable
	/* === */
	/* atau ini: */
	// var address1 Address = Address{"Sumbawa", "Nusa Tenggara Barat", "Indonesia"}
	// var address2 *Address = &address1
	/* === */

	// address2.City = "KSB"

	// fmt.Println(address1)
	// fmt.Println(address2)

	/* Operator "*" */
	var address1 Address = Address{"Sumbawa", "Nusa Tenggara Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "KSB"

	*address3 = Address{"Kupang", "Nusa Tenggara Timur", "Indonesia"} // put the operator on this line

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	/* function new */
	var address4 *Address = new(Address)
	address4.City = "Mataram"
	address4.Province = "Nusa Tenggara Barat"
	fmt.Println(address4)

	var alamat = Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}

	var alamatPointer *Address = &alamat

	ChangeCountryToIndonesia(alamatPointer)

	fmt.Println(alamat)
}