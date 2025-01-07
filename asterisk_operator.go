package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Medan",
		Province: "Sumatera Utara",
		Country:  "Indonesia",
	}

	// address2 := &address1
	var address2 *Address = &address1
	address2.City = "Deli Serdang"

	fmt.Println(address1) // ikut berubah
	fmt.Println(address2)

	*address2 = Address{
		City:     "Pekanbaru",
		Province: "Riau",
		Country:  "Indonesia",
	}

	fmt.Println(address1) // address1 sekarang memorinya di address2
	fmt.Println(address2)
}
