package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int 
}

type UserSlice []User

func (val UserSlice) Len() int {
	return len(val)
}

func (val UserSlice) Less(i, j int) bool {
	return val[i].Age < val[j].Age
}

func (val UserSlice) Swap(i, j int) {
	val[i], val[j] = val[j], val[i]
}

func main() {
	users := []User {
		{"Eko", 30},
		{"Joko", 35},
		{"Budi", 31},
		{"Rudi", 25},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}