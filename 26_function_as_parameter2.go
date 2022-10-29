/*
FUNCTION AS PARAMETER or FUNCTION AS ARGUMEN (2)

FUNCTION TYPE DECLARATIONS
> Kadang jika function terlalu panjang, cukup ribet untuk menuliskannya di dalam parameter.
> Type Declaration juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter.
*/

package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Sandi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}