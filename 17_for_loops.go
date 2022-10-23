/*
nothing another loop in Go like a While Loop. Only For Loops in GO
*/

package main 

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter, " in for loops section.")
		counter++
	}

	/*
	FOR LOOPS with STATEMENT
	1. before(init) statenment
	2. after(post) statement
	*/

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i, " in for loops section with statement.")
	}

	/*
	FOR RANGE
	*/

	//before in for range:
	slice := []string{"Kurnia", "Sandi", "Programmer"}

	for s := 0; s < len(slice); s++ {
		fmt.Println(slice[s])
	}
	//end

	for j, value := range slice {
		fmt.Println("Index", j, "=", value)
	}

	//print result of FOR LOOPS without index
	for _, valueK := range slice {
		fmt.Println(valueK)
	}

	//using map
	persons := make(map[string]string)
	persons["name"] = "Kurnia"
	persons["title"] = "Programmer"

	for key, value := range persons {
		fmt.Println(key, "=", value)
	}

}