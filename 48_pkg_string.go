/*
PACKAGE STRING
- manipulate the string values
- golang.go/pkg/string
*/

package main 

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Contains cek: ", strings.Contains("Kurnia Sandi Keitaro", "Kurnia"))
	fmt.Println("Contains cek: ", strings.Contains("Kurnia Sandi Keiraro", "Budi"))

	fmt.Println("Split cek: ", strings.Split("Kurnia Sandi Keitaro", " "))

	fmt.Println("ToLower cek: ", strings.ToLower("Kurnia Sandi Keitaro"))
	fmt.Println("ToUpper cek: ", strings.ToUpper("Kurnia Sandi Keitaro"))
	fmt.Println("ToTitle cek: ", strings.ToTitle("Kurnia Sandi Keitaro"))

	fmt.Println("before Trim cek: ", "                 Kurnia Sandi                  ")
	fmt.Println("after Trim cek: ", strings.Trim("                 Kurnia Sandi                  ", " ")) // mengilangkan spasi di kiri dan kanan

	fmt.Println("ReplaceAll cek: ", strings.ReplaceAll("Kurnia Sandi Kurnia", "Kurnia", "Keitaro"))
}