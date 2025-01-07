package main

import (
	"errors"
	"fmt"
)

func Division(v float64, d float64) (float64, error) {
	if d == 0 {
		return 0, errors.New("Cannot divide by the value 0")
	} else {
		return v / d, nil
	}
}

func main() {
	res, err := Division(20, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(res)
	}
}
