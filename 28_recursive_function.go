/*
RECURSIVE FUNCTION
> Calling self function
> Ex: Factorial (5! = 5x4x3x2x1)
*/

package main

import "fmt"

//without recursive function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//with recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	
	loop := factorialLoop(5)
	loopFactorial := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(loopFactorial)

}