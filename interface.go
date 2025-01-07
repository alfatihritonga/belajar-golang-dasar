package main

import "fmt"

type HasName interface {
	getName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (p Person) getName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) getName() string {
	return a.Name
}

func main() {
	person := Person{Name: "Fatih"}
	SayHello(person)

	animal := Animal{Name: "Kucing Akmal"}
	SayHello(animal)
}
