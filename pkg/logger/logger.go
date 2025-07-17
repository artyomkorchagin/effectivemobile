package logger

import (
	"log"
	"os"
)

func New() *log.Logger {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	logger := log.New(log.Writer(), "online-shop", 0)
	logger.SetOutput(f)
	return logger
}
