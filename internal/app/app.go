package app

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/artyomkorchagin/effectivemobile/internal/config"
	"github.com/artyomkorchagin/effectivemobile/internal/router"
	servicesubscription "github.com/artyomkorchagin/effectivemobile/internal/services/subscription"
	psqlsubscription "github.com/artyomkorchagin/effectivemobile/internal/storage/postgresql"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

const shutdownTimeout = 10 * time.Second

type App struct {
	config *config.Config
	db     *sql.DB
	logger *zap.Logger
	server *http.Server
}

func New(
	cfg *config.Config,
	db *sql.DB,
	logger *zap.Logger,
	validate *validator.Validate,
) *App {

	subRepo := psqlsubscription.NewRepository(db)
	subSvc := servicesubscription.NewService(subRepo, logger, validate)
	handler := router.NewHandler(subSvc, logger)
	r := handler.InitRouter()

	srv := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: r,
	}

	return &App{
		config: cfg,
		db:     db,
		logger: logger,
		server: srv,
	}
}

func (a *App) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		a.logger.Info("Server starting", zap.String("addr", a.server.Addr))
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Error("Server failed", zap.Error(err))
		}
	}()

	<-quit
	a.logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("Server shutdown failed", zap.Error(err))
		return err
	}

	a.logger.Info("Server stopped gracefully")
	return nil
}

func RunMigrations(db *sql.DB, logger *zap.Logger) error {
	logger.Info("Running database migrations")
	if err := psqlsubscription.RunMigrations(db); err != nil {
		return err
	}
	logger.Info("Migrations completed successfully")
	return nil
}
