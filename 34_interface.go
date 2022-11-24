/*
INTERFACE
- tipe data abstract yang tidak memiliki implementasi langsung.
- interface berisi definisi2 method
- biasa digunakan sbg kontrak
*/

package main 

import "fmt"

type HasName interface {
	GetName() string // method dengan nama "GetName" dengan return berupa "string"
}

func SayHello(hasName HasName){
	fmt.Println("Hello ", hasName.GetName())
}

// func main() {
// 	var eko HasName

// 	SayHello(eko)
// }

/*
* IMPLEMENTASI INTERFACE
* setuap tipa data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai intrface tsb.
* Sehingga kita tidak perlu mengimplementasikan interface secara manual.
* Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface. kita harus menyebutkan secara eksplisit akan menggunakan interface mana.
*/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var whoami Person
	whoami.Name = "Sandi"

	SayHello(whoami)

	cat := Animal {
		Name: "Molly",
	}

	//note: bisa pakek struktur data implementasi Person

	SayHello(cat)
}


