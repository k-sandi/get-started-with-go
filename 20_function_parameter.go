/*
FUNCTION PARAMETER
1. Parameter wajib ada isinya.
*/

package main 

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Kurnia", "Sandi")
}