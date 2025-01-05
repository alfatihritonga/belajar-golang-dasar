package main

import "fmt"

func checkRole(name string) bool {
	if name != "admin" {
		return false
	}
	return true
}

func main() {
	auth := checkRole("admin")

	if auth {
		fmt.Println("Akses diterima!")
	} else {
		fmt.Println("Akses ditolak!")
	}
}
