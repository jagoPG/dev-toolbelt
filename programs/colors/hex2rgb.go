package colors

import (
	"errors"
	"fmt"
	"jagoPG/utils/arguments"
	"regexp"
	"strconv"
)

// Given an Hex color with 6 characters, return the RGB value
// Example: hex2RGB(#FFFFFF) -> 255, 255, 255
func ConvertHexToRGB(args []string) (string, error) {
	isValid, err := arguments.ValidateArgs(args, 1)
	if !isValid {
		return "", err
	}
	if err := validateHexColor(args[0]); err != nil {
		return "", err
	}
	hex := Hex{args[0]}

	return hex.ToRGB()
}

func validateHexColor(arg string) error {
	if matched, _ := regexp.MatchString("#[0-9A-Fa-f]{6}", arg); !matched {
		return errors.New("does not match an hexadecimal color")
	}

	return nil
}

type Hex struct {
	r string
}

func (h *Hex) ToRGB() (string, error) {
	r, err := convertHexToInt(h.r[1:3])
	if err != nil {
		return "", err
	}

	g, err := convertHexToInt(h.r[3:5])
	if err != nil {
		return "", err
	}

	b, err := convertHexToInt(h.r[5:7])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v, %v, %v", r, g, b), nil
}

func convertHexToInt(hex string) (int64, error) {
	if result, err := strconv.ParseInt(hex, 16, 64); err != nil {
		return 0, errors.New("invalid hexadecimal number")
	} else {
		return result, nil
	}
}
