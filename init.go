package main

import (
	"belajar-golang-dasar/database"
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
