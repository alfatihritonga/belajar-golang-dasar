package main

import "fmt"

func main() {
	var salaries [3]int

	salaries[0] = 12000000
	salaries[1] = 7000000
	salaries[2] = 5000000

	fmt.Println(salaries)
	fmt.Println(salaries[0])

	// membuat array secara langsung
	var countries = [...]string{
		"Indonesia",
		"Malaysia",
		"Thailand",
		"Singapore",
		"Vietnam",
		"Laos",
	}

	fmt.Println(countries)
	fmt.Println(len(countries))
}
