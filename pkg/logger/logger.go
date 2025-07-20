package logger

import (
	"log"
	"os"
)

type Logger struct {
	Logger *log.Logger
	file   *os.File
}

// This is for writing logs into file
// Apperantly, they don't do that in prod
func New() Logger {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger := log.New(log.Writer(), "effectivemobile --- ", 0)
	logger.SetOutput(f)
	return Logger{logger, f}
}

func (l Logger) Close() {
	l.file.Close()
}
