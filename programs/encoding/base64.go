package encoding

import (
	"encoding/base64"
	"jagoPG/utils/arguments"
)

func EncodeBase64(args []string) (string, error) {
	isValid, err := arguments.ValidateArgs(args, 1)
	if !isValid {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(args[0])), nil
}

func DecodeBase64(args []string) (string, error) {
	isValid, err := arguments.ValidateArgs(args, 1)
	if !isValid {
		return "", err
	}
	result, err := base64.StdEncoding.DecodeString(args[0])

	return string(result), err
}
