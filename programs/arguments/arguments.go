package arguments

import (
	"errors"
)

func ValidateArgs(args []string, amountRequired int) (bool, error) {
	if len(args) != amountRequired {
		return false, errors.New("invalid amount of arguments")
	}

	return true, nil
}
