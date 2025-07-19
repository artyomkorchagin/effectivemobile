package main

import (
	"database/sql"
	"log"

	"github.com/artyomkorchagin/effectivemobile/internal/config"
	"github.com/artyomkorchagin/effectivemobile/internal/router"
	servicesubscription "github.com/artyomkorchagin/effectivemobile/internal/services/subscription"
	psqlsubscription "github.com/artyomkorchagin/effectivemobile/internal/storage/postgresql"
	"github.com/artyomkorchagin/effectivemobile/pkg/logger"
	"github.com/joho/godotenv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

//	@title			Effective Mobile Task GO Junior
//	@version		1.0

//	@contact.name	Artyom Korchagin
//	@contact.email	artyomkorchagin333@gmail.com

//	@host		localhost:3000
//	@BasePath	/

func main() {
	mylogger := logger.New()
	defer mylogger.Close()

	db, err := sql.Open("pgx", config.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	subRepo := psqlsubscription.NewRepository(db)
	subSvc := servicesubscription.NewService(subRepo)

	handler := router.NewHandler(subSvc, mylogger.Logger)
	r := handler.InitRouter()
	r.Run(":3000")
}
