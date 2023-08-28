package main

import "fmt"

func main() {
	var ujian int = 80
	var absensi int = 80

	var lulusUjian bool = ujian >= 80
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusUjian && lulusAbsensi

	fmt.Println(lulus)
}
