package environment

import (
	"os"
	"strings"
)

// GetEnv get key environment variable if exist, otherwise return defaultValue
func GetStrEnv(key, defaultValue string) string {
	if value := os.Getenv(key);  value != "" {
		return value 
	}
	return defaultValue
}

func GetBoolEnv(key string, defaultValue bool) bool{
	if value := os.Getenv(key);  value != "" {
		if strings.ToLower(value) == "true" {
			return true
		}
		if strings.ToLower(value) == "false" {
			return false
		}
	}
	return defaultValue
}
