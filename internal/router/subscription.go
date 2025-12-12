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
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body types.SubscriptionCreateRequestJSON  true  "Create subscription"
// @Success      201  "No Content"
// @Failure      400  {object}  object{error=string} "Bad request"
// @Failure      500  {object}  object{error=string} "Internal server error"
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
	c.JSON(http.StatusCreated, nil)
	return nil
}

// GetSubscription godoc
// @Summary Get subscription by ID
// @Description Fetches a single subscription by its numeric ID.
// @Tags subscriptions
// @Produce json
// @Param id path int64 true "Subscription ID"
// @Success 200 {object} types.SubscriptionJSON "Subscription details"
// @Failure 400 {object} object{error=string} "Invalid ID format"
// @Failure 404 {object} object{error=string} "Subscription not found"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /subscriptions/{id} [get]
func (h *Handler) getSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseInt(idRaw, 10, 64)

	if err != nil {
		return types.ErrBadRequest(err)
	}

	sub, err := h.subscriptionService.GetSubscription(c, id)
	if err != nil {
		return err
	}

	h.logger.Info("Successfuly got subscription", zap.Int64("id", id))
	subj := types.NewSubscriptionJSON(*sub)
	c.JSON(http.StatusOK, subj)
	return nil
}

// UpdateSubscription godoc
// @Summary Partially update a subscription
// @Description Updates only the provided fields of an existing subscription. All fields are optional.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body types.SubscriptionUpdateRequestJSON true "Fields to update (at least one required)"
// @Success 200 "No Content"
// @Failure 400 {object} object{error=string} "Invalid input"
// @Failure 404 {object} object{error=string} "Subscription not found"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /subscriptions [patch]
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

	h.logger.Info("Updated subscription successfully", zap.Any("Subscription update request", sur))
	c.JSON(http.StatusOK, nil)
	return nil
}

// DeleteSubscription godoc
// @Summary Delete a subscription
// @Description Permanently deletes a subscription by its ID.
// @Tags subscriptions
// @Produce json
// @Param id path int64 true "Subscription ID"
// @Success 204 "No Content"
// @Failure 400 {object} object{error=string} "Invalid ID format"
// @Failure 404 {object} object{error=string} "Subscription not found"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /subscriptions/{id} [delete]
func (h *Handler) deleteSubscription(c *gin.Context) error {
	idRaw := c.Param("id")
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		return types.ErrBadRequest(err)
	}

	if err := h.subscriptionService.DeleteSubscription(c, id); err != nil {
		return err
	}
	h.logger.Info("Deleted subscription successfully: ", zap.Int64("id", id))
	c.JSON(http.StatusNoContent, nil)
	return nil
}

// GetAllSubscriptions godoc
// GetAllSubscriptions retrieves all subscriptions.
// @Summary Get all subscriptions
// @Description Returns a list of all active and inactive subscriptions.
// @Tags subscriptions
// @Produce json
// @Success 200 {array} types.Subscription "List of subscriptions"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /subscriptions [get]
func (h *Handler) getAllSubscriptions(c *gin.Context) error {
	subs, err := h.subscriptionService.GetAllSubscriptions(c)
	if err != nil {
		return err
	}
	h.logger.Info("Got all subscriptions successfully", zap.Any("subscriptions", subs))

	subsj := make([]*types.SubscriptionJSON, len(subs))
	for i, sub := range subs {
		subsj[i] = types.NewSubscriptionJSON(*sub)
	}
	h.logger.Info("Transformed subscriptions to json", zap.Any("subscriptionsJSON", subsj))
	c.JSON(http.StatusOK, subsj)
	return nil
}

// GetSumOfSubscriptions godoc
// @Summary Get total sum of subscriptions
// @Description Calculates the total price (in cents) of subscriptions matching the optional filters.
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "Filter by user UUID" Format(uuid)
// @Param service_name query string false "Filter by service name (min 5, max 30 chars)"
// @Param start_date query string false "Filter by start date" Format(MM-YYYY)
// @Param end_date query string false "Filter by end date" Format(MM-YYYY)
// @Success 200 {object} object{sum=int} "Total sum in cents"
// @Failure 400 {object} object{error=string} "Invalid filter parameters"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /subscriptions/sum [get]
func (h *Handler) getSumOfSubscriptions(c *gin.Context) error {
	fj := types.FilterQuery{}

	if err := c.BindQuery(&fj); err != nil {
		return types.ErrBadRequest(err)
	}

	filter, err := types.NewFilter(fj)
	if err != nil {
		return types.ErrBadRequest(err)
	}
	h.logger.Info("Parsed this filter", zap.Any("filter", filter))
	sum, err := h.subscriptionService.GetSumOfSubscriptions(c, filter)
	if err != nil {
		return err
	}
	h.logger.Info("Got sum of subscriptions successfully", zap.Int64("sum", sum))
	c.JSON(http.StatusOK, gin.H{"sum": sum})
	return nil
}
