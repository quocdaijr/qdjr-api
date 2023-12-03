package helpers

import (
	"os"
)

type BaseHelper struct{}

func (_ BaseHelper) GetEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
