package main

import "fmt"

func random() any {
	// return "OK"
	// return 22
	return true
}

func main() {
	res := random()

	// resultString := res.(string)
	// fmt.Println(resultString)

	// resultInt := res.(int)
	// fmt.Println(resultInt)

	switch value := res.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	case bool:
		fmt.Println("bool:", value)
	default:
		fmt.Println("unknown", value)
	}
}
