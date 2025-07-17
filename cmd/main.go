package main

import (
	"log"

	"github.com/artyomkorchagin/effectivemobile/internal/router"
	"github.com/artyomkorchagin/effectivemobile/pkg/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	logger := logger.New()
	defer logger.CloseLogger()

	handler := router.NewHandler(nil, logger.Logger)
	r := handler.InitRouter()
	r.Run(":3000")
}
