package router

import (
	"log"
	"net/http"

	_ "github.com/artyomkorchagin/effectivemobile/docs"
	"github.com/artyomkorchagin/effectivemobile/internal/middleware"
	servicesubscription "github.com/artyomkorchagin/effectivemobile/internal/services/subscription"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		main.GET("/subscriptions/:id", h.wrap(h.getSubscription))
		main.POST("/subscriptions", h.wrap(h.createSubscription))
		main.PATCH("/subscriptions", h.wrap(h.updateSubscription))
		main.DELETE("/subscriptions/:id", h.wrap(h.deleteSubscription))
		main.GET("/subscriptions", h.wrap(h.getAllSubscriptions))
		main.GET("/subscriptions/sum", h.wrap(h.getSumOfSubscriptions))

		main.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
