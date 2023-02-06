package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // WIth this we can manage the error
	}
	return a / b, nil
}

func main() {

	result, err := div(5, 0) // Receive the error then evaluate it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
