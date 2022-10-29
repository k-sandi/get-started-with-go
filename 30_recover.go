/*
DEFER, PANIC and RECOVER

# Recover
> function yang bisa digunakan untuk menangkap data panic.
> with recover the panic processing will be stopped, then the program still running.
*/

package main 

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai!")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Validasi")
}