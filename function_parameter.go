package main

import "fmt"

func greetingTo(name string, role string) {
	fmt.Println("Selamat Datang,", name)
	fmt.Println("Anda login sebagai", role)
}

func main() {
	greetingTo("Alfatih", "Admin")
	greetingTo("Syilfa", "Mahasiswa")
}
