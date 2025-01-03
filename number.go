package main

import "fmt"

func main() {
	// int
	fmt.Println("int8 : ", -128, "s/d", 127)
	fmt.Println("int16 : ", -32768, "s/d", 32767)
	fmt.Println("int32 : ", -2147483648, "s/d", 2147483647)
	fmt.Println("int64 : ", -9223372036854775808, "s/d", 9223372036854775807)

	// float
	fmt.Println("Floating Point : ", 12.89)
}

// Note

// di Golang hanya ada 2 jenis tipe data Number:
// - Integer
// - Floating Point

// Integer terbagi menjadi 2:
// - int
// - uint (unsigned integer)

// int dapat menyimpan nilai minus (-) sedangkan uint tidak

// Alias
// byte => uint8
// rune => int32
// int => minimal int32
// uint => minimal uint32
