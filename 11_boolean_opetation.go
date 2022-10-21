/* 
Operasi Boolean adalah operasi yg hanya bisa dilakukan pada boolean saja
Beberapa operasi boolean : &&, || dan !
****************
T && T = T
T && F = F
F && T = F
F && F = F
****************
T || T = T
T || F = T
F || T = T
F || F = F
****************
! T = F
! F = T
*/

package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool
	var lulusAbsensi bool
	var lulus bool 

	lulusNilaiAkhir = nilaiAkhir > 80
	lulusAbsensi = absensi > 80

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	lulus = lulusAbsensi && lulusNilaiAkhir
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 90 && absensi >= 80)
}