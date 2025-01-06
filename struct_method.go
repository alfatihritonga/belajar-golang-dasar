package main

import "fmt"

type Product struct {
	Name     string
	Price    int
	Category string
}

type Discount func(int) int

func (product Product) getDiscount(discount int) int {
	totalDiscount := (float32(discount) / float32(100)) * float32(product.Price)
	return int(totalDiscount)
}

func main() {
	product := Product{
		Name:     "Lenovo IdeaPad Slim 3",
		Price:    6600000,
		Category: "Laptop",
	}

	discount := 20
	total := product.Price - product.getDiscount(discount)

	fmt.Println(product.Name, "\tRp.", product.Price)
	fmt.Println("Dicount(%)\t:", discount)
	fmt.Println("Total\t\t: Rp.", total)
}
