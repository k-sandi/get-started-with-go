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

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "MayLagi"
	fmt.Println(months)
	
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Kurnia")
	fmt.Println(slice3)
	slice3[1] = "Bukan-December"
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Kurnia"
	newSlice[1] = "Sandi"

	// fmt.Println(fmt.Sprintf("value: %d", newSlice))
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// "Length: " + fmt.Println(len(newSlice))
	// "Capacity: " + fmt.Println(cap(newSlice))

	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	/* 
	Hati-hati membuat array
	*/

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}