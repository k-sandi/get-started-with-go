/*
FUCTION RETURN VALUE
1) Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tsb.
2) Jika function tersebut dideklarasikan  dengan tipe data pengembalian, maka wajib di dalam functionnya kita harus mengembalikan data,
3) Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya.
*/

package main 

import (
	"strconv"
	"fmt"
)

func person(firstName string, age int) string {
	result := "Nama saya " + firstName + "dan Usia saya " + strconv.Itoa(age) 
	return result
}

func main() {
	result := person("Kurnia", 25)
	fmt.Println(result)
}