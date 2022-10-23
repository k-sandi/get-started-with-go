/*
BREAK & CONTINUE
1. Break: Stop the looping for all the script after that
2. Continue: Skip to next block or loop
*/

package main

import "fmt"

func main() {

	//BREAK
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke", i)
	}

	//CONTINUE
	for j := 0; j < 10; j++ {
		if j % 2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke", j)
	}
}