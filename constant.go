package main

import "fmt"

func main() {
	const firstName = "Muhammad Al-fatih"
	const lastName = "Ritonga"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		role = "Fullstack Programmer"
	)
}

// Note

// const tidak dapat diubah nilainya
// const juga tidak error ketika tidak dipanggil didalam kode program
