package main

import "fmt"

func main() {
	student := map[string]string{
		"name":    "Al-fatih",
		"major":   "Information Systems",
		"address": "Medan",
	}

	fmt.Println(student)
	fmt.Println(student["name"])
	fmt.Println(student["major"])

	product := make(map[string]string)
	product["name"] = "Samsung A22"
	product["spec"] = "Lorem Ipsum"
	product["wrong"] = "Ups"

	fmt.Println(product)

	delete(product, "wrong")

	fmt.Println(product)
}
