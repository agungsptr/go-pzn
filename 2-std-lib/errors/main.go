package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found error")
)

func saveData(id string) error {
	if id == "" {
		return ErrValidation
	}

	if id != "agung" {
		return ErrNotFound
	}

	return nil
}

func main() {
	err := saveData("1")

	if err != nil {
		if errors.Is(err, ErrValidation) {
			fmt.Printf("Error: %s\n", err.Error())
		} else if errors.Is(err, ErrNotFound) {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Printf("Error: %s\n", err.Error())
		}
		return
	}

	fmt.Println("Success")
}
