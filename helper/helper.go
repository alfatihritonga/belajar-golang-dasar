package helper

import "fmt"

var version = "1.0.0" // hanya bisa diakses didalam package yg sama
var Application = "Golang"

// function dibawah tidak dapat diakses diluar package helper
func sayGoodBye(v string) string {
	return "Good Bye " + v
}

func SayHello(v string) string {
	return "Hello " + v
}

func Contoh() {
	fmt.Println(sayGoodBye("Fatih"))
	fmt.Println(version)
}
