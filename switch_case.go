package main

import "fmt"

func main() {
	name := "Syilfa"

	switch name {
	case "Alfatih":
		fmt.Println("Hello, Alfatih")
	case "Syilfa":
		fmt.Println("Hello, Syilfa")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	number := 0

	switch {
	case number > 80:
		fmt.Println("Angka terlalu tinggi")
	case number > 60:
		fmt.Println("Angka masih tinggi")
	case number > 45:
		fmt.Println("Angka sudah pas")
	case number > 20:
		fmt.Println("Angka masih rendah")
	case number > 0:
		fmt.Println("Angka terlalu rendah")
	default:
		fmt.Println("Angka tidak boleh 0")
	}
}
