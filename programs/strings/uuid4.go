package strings

import (
	"github.com/google/uuid"
)

func GenerateUUID4(args []string) (string, error) {
	uuid4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid4.String(), nil
}
