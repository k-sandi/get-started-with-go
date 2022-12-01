/*
PACKAGE
- Package adalah tempat yang bisa digunakan mengorganisir kode program yang kita buat di Go-Lang.
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat.
- Package sendiri sebenarnya hanya direktori di OS kita.

IMPORT
- Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama.
- Jika kita ingin mengakses file Go-Lang yang berada di luar package, maka kita bisa menggunakan import.
*/

package main 

import "get-started-with-go/helper"

func main() {
	helper.SayHello("Sandi")
}