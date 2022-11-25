/*
* POINTER
* # Past by Value
* Secara default di Go semua variable itu di passing by value, bukan by reference
* Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi valuenya.
*
* # Pointer
* Pointer adalah kemampuan membuat reference ke lokasi data di memorry yang sama tanpa menduplikasi data yang sudah ada.
* Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
*
* Bagiamana cara melakukan pointer?
*
* # Cara I : Operator "&"
* Bikin pointer dari variable yang sudah ada
* 
* # Cara II : Operator "*"
* Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
* Semua variable yang mengacu ke data yang sama tidak akan berubah.
* Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator "*".
*
* # Cara III: Function New
* Function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
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
}
