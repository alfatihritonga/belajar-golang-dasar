package main

import "fmt"

func main() {
	students := [...]string{
		"Alfatih",
		"Syilfa",
		"Arly",
		"Rudiansyah",
		"Dedi",
		"Alim",
	}

	slice1 := students[:] // slice semua data pada array
	fmt.Println(slice1)   // [ Alfatih, Syilfa, Arly, Rudiansyah, Dedi, Alim ]

	slice2 := students[3:6] // slice data dari index ke-3 sampai index ke-5 pada array
	fmt.Println(slice2)     // [ Rudiansyah, Dedi, Alim ]

	slice3 := students[4:] // slice data dari index ke-4 sampai index terakhir pada array
	fmt.Println(slice3)    // [ Dedi, Alim ]

	slice4 := students[:2] // slice data dari index pertama sampai index ke-1 pada array
	fmt.Println(slice4)    // [ Alfatih, Syilfa ]

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[:5]
	daySlice1[0] = "Senin Baru"
	daySlice1[1] = "Selasa Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Sabtu Baru")
	daySlice2 = append(daySlice2, "Minggu Baru")
	daySlice2 = append(daySlice2, "Libur Baru")

	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 7)
	newSlice[0] = "Senin"
	newSlice[1] = "Selasa"
	// newSlice[2] = "Rabu" // error, harusnya menggunakan append() untuk memasukkan data baru

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Rabu")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Kamis"

	fmt.Println(newSlice) // Kamis, Selasa, Rabu // karna, newSlice2 masih berkaitan dengan newSlice
	fmt.Println(newSlice2)

	newSlice3 := append(newSlice2, "Kamis")
	newSlice3 = append(newSlice3, "Jum'at")
	newSlice3 = append(newSlice3, "Sabtu")
	newSlice3 = append(newSlice3, "Minggu")
	newSlice3 = append(newSlice3, "Hari Baru")
	fmt.Println(newSlice3)
	fmt.Println(len(newSlice3))
	fmt.Println(cap(newSlice3))

	newSlice3[0] = "Senin"
	fmt.Println(newSlice3)
	fmt.Println(newSlice2) // newSlice[0] tetap kamis, karna newSlice3 sudah menyimpan array baru sehingga ketika data pertama diubah tidak berlaku dengan newSlice dan newSlice2
	fmt.Println(newSlice)

	// copy data slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
