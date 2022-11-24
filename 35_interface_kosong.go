/* #INTERFACE KOSONG
* Go-Lang bukan OOP
* Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya.
*
* PENGGUNAAN INTERFACE KOSONG DI GO-LANG
* Contoh :
* fmt.Println(a ...interface{}) 
* panic(v interface{})
* recover() interface{}
* dll
*
*/

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
}