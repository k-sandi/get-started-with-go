/*
FUNCTION IN GO

*/

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	// sayHello()
	// sayHello()

	for i := 0; i < 10; i++ {
		sayHello()
	}

}