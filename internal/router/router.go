package router

import (
	"log"

	"github.com/artyomkorchagin/effectivemobile/internal/middleware"
	"github.com/artyomkorchagin/effectivemobile/internal/services/subscription"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	subscriptionService *subscription.Service
	logger              *log.Logger
}

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.LoggerMiddleware(h.logger))

	return router
}
