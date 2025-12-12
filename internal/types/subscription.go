package types

import (
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

// SubscriptionJSON godoc
// @Description All date fields are in MM-YYYY format. Price is in cents.
type SubscriptionJSON struct {
	// Unique numeric ID of the subscription
	// @example 123
	ID int64 `json:"id"`

	// Name of the service (e.g., "Spotify", "Yandex Plus")
	// @example "Yandex"
	ServiceName string `json:"service_name"`

	// Price in cents (e.g., 999 = $9.99)
	// @example 999
	Price int64 `json:"price"`

	// User UUID associated with this subscription
	// @example "550e8400-e29b-41d4-a716-446655440000"
	UserUUID uuid.UUID `json:"user_id"`

	// Start date in MM-YYYY format
	// @example "03-2025"
	StartDate string `json:"start_date"`

	// Optional end date in MM-YYYY format
	// @example "12-2025"
	EndDate string `json:"end_date,omitempty"`
}

type Subscription struct {
	ID          int64      `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int64      `json:"price"`
	UserUUID    uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}

// SubscriptionUpdateRequestJSON godoc
// @Description At least one field must be provided. Dates in MM-YYYY format. Price in cents.
type SubscriptionUpdateRequestJSON struct {
	// Subscription ID (required to identify the record)
	// @example 123
	ID int64 `json:"id" binding:"required"`

	// Optional new service name (5–30 characters)
	// @example "Yandex Music"
	ServiceName string `json:"service_name,omitempty"`

	// Optional new price in cents
	// @example 1299
	Price int64 `json:"price,omitempty"`

	// Optional new user UUID
	// @example "a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8"
	UserUUID string `json:"user_id,omitempty"`

	// Optional new start date (MM-YYYY)
	// @example "04-2025"
	StartDate string `json:"start_date,omitempty" validate:"month_year"`

	// Optional new end date (MM-YYYY)
	// @example "11-2026"
	EndDate string `json:"end_date,omitempty" validate:"month_year"`
}

type SubscriptionUpdateRequest struct {
	ID          int64      `validate:"required"`
	ServiceName string     `validate:"omitempty,min=5,max=30"`
	Price       int64      `validate:"omitempty,min=0"`
	UserUUID    uuid.UUID  `validate:"omitempty,uuid"`
	StartDate   *time.Time `validate:"omitempty"`
	EndDate     *time.Time `validate:"omitempty"`
}

// SubscriptionCreateRequestJSON godoc
// @Description All fields except end_date are required. Dates must be in MM-YYYY format. Price is in cents.
type SubscriptionCreateRequestJSON struct {
	// Service name (5–30 characters)
	// @example "Yandex"
	ServiceName string `json:"service_name" binding:"required"`

	// Price in cents (e.g., 999 = $9.99)
	// @example 999
	Price int64 `json:"price" binding:"required"`

	// User UUID (must be valid)
	// @example "550e8400-e29b-41d4-a716-446655440000"
	UserUUID string `json:"user_id" binding:"required"`

	// Start date in MM-YYYY format (e.g., "03-2025")
	// @example "03-2025"
	StartDate string `json:"start_date" binding:"required" validate:"month_year"`

	// Optional end date in MM-YYYY format
	// @example "12-2025"
	EndDate string `json:"end_date" validate:"omitempty,month_year"`
}
type SubscriptionCreateRequest struct {
	ServiceName string     `validate:"required,min=5,max=30"`
	Price       int64      `validate:"required,min=0"`
	UserUUID    uuid.UUID  `validate:"required,uuid"`
	StartDate   time.Time  `validate:"required"`
	EndDate     *time.Time `validate:"omitempty"`
}

func NewSubscriptionCreateRequest(scrj SubscriptionCreateRequestJSON) (*SubscriptionCreateRequest, error) {
	userID, err := uuid.Parse(scrj.UserUUID)
	if err != nil {
		return nil, err
	}

	startDate, err := helpers.ParseTime(scrj.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate *time.Time
	if scrj.EndDate != "" {
		end, err := helpers.ParseTime(scrj.EndDate)
		if err != nil {
			return nil, err
		}
		endDate = &end
	}

	return &SubscriptionCreateRequest{
		ServiceName: scrj.ServiceName,
		Price:       scrj.Price,
		UserUUID:    userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}

func NewSubscriptionUpdateRequest(surj SubscriptionUpdateRequestJSON) (*SubscriptionUpdateRequest, error) {

	var userID uuid.UUID
	if surj.UserUUID != "" {
		var err error
		userID, err = uuid.Parse(surj.UserUUID)
		if err != nil {
			return nil, err
		}
	}

	var startDate *time.Time
	if surj.StartDate != "" {
		start, err := helpers.ParseTime(surj.StartDate)
		if err != nil {
			return nil, err
		}
		startDate = &start
	}

	var endDate *time.Time
	if surj.EndDate != "" {
		end, err := helpers.ParseTime(surj.EndDate)
		if err != nil {
			return nil, err
		}
		endDate = &end
	}

	return &SubscriptionUpdateRequest{
		ID:          surj.ID,
		ServiceName: surj.ServiceName,
		Price:       surj.Price,
		UserUUID:    userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}

func NewSubscriptionJSON(sub Subscription) *SubscriptionJSON {
	var end string
	if sub.EndDate != nil {
		end = sub.EndDate.Format("01-2006")
	}

	return &SubscriptionJSON{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserUUID:    sub.UserUUID,
		StartDate:   sub.StartDate.Format("01-2006"),
		EndDate:     end,
	}
}
