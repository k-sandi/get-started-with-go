/* 
Tipe Data Array


*/

package main

import "fmt"

func main() {
	var names [3] string

	names[0] = "Kurnia"
	names[1] = "Sandi"
	names[2] = "Booriel"

	//fmt.Println(names[0])
	//fmt.Println(names[1])
	//fmt.Println(names[2])

	var values = [3]int {
		90,
		80,
		95,
	}

	//fmt.Println(values)
	//fmt.Println(values[1]);

	/*
	Function Array
	len(array)
	*/

	fmt.Println(len(names))
	fmt.Println(len(values))
}