package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// kode pass by value
	// address1 := Address{
	// 	City:     "Medan",
	// 	Province: "Sumatera Utara",
	// 	Country:  "Indonesia",
	// }

	// address2 := address1
	// address2.City = "Deli Serdang"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah menjadi Deli Serdang

	// kode pass by reference
	address1 := Address{
		City:     "Medan",
		Province: "Sumatera Utara",
		Country:  "Indonesia",
	}

	// address2 := &address1
	var address2 *Address = &address1
	address2.City = "Deli Serdang"

	fmt.Println(address1) // berubah
	fmt.Println(address2)
}
