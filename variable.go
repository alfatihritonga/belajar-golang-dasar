package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Al-fatih Ritonga"
	fmt.Println(name)

	// Deklarasi variable tanpa menyebutkan tipe datanya
	var role = "Fullstack Programmer"
	fmt.Println(role)

	// Deklarasi variable tanpa menggunakan kata kunci var
	salary := 60000000
	fmt.Println(salary)

	// Deklarasi Multiple Variable
	var (
		firstName  = "Muhammad"
		middleName = "Al-fatih"
		lastName   = "Ritonga"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
