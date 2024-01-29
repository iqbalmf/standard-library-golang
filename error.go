package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("Not Found Error")
)

func CheckErrorById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "iqbal" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := CheckErrorById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
