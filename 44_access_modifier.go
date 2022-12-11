/*
ACCESS MODIFIER
- Di bahasa pemrograman lain biasanya ada keyword yang dapat digunakan untuk menentukan access modifier terhadap suatu function atau variabel.
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau nama variable.
- Jika namanya diawali dengan huruf besar, maka dapat diakses dari package lain.
- Jika namanya diawali dengan huruf kecil, maka tidak dapat diakses dari pakcage lain.
*/

package main

import (
	"get-started-with-go/helper"
	"fmt"
)

func main() {
	helper.SayHello("Sandi")
	// helper.sayGoodBye("Sandi") //error not exported because the name started by lowercase
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //error not exported because the name started by lowercase
}