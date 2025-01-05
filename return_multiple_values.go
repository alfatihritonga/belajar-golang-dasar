package main

import "fmt"

func getFullName() (string, string) {
	return "Al-fatih", "Ritonga"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// menghiraukan return value
	firstName, _ := getFullName()
	fmt.Println("Hello,", firstName)
}
