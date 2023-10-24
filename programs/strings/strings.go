package strings

import (
	"fmt"
	"unicode/utf8"
)

func Count(s []string) (string, error) {
	return fmt.Sprintf("%d", utf8.RuneCountInString(s[0])), nil
}
