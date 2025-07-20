package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/gin-gonic/gin"
)

// CreateSubscription godoc
// @Summary      Create a subscription
// @Description  Create a new subscription
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        subscription  body      types.SubscriptionCreateRequest  true  "Create subscription"
// @Success      200  "No Content"
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions [post]
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
			Err:  fmt.Errorf("Error inserting subscription into DB: %v", err),
		}
	}
	return nil
}

// GetSubscription godoc
// @Summary      Get a subscription by ID
// @Description  Retrieve a subscription by its ID
// @Tags         subscription
// @Produce      json
// @Param        id   path    int     true  "Subscription ID"
// @Success      200  {object}  types.Subscription
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/{id} [get]
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

// UpdateSubscription godoc
// @Summary      Partially update a subscription
// @Description  Update only the provided fields of a subscription
// @Tags         subscription
// @Accept       json
// @Produce      json
// @Param        subscription body    types.SubscriptionUpdateRequest  true  "Fields to update"
// @Success      200  "No Content"
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions [patch]
func (h *Handler) updateSubscription(c *gin.Context) error {
	var sur types.SubscriptionUpdateRequest

	if err := c.BindJSON(&sur); err != nil {
		return HTTPError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("Error binding JSON to struct: %v", err),
		}
	}
	if err := h.subscriptionService.UpdateSubscription(c, &sur); err != nil {
		return HTTPError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("Error updating subscription : %v", err),
		}
	}
	return nil
}

// DeleteSubscription godoc
// @Summary      Delete a subscription
// @Description  Delete a subscription by ID
// @Tags         subscription
// @Produce      json
// @Param        id   path    int     true  "Subscription ID"
// @Success      200  "No Content"
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/{id} [delete]
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

// GetAllSubscriptions godoc
// @Summary      Get all subscriptions
// @Description  Retrieve a list of all subscriptions
// @Tags         subscription
// @Produce      json
// @Success      200  {array}  types.Subscription
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions [get]
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

// GetSumOfSubscriptions godoc
// @Summary      Get total sum of subscriptions
// @Description  Calculate the total revenue from subscriptions matching the filter
// @Tags         subscription
// @Produce      json
// @Param        filter body    types.Filter  true  "Fields to filter"
// @Success      200  {object}  int "Total sum"
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/sum [get]
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
