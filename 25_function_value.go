/*
FUNCTION VALUE
> first class citizen
> also data type dan can saved into variable
*/

package main 

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Sandi")

	fmt.Println(result)
	fmt.Println(getGoodBye("Sandi"))
}