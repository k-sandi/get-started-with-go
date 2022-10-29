/*
FUNCTION AS PARAMETER or FUNCTION AS ARGUMEN
> Function tidak hanya bisa disimpan di dalam variabel sebagai value. 
> namun bisa sebagai parameter untuk function lain
*/

package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
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