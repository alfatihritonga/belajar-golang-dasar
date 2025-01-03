package main

import "fmt"

func main() {
	nilai16 := int16(128)
	nilai32 := int32(nilai16)
	nilai8 := int8(nilai16)

	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai8)

	name := "Alfatih"
	charA := name[0]
	stringA := string(charA)

	fmt.Println(name)
	fmt.Println(charA)
	fmt.Println(stringA)
}

// Note

// Jika nilai dari tipe data yang mau dikonversi melebihi batas maksimalnya,
// maka akan terjadi Number Overflow, dimana hasilnya akan dimulai dari nilai minimumnya
