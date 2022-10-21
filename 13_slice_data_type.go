/*
Tipe Data Slice
- potongan dari array
- mirip dengan array. Tapi, ukuran slice bisa berubah
- slice dan array selalu terkoneksi. Jadi, dibelakang slice itu ada array

Detail:
- Data Pointer adalah acuan data pertama yang ada di array
- Data Length adalah panjang dari slice
- Capacity adalah kapasitas dari slice, dimana length todak boleh lebih dari capacity

Membuat Slice
- array[low:hight] = membuat slice dari array index low s/d index sebelum high.
- array[low:] = membuat slice dari array index low s/d index akhir di array
- array[:high] = membuat slice dari array dimulai index 0 s/d index sebelum high
- array[:] = Membuat slice dari array dimulai index 0 s/d index akhir di array
*/

package main

import "fmt"

func main() {
	var months = [...] string {
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[1] = "MayUbah"
	// fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)
}