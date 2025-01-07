package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func UpdateData(id string, data any) error {
	if id == "" {
		return &validationError{"missing parameter id when update data."}
	}

	if id != "1" {
		return &notFoundError{"id not found."}
	}

	return nil
}

func main() {
	err := UpdateData("12", nil)
	if err != nil {
		switch errorType := err.(type) {
		case *validationError:
			fmt.Println("Fatal Error:", errorType.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", errorType.Error())
		default:
			fmt.Println("Unknwon Error:", errorType.Error())
		}
	} else {
		fmt.Println("Success update data.")
	}
}
