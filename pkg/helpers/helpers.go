package helpers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

func RunMake(target, root string) error {
	cmd := exec.Command("make", target)
	cmd.Stderr = os.Stderr
	cmd.Dir = root

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run make %s: %v", target, err)
	}
	return nil
}

func GetProjectRoot() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return path, nil
}
