package main

import "fmt"

func main() {
	var a = 10
	var b = 15

	result := a > b
	fmt.Println(a, ">", b, "=", result)

	result = a < b
	fmt.Println(a, "<", b, "=", result)

	result = a >= b
	fmt.Println(a, ">=", b, "=", result)

	result = a <= b
	fmt.Println(a, "<=", b, "=", result)

	result = a == b
	fmt.Println(a, "==", b, "=", result)

	result = a != b
	fmt.Println(a, "!=", b, "=", result)

	// perbandingan nilai string
	var name1 = "Alfatih"
	var name2 = "Al-fatih"

	result = name1 == name2
	fmt.Println(result)
}
