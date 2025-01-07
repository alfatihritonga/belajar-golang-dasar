package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	res := helper.SayHello("Fatih")
	fmt.Println(res)

	// fmt.Println(helper.version) // tidak dapat diakses
	fmt.Println(helper.Application) // bisa diakses diluar packagenya
	// fmt.Println(helper.sayGoodBye("Fatih")) // tidak dapat diakses
	helper.Contoh()
}
