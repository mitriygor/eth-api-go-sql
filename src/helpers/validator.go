package helpers

import (
	"regexp"
	"strings"
)

// IsValidIdentifier checks if the given identifier is valid.
// It checks if the identifier has the correct prefix, length, and format.
func IsValidIdentifier(id string, length int) bool {
	if !strings.HasPrefix(id, "0x") && !strings.HasPrefix(id, "0X") {
		return false
	}

	if len(id) != length {
		return false
	}

	matched, err := regexp.MatchString("^[0-9a-fA-F]+$", id[2:])
	if err != nil {
		return false
	}

	return matched
}
