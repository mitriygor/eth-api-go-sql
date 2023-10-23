package helpers

import (
	"eth-api/src/helpers/logger"
	"fmt"
	"regexp"
	"strconv"
)

// IsInt checks if the given string can be converted to an integer.
func IsInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// IsHex checks if the given string is a valid hexadecimal value.
func IsHex(s string) bool {
	pattern := `^0x[0-9a-fA-F]+$`
	match, _ := regexp.MatchString(pattern, s)
	return match && len(s) < 32
}

// StringToInt converts the given string to an integer.
func StringToInt(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		logger.Error("eth-api::ERROR::StringToInt", "error", err)
		return -1
	}

	return intValue
}

// HexToInt converts the given hexadecimal string to an integer.
func HexToInt(hexStr string) (int, error) {
	n, err := strconv.ParseInt(hexStr, 0, 64)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

// IntToHex converts the given integer to a hexadecimal string.
func IntToHex(num int) string {
	n := int64(num)
	return fmt.Sprintf("0x%x", n)
}
