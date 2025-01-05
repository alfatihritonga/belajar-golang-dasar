package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func toxicFilter(name string) string {
	switch name {
	case "Anjing":
		return "***"
	case "Babi":
		return "***"
	default:
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", toxicFilter)

	filter := toxicFilter
	sayHelloWithFilter("Babi", filter)
	sayHelloWithFilter("No Toxic", filter)
}
