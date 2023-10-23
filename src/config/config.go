package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// Config stores configuration values for the application.
type Config struct {
	ApiUrl        string
	ApiKey        string
	Limit         int
	HexLength     int
	HashLength    int
	AddressLength int
}

// New creates a new Config object with values from environment variables or default values.
func New() *Config {
	loadEnv()
	return &Config{
		ApiUrl:        getStr("API_URL", "https://api.etherscan.io/api"),
		ApiKey:        getStr("API_KEY", ""),
		Limit:         getNum("LIMIT", 50),
		HexLength:     getNum("HEX_LENGTH", 9),
		HashLength:    getNum("HASH_LENGTH", 66),
		AddressLength: getNum("ADDRESS_LENGTH", 42),
	}
}

// loadEnv loads environment variables from the .env file.
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

// getStr retrieves a string value from the environment variables or returns a default value.

func getStr(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getNum retrieves an integer value from the environment variables or returns a default value.
func getNum(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		num, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return num
	}
	return defaultValue
}
