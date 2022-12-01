/*
POINTER DI METHOD

- Walaupun  method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value.
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method.
*/

package main 

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method:", man.Name)
}

func main() {
	sandi := Man{"Sandi"}
	sandi.Married()

	fmt.Println("Di akhir jadi:", sandi.Name)
}