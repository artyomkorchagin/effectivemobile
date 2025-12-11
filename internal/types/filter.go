package types

import (
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

type FilterJSON struct {
	UserUUID    string `json:"user_id" binding:"omitempty,uuid"`
	ServiceName string `json:"service_name" binding:"omitempty,min=5,max=30"`
	StartDate   string `json:"start_date" binding:"omitempty,month_year"`
	EndDate     string `json:"end_date" binding:"omitempty,month_year"`
}

type Filter struct {
	UserUUID    *uuid.UUID
	ServiceName *string
	StartDate   *time.Time
	EndDate     *time.Time
}

func NewFilter(fj FilterJSON) (*Filter, error) {
	var userUUID *uuid.UUID
	if fj.UserUUID != "" {
		id, err := uuid.Parse(fj.UserUUID)
		if err != nil {
			return nil, ErrBadRequest(fmt.Errorf("invalid user_id: %w", err))
		}
		userUUID = &id
	}

	var serviceName *string
	if fj.ServiceName != "" {
		serviceName = &fj.ServiceName
	}

	var startDate *time.Time
	if fj.StartDate != "" {
		t, err := helpers.ParseTime(fj.StartDate)
		if err != nil {
			return nil, ErrBadRequest(fmt.Errorf("invalid start_date: %w", err))
		}
		startDate = &t
	}

	var endDate *time.Time
	if fj.EndDate != "" {
		t, err := helpers.ParseTime(fj.EndDate)
		if err != nil {
			return nil, ErrBadRequest(fmt.Errorf("invalid end_date: %w", err))
		}
		endDate = &t
	}

	return &Filter{
		UserUUID:    userUUID,
		ServiceName: serviceName,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil
}
