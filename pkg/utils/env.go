package utils

import (
	"os"
)

func Getenv(key string) string {
	value:= os.Getenv(key)
	return value
}
