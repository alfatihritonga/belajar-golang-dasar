package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("Selesai")
}
