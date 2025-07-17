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

func NewHandler(subscriptionService *subscription.Service, logger *log.Logger) *Handler {
	return &Handler{
		subscriptionService: subscriptionService,
		logger:              logger,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.LoggerMiddleware(h.logger))

	main := router.Group("/")
	{
		// basic CRUDL routes
		main.GET("/subscription/:id", func(c *gin.Context) {})
		main.POST("/subscription", h.AddSubscription)
		main.PUT("/subscription/:id", func(c *gin.Context) {})
		main.DELETE("/subscription/:id", func(c *gin.Context) {})
		main.GET("/subscriptions", func(c *gin.Context) {})

		main.GET("/status", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	return router
}
