package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	fmt.Println("Selesai")

	days := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}

	// for biasa
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// for range
	for index, day := range days {
		fmt.Println("index ke-", index, ":", day)
	}

	// for range mengabaikan index dari data collectionnya
	for _, day := range days {
		fmt.Println(day)
	}
}
