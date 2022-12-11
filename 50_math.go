/*
PACKAGE MATH
*/

package main 

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Round cek untuk nilai paling dekat ke atas: ", math.Round(1.7))
	fmt.Println("Round cek untuk paling dekat ke bawah: ", math.Round(1.3))

	fmt.Println("Floor cek untuk membulatkan nilai ke atas secara paksa: ", math.Round(1.7))
	fmt.Println("Ceil cek untuk membulatkan nilai ke atas secara paksa: ", math.Round(1.3))

	fmt.Println("Max cek untuk mengambil nilai terbesar: ", math.Max(11,12))
	fmt.Println("Min cek untuk mengambil nilai terkecil: ", math.Min(11,12))
}