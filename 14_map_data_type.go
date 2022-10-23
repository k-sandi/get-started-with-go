/*
Map => kay (unique) and value
Map can have many values
*/

package main 

import "fmt"

func main() {
	person := map[string]string{
		"name": "Kurnia",
		"address": "Sumbawa",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	/* 
		Function Map
		1. len(map) 					|=> get count of data
		2. map[key] 					|=> get data from map by key
		3. map[key]=value 				|=> change value in spesific key
		4. make(map[TypeKey]TypeValue) 	|=> create a new map
		5. delete(map,key)				|=> delete map data by key
	*/

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Sandi"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}



