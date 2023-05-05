package utils

import "os"

func GetEnv(key,defaultValue string) string {
	result := os.Getenv(key)

	if result == ""{
		return defaultValue
	}

	return result
}
