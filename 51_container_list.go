package main

import (
	"fmt"
	"container/list"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khaneddy")
	data.PushFront("Budi")

	// for e := data.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}