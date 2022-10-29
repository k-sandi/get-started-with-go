/*
DEFER, PANIC and RECOVER

# DEVER
> function yang dapat dijadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi.
> Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi.
*/

package main 

import "fmt"

// Kode Program Defer
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() //setelah selesai runapplication, maka harus diakhir dengan mengeksekusi function logging
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
	// logging()
}

func main() {
	runApplication(10)
}