package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddSubscription(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(jsonData)
}
