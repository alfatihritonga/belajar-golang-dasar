package main

import "fmt"

func main() {
	type NoKTP string

	var ktpFatih NoKTP = "123456789"

	// konversi tipe data ke tipe data yang kita buat
	contoh := "987654321"
	ktpContoh := NoKTP(contoh)

	fmt.Println(ktpFatih)
	fmt.Println(ktpContoh)
}
