package router

import (
	"net/http"
	"strconv"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createSubscription(c *gin.Context) error {

	var scr types.SubscriptionCreateRequest

	if err := c.BindJSON(&scr); err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	if err := h.subscriptionService.CreateSubscription(c, &scr); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return nil
}

func (h *Handler) getSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 64)
	if err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	sub, err := h.subscriptionService.GetSubscription(c, id)
	if err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	c.JSON(http.StatusOK, sub)
	return nil
}

func (h *Handler) updateSubscription(c *gin.Context) error {

	var sub types.Subscription

	if err := c.BindJSON(&sub); err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	if err := h.subscriptionService.UpdateSubscription(c, &sub); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return nil
}

func (h *Handler) deleteSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 64)
	if err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	if err := h.subscriptionService.DeleteSubscription(c, id); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return nil
}

func (h *Handler) getAllSubscriptions(c *gin.Context) error {
	subs, err := h.subscriptionService.GetAllSubscriptions(c)
	if err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	c.JSON(http.StatusOK, subs)
	return nil
}
