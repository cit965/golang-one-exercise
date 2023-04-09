package main

import (
	"errors"
	"fmt"
)

const badInput = "abc"

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

func main() {
	input := badInput

	err := validateInput(input)
	var badInputErr *BadInputError

	if errors.Is(err, badInputErr) {
		fmt.Printf("is bad input error occured: %s\n", badInputErr)
	}

	if errors.As(err, &badInputErr) {
		fmt.Printf("as bad input error occured: %s\n", badInputErr)
	}
}
