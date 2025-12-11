package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/app"
	"github.com/artyomkorchagin/effectivemobile/internal/config"
	"github.com/artyomkorchagin/effectivemobile/internal/logger"
	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

// @title			Effective Mobile Task GO Junior
// @version			1.0
// @contact.name	Artyom Korchagin
// @contact.email	artyomkorchagin333@gmail.com
// @host			localhost:3000
// @BasePath		/

var validate *validator.Validate

func init() {

	validate = validator.New(validator.WithRequiredStructEnabled())

	// ✅ REGISTER CUSTOM VALIDATOR IMMEDIATELY
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("month_year", func(fl validator.FieldLevel) bool {
			s := fl.Field().String()
			if s == "" {
				return true
			}
			_, err := helpers.ParseTime(s)
			return err == nil
		})
	}
}

func main() {

	cfg, err := config.LoadConfig(validate)
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

	zapLogger.Info("Starting application", zap.Uint("port", cfg.Server.Port))

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

	application := app.New(cfg, db, zapLogger, validate)
	if err := application.Run(); err != nil {
		zapLogger.Fatal("application error", zap.Error(err))
	}
}
