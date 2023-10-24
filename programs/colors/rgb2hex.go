package colors

import (
	"fmt"
	"jagoPG/utils/arguments"
	"strconv"
)

// Given a RGB color, return the hex value
// Example: rgb2hex(255, 255, 255) -> #FFFFFF
func ConvertRGBToHex(args []string) (string, error) {
	isValid, err := arguments.ValidateArgs(args, 3)
	if !isValid {
		return "", err
	}

	r, err := argToRGB(args)
	if err != nil {
		return "", err
	}

	return r.ToHex(), nil
}

func argToRGB(args []string) (*RGB, error) {
	var r, g, b uint64
	var err error

	if r, err = strconv.ParseUint(args[0], 10, 64); err != nil {
		return nil, err
	}
	if g, err = strconv.ParseUint(args[1], 10, 64); err != nil {
		return nil, err
	}
	if b, err = strconv.ParseUint(args[2], 10, 64); err != nil {
		return nil, err
	}

	return &RGB{r, g, b}, nil
}

type RGB struct {
	r, g, b uint64
}

func (r *RGB) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", r.r, r.g, r.b)
}
