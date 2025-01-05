package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// contoh
	bilGenap := make([]int, 0, 10)

	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		bilGenap = append(bilGenap, i)
	}
	fmt.Println("Bilangan genap dari 0 - 10 =", bilGenap)
}
