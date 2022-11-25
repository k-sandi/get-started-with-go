/*
* ERROR INTERFACE
* Go mempunyai interface yang digunakan sebagai kontrak untuk membuat error dan nama interfacenya adalah error.
*
* MEMBUAT ERROR
* Go sudah menyediakan libary untuk membuat helper secara mudah, yang terdapat di package errors.
*/

package main 

import (
	"fmt"
	"errors"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh O")
	} else {
		result := nilai / pembagi
		return result, nil //karena interfaca, maka kita bisa gunakan nil
	}
}

func main() {
	/* Percobaan pertama */
	// var contohError error = errors.New("Upps Error")
	// fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 10)
	if err == nil { // Jika tidak error atau hasil error sama dengan nil
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

	hasil2, err := Pembagi(100, 0)
	if err == nil { // Jika tidak error atau hasil error sama dengan nil
		fmt.Println("Hasil", hasil2)
	} else {
		fmt.Println("Error", err.Error())
	}
}