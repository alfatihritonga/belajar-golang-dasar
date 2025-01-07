package main

import "fmt"

func funcReturnAny() interface{} {
	// return 1

	// return "any"

	// return true

	// return map[string]string{}

	// return []string{}

	return [...]string{}
}

func main() {
	any := funcReturnAny()
	fmt.Println(any)
}
