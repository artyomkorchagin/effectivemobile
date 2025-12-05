package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/app"
	"github.com/artyomkorchagin/effectivemobile/internal/config"
	"github.com/artyomkorchagin/effectivemobile/internal/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)
//test
// @title			Effective Mobile Task GO Junior
// @version			1.0
// @contact.name	Artyom Korchagin
// @contact.email	artyomkorchagin333@gmail.com
// @host			localhost:3000
// @BasePath		/

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	devMode := logger.Production
	if cfg.LogMode == "DEV" {
		devMode = logger.Development
	}

	zapLogger, err := logger.New(devMode)
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	defer zapLogger.Sync()

	zapLogger.Info("Starting application", zap.String("port", cfg.Server.Port))

	db, err := sql.Open("pgx", cfg.GetDSN())
	if err != nil {
		zapLogger.Fatal("failed to open database", zap.Error(err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			zapLogger.Error("error closing database", zap.Error(err))
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		zapLogger.Fatal("failed to ping database", zap.Error(err))
	}

	if err := app.RunMigrations(db, zapLogger); err != nil {
		zapLogger.Fatal("failed to run migrations", zap.Error(err))
	}

	application := app.New(cfg, db, zapLogger)
	if err := application.Run(); err != nil {
		zapLogger.Fatal("application error", zap.Error(err))
	}
}
