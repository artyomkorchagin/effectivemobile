package router

import (
	"net/http"
	"strconv"

	"github.com/artyomkorchagin/effectivemobile/internal/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	var scrj types.SubscriptionCreateRequestJSON

	if err := c.BindJSON(&scrj); err != nil {
		return types.ErrBadRequest(err)
	}

	scr, err := types.NewSubscriptionCreateRequest(scrj)
	if err != nil {
		return types.ErrBadRequest(err)
	}

	if err := h.subscriptionService.CreateSubscription(c, scr); err != nil {
		return err // обработается с помощью wrap(handlerFunc) gin.HandlerFunc
	}

	h.logger.Info("Successfully created subscription", zap.Any("subscription create request", scr))
	c.JSON(http.StatusOK, nil)
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
// @Failure      404  {object}  HTTPError "Not found"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/{id} [get]
func (h *Handler) getSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 64)

	if err != nil {
		return types.ErrBadRequest(err)
	}

	sub, err := h.subscriptionService.GetSubscription(c, id)
	if err != nil {
		return err
	}

	h.logger.Info("Successfuly got subcription", zap.Any("uuid", id))

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
// @Failure      404  {object}  HTTPError "Not found"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions [patch]
func (h *Handler) updateSubscription(c *gin.Context) error {
	var surj types.SubscriptionUpdateRequestJSON

	if err := c.BindJSON(&surj); err != nil {
		return types.ErrBadRequest(err)
	}

	sur, err := types.NewSubscriptionUpdateRequest(surj)
	if err != nil {
		return types.ErrBadRequest(err)
	}

	if err := h.subscriptionService.UpdateSubscription(c, sur); err != nil {
		return err
	}

	h.logger.Info("Updated subscription successfully", zap.Any("Subcription update request", sur))
	c.JSON(http.StatusOK, nil)
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
// @Failure      404  {object}  HTTPError "Not found"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/{id} [delete]
func (h *Handler) deleteSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 64)
	if err != nil {
		return types.ErrBadRequest(err)
	}

	if err := h.subscriptionService.DeleteSubscription(c, id); err != nil {
		return err
	}
	h.logger.Info("Deleted subscription successfully: ", zap.Uint64("id", id))
	c.JSON(http.StatusOK, nil)
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
		return err
	}
	h.logger.Info("Got all subscriptions successfully", zap.Any("subscriptions", subs))
	c.JSON(http.StatusOK, subs)
	return nil
}

// GetSumOfSubscriptions godoc
// @Summary      Get total sum of subscriptions
// @Description  Calculate the total revenue from subscriptions matching the filter
// @Tags         subscription
// @Produce      json
// @Param        user_id     query    string  false  "User UUID"
// @Param        service_name  query  string  false  "Service Name"
// @Param        start_date  query  string  false  "Start Date (format: MM-YYYY)"
// @Param        end_date    query  string  false  "End Date (format: MM-YYYY)"
// @Success      200  {object}  int "Total sum"
// @Failure      400  {object}  HTTPError "Bad request"
// @Failure      500  {object}  HTTPError "Internal server error"
// @Router       /subscriptions/sum [get]
func (h *Handler) getSumOfSubscriptions(c *gin.Context) error {
	filter := types.Filter{}

	if err := c.Bind(&filter); err != nil {
		return types.ErrBadRequest(err)
	}

	sum, err := h.subscriptionService.GetSumOfSubscriptions(c, filter)
	if err != nil {
		return err
	}
	h.logger.Info("Got sum of subscriptions successfully", zap.Uint("sum", sum))
	c.JSON(http.StatusOK, gin.H{"sum": sum})
	return nil
}
