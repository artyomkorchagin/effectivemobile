package helpers

import (
	"log"
	"os"
	"time"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func ParseTime(s string) (time.Time, error) {
	parsedTime, err := time.Parse("01-2006", s)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func GetProjectRoot() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return path, nil
}
