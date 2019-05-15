package db

import (
	"fmt"
	"regexp"
)

var validKey = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func validateKey(key string) error {
	if validKey.MatchString(key) {
		return nil
	}

	return fmt.Errorf("Invalid key, must be alphanumeric")
}
