package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Fatih"}
	man.Married()
	fmt.Println(man.Name)
}
