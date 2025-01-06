package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
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
