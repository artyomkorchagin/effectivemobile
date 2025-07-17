package router

import (
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	subscriptionServices *
	logger   *log.Logger
}

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.LoggerMiddleware(h.logger))


	return router
}
