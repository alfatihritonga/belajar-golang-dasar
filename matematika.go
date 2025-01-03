package main

import "fmt"

func main() {
	// operasi matematika
	a := 10
	b := 5

	penjumlahan := a + b
	pengurangan := a - b
	perkalian := a * b
	pembagian := a / b
	sisaBagi := a % b

	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(pembagian)
	fmt.Println(sisaBagi)

	// augmented assignment
	c := 10

	c += 20
	fmt.Println(c)

	c -= 5
	fmt.Println(c)

	c *= 2
	fmt.Println(c)

	c /= 2
	fmt.Println(c)

	c %= 5
	fmt.Println(c)

	// unary operator
	d := 10
	d++
	fmt.Println(d)

	d--
	fmt.Println(d)
}
