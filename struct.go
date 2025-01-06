package main

import "fmt"

type Student struct {
	Id            int
	Name, Address string
	IsActive      bool
}

func main() {
	var fatih Student

	fatih.Id = 1
	fatih.Name = "Muhammad Al-fatih Ritonga"
	fatih.Address = "Medan"
	fatih.IsActive = true

	fmt.Println(fatih.Id)
	fmt.Println(fatih.Name)
	fmt.Println(fatih.Address)
	fmt.Println(fatih.IsActive)
	fmt.Println(fatih)

	syilfa := Student{
		Id:       2,
		Name:     "Syilfa Salsabila Siregar",
		Address:  "Medan",
		IsActive: true,
	}
	fmt.Println(syilfa)

	eko := Student{3, "Eko Kurniawan", "Indonesia", false}
	fmt.Println(eko)
}
