package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPError struct {
	Code int
	Err  error
}

func (e HTTPError) Error() string {
	return e.Err.Error()
}

type handlerFunc func(c *gin.Context) error

func (h *Handler) wrap(fn handlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		if err != nil {
			if httpErr, ok := err.(HTTPError); ok {
				h.logger.Println(httpErr.Code, err)
				c.JSON(httpErr.Code, gin.H{"error": httpErr.Err.Error()})
			} else {
				h.logger.Println(http.StatusInternalServerError, err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	}
}
