package main

import "fmt"

func main() {
	nilaiUAS := 80
	nilaiUTS := 80

	if nilaiUAS >= 80 && nilaiUTS >= 80 {
		fmt.Println("Selamat, anda lulus semester ini!")
	} else if nilaiUAS >= 80 || nilaiUTS >= 80 {
		fmt.Println("Anda lulus bersyarat")
	} else {
		fmt.Println("Maaf, anda gagal dalam semester ini")
	}
}
