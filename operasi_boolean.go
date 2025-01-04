package main

import "fmt"

func main() {
	var nilaiUAS = 80
	var nilaiAbsensi = 75

	var lulusNilaiUAS bool = nilaiUAS >= 80
	var lulusNilaiAbsensi bool = nilaiAbsensi >= 80

	var lulus bool = lulusNilaiUAS && lulusNilaiAbsensi

	fmt.Println(lulus)
}
