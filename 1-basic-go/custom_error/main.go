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

func saveData(id string) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "1" {
		return &notFoundError{"not found error"}
	}
	return nil
}

func main() {
	err := saveData("0")
	if err != nil {
		// Using if
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundErr.Error())
		} else {
			fmt.Println("Unknown error:", err.Error())
		}

		// Using switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", err.Error())
		}

		return
	}

	fmt.Println("Success")
}
