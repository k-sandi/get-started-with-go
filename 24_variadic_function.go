/*
VARIADIC FUNCTION
> function yang memiliki variabel arguments (varargs) 
> Dynamic params
*/

package main 

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	//Slice Paramater
	slice := []int{10, 20, 30, 40, 50}
	totalSlice := sumAll(slice...)
	fmt.Println(totalSlice)
}