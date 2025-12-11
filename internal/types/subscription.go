package types

import (
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint64     `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       uint       `json:"price"`
	UserUUID    uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}

type SubscriptionUpdateRequestJSON struct {
	ID          uint64 `json:"id" binding:"required"`
	ServiceName string `json:"service_name,omitempty"`
	Price       uint   `json:"price,omitempty"`
	UserUUID    string `json:"user_id,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
}

type SubscriptionUpdateRequest struct {
	ID          uint64     `validate:"required"`
	ServiceName string     `validate:"min=5,max=30"`
	Price       uint       `validate:"required,min=0"`
	UserUUID    uuid.UUID  `validate:"uuid"`
	StartDate   *time.Time `validate:"month_year"`
	EndDate     *time.Time `validate:"month_year"`
}

type SubscriptionCreateRequestJSON struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	UserUUID    string `json:"user_id" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"-"`
}

type SubscriptionCreateRequest struct {
	ServiceName string     `validate:"required,min=5,max=30"`
	Price       uint       `validate:"required,min=0"`
	UserUUID    uuid.UUID  `validate:"required,uuid"`
	StartDate   time.Time  `validate:"required,month_year"`
	EndDate     *time.Time `validate:"month_year"`
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
		ServiceName: surj.ServiceName,
		Price:       surj.Price,
		UserUUID:    userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
