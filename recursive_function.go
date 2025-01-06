package main

import "fmt"

// menghitung nilai faktorial dengan func menggunakan looping
func faktorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// menghitung nilai faktorial dengan func recursive
func faktorialRecursive(value int) int {
	if value != 1 {
		return value * faktorialLoop(value-1)
	}
	return 1
}

func main() {
	fmt.Println(faktorialLoop(20))
	fmt.Println(faktorialRecursive(20))
}
