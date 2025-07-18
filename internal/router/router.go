package router

import (
	"log"
	"net/http"

	"github.com/artyomkorchagin/effectivemobile/internal/middleware"
	servicesubscription "github.com/artyomkorchagin/effectivemobile/internal/services/subscription"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	subscriptionService *servicesubscription.Service
	logger              *log.Logger
}

func NewHandler(subscriptionService *servicesubscription.Service, logger *log.Logger) *Handler {
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
		main.GET("/subscription/:id", h.wrap(h.getSubscription))
		main.POST("/subscription", h.wrap(h.createSubscription))
		main.PUT("/subscription", h.wrap(h.updateSubscription))
		main.DELETE("/subscription/:id", h.wrap(h.deleteSubscription))
		main.GET("/subscriptions", h.wrap(h.getAllSubscriptions))
		main.GET("/sum-of-subscriptions", h.wrap(h.getSumOfSubscriptions))

		main.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	return router
}
