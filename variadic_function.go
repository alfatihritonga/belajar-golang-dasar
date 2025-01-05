package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(5, 10, 15, 20)
	fmt.Println(total)

	// variadic arguments dengan data slice
	numbers := []int{5, 10, 15, 20}
	total = sumAll(numbers...)
	fmt.Println(total)
}
