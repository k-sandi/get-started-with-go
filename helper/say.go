package helper

import "fmt"

var version = "1.0.0"
var Application = "belajar golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Good bye", name)
}
