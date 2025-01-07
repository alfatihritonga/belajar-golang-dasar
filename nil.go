package main

import "fmt"

// func dataString() string {
// 	return nil
// }

func dataInterface() interface{} {
	return nil
}

func dataMap() map[string]string {
	return nil
}

func dataSlice() []int {
	return nil
}

func main() {
	fmt.Println(dataInterface())

	dataMap := dataMap()
	if dataMap == nil {
		fmt.Println("Data Map Kosong.")
	} else {
		fmt.Println("Data Map Ada.")
	}

	dataSlice := dataSlice()
	if dataSlice == nil {
		fmt.Println("Data Slice Kosong.")
	} else {
		fmt.Println("Data Slice Ada.")
	}
}
