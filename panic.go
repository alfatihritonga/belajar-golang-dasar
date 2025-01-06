package main

import (
	"fmt"
)

type Checklog func(bool)

func runApplication(status bool, check Checklog) {
	defer endApplication() // function tetap dieksekusi meskipun aplikasi terjadi error

	fmt.Println("Starting Application...")
	check(status)

	// endApplication() // tidak akan dieksekusi ketika check(status) error
}

func endApplication() {
	fmt.Println("Application stop running.")
}

func loggingStatus(status bool) {
	if !status {
		panic("Ups.. Application failed to start.")
	}

	fmt.Println("Application running.")
}

func main() {
	runApplication(false, loggingStatus)
}
