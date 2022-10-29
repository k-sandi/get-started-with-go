/*
DEFER, PANIC and RECOVER

# PANIC
> function yang dapat untuk menghentikan program
> biasanya di-call ketika terjadi error pada saat program kita berjalan
> Saat panic function di-call, program akan terhenti, namun defer function tetap akan dieksekusi.
*/

package main 

import "fmt"

func endApp() {
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
}