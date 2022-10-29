/*
CLOSURES
> Kemampuan function berinteraksi dengan data-data disekitarnya dalam scope yang sama.
> Harap gunakan dengan bijak!
*/

// scope: lingkup kerja variabel atau function

package main 

import "fmt"

func main() {
	name := "sandi"
	counter := 0
	increment := func() {
		name := "budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}