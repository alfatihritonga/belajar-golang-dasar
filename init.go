package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	connect := database.ConnectDatabase()

	if !connect {
		fmt.Println("Failed to connect database.")
	} else {
		fmt.Println("Database connection successful")
	}
}
