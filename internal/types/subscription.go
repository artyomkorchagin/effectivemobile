package types

import (
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint64     `json:"id" binding:"required"`
	ServiceName string     `json:"service_name" binding:"required" validate:"required"`
	Price       uint       `json:"price" binding:"required" validate:"required"`
	UserUUID    uuid.UUID  `json:"user_id" binding:"required" validate:"required,uuid"`
	StartDate   time.Time  `json:"start_date" binding:"required" validate:"month_year"`
	EndDate     *time.Time `json:"end_date" binding:"-" validate:"month_year"`
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
	ID          uint64 `validate:"required"`
	ServiceName string `validate:"min=5,max=30"`
	Price       uint
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
	ServiceName string `validate:"min=5,max=30"`
	Price       uint
	UserUUID    uuid.UUID  `validate:"required,uuid"`
	StartDate   time.Time  `validate:"required,month_year"`
	EndDate     *time.Time `validate:"month_year"`
}

func NewSubscriptionCreateRequest(scrj SubscriptionCreateRequestJSON) (*SubscriptionCreateRequest, error) {

	userid, err := uuid.Parse(scrj.UserUUID)
	if err != nil {
		return nil, err
	}

	start, err := helpers.ParseTime(scrj.StartDate)
	if err != nil {
		return nil, err
	}
	end, err := helpers.ParseTime(scrj.StartDate)
	if err != nil {
		return nil, err
	}

	return &SubscriptionCreateRequest{
		ServiceName: scrj.ServiceName,
		Price:       scrj.Price,
		UserUUID:    userid,
		StartDate:   start,
		EndDate:     &end,
	}, nil
}

func NewSubscriptionUpdateRequest(surj SubscriptionUpdateRequestJSON) (*SubscriptionUpdateRequest, error) {

	userid, err := uuid.Parse(surj.UserUUID)
	if err != nil {
		return nil, err
	}

	start, err := helpers.ParseTime(surj.StartDate)
	if err != nil {
		return nil, err
	}

	end, err := helpers.ParseTime(surj.StartDate)
	if err != nil {
		return nil, err
	}

	return &SubscriptionUpdateRequest{
		ServiceName: surj.ServiceName,
		Price:       surj.Price,
		UserUUID:    userid,
		StartDate:   &start,
		EndDate:     &end,
	}, nil
}
