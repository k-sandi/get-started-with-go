// PACKAGE TIME

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Local())

	utc := time.Date(2020, 10, 10, 10, 10 ,10, 10, time.UTC)
	fmt.Println(utc)

	// layout := "RFC3339"
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Println(parse)
}