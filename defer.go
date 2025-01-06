package main

import "fmt"

func logging() {
	fmt.Println("Application running!")
}

func runApp() {
	defer logging()
	fmt.Println("Starting Application...")
}

func main() {
	runApp()
}
