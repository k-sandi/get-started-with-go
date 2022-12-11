/*
PACKAGE STRCONV (STRING CONVERSION)

convert data type from string to any data type, or any data type to string.

https://golang.org/pkg/strconv
*/

package main 

import (
	"fmt"
	"strconv"
)

func main() {

	/*
	* you must define the value in real value but the datatype was string.
	* do : "true" or "1" for true, "false" or "0" for false.
	* don't : "salah", "benar" and any language for true or false value.
	*/
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}


	/* 
	* in number parse you must define from the number value in string, base int and bitSize
	* argument 1: number value in string
	* argument 2: base int or base size (decimal: 10, binary:2, octal:8, hexadecimal: 16)
	* argument 3: bitSzie (32/64)
	*/
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(10000000, 10)
	fmt.Println(value)

	// Atoi or Itoa for convert string to integer (this is simple)
	valueInt, err := strconv.Atoi("2000000")
	fmt.Println(valueInt)

}