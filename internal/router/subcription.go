package router

import (
	"fmt"
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
			Err:  fmt.Errorf("Error binding JSON to struct: %v", err),
		}
	}

	if err := h.subscriptionService.CreateSubscription(c, &scr); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error inserting subscription into DB: ", err),
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
			Err:  fmt.Errorf("Error parsing parameter id (idRaw: %v) as uint64: %v", idRaw, err),
		}
	}
	sub, err := h.subscriptionService.GetSubscription(c, id)
	if err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error retrieving subscription (id: %d) from DB: %v", id, err),
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
			Err:  fmt.Errorf("Error binding JSON to struct: ", err),
		}
	}
	if err := h.subscriptionService.UpdateSubscription(c, &sub); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error updating subscription : ", err),
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
			Err:  fmt.Errorf("Error parsing parameter id (idRaw: %v) as uint64: %v", idRaw, err),
		}
	}
	if err := h.subscriptionService.DeleteSubscription(c, id); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error deleting subscription (id: %d) from DB: %v", id, err),
		}
	}
	return nil
}

func (h *Handler) getAllSubscriptions(c *gin.Context) error {
	subs, err := h.subscriptionService.GetAllSubscriptions(c)
	if err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error retrieving all subscriptions from DB: %v", err),
		}
	}
	c.JSON(http.StatusOK, subs)
	return nil
}

func (h *Handler) getSumOfSubscriptions(c *gin.Context) error {
	filter := types.Filter{}
	if err := c.Bind(&filter); err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Error binding parameters to struct as JSON: %v", err),
		}
	}
	sum, err := h.subscriptionService.GetSumOfSubscriptions(c, &filter)
	if err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error retrieving sum from DB: %v", err),
		}
	}
	c.JSON(http.StatusOK, gin.H{"sum": sum})
	return nil
}
