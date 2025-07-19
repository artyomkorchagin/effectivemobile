package types

// Subscription represents a user's subscription to a service.
type Subscription struct {
	// Unique identifier for the subscription
	ID uint64 `json:"id" binding:"required"`

	// Name of the service
	ServiceName string `json:"service_name" binding:"required"`

	// Monthly price of the subscription in USD cents or integer units
	Price uint `json:"price" binding:"required"`

	// Unique identifier of the user who owns this subscription
	UserUUID string `json:"user_id" binding:"required"`

	// Start date of the subscription in "MM-YYYY" format
	StartDate string `json:"start_date" binding:"required"`

	// End date of the subscription in "MM-YYYY" format
	EndDate string `json:"end_date" binding:"-"`
}

// SubscriptionUpdateRequest represents a request to update a subscription (PATCH).
// Only the fields provided will be updated.
type SubscriptionUpdateRequest struct {
	// ID is the unique identifier of the subscription to update
	// Required: yes
	ID uint64 `json:"id" binding:"required"`

	// ServiceName is the new name of the service (optional)
	// Example: "Yandex Plus"
	ServiceName string `json:"service_name,omitempty"`

	// Price is the new monthly price in USD cents or integer units (optional)
	// Example: 1000
	Price uint `json:"price,omitempty"`

	// UserUUID is the new unique identifier of the user (optional)
	// Example: "123e4567-e89b-42d3-a456-556642440000"
	UserUUID string `json:"user_id,omitempty"`

	// StartDate is the new start date in "MM-YYYY" format (optional)
	// Example: "01-2025"
	StartDate string `json:"start_date,omitempty"`

	// EndDate is the new end date in "MM-YYYY" format (optional)
	// Example: "02-2025"
	EndDate string `json:"end_date,omitempty"`
}

// SubscriptionCreateRequest is used to create a new subscription.
type SubscriptionCreateRequest struct {
	// Name of the service
	ServiceName string `json:"service_name" binding:"required"`

	// Monthly price of the subscription in rubles
	Price uint `json:"price" binding:"required"`

	// Unique identifier of the user who owns this subscription (uuid)
	UserUUID string `json:"user_id" binding:"required"`

	// Start date of the subscription in "MM-YYYY" format
	StartDate string `json:"start_date" binding:"required"`

	// End date of the subscription in "MM-YYYY" format; optional
	EndDate string `json:"end_date" binding:"-"`
}

func NewSubscriptionCreateRequest(serviceName string, price uint, userUUID, startDate, endDate string) SubscriptionCreateRequest {
	return SubscriptionCreateRequest{
		ServiceName: serviceName,
		Price:       price,
		UserUUID:    userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
