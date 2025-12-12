package types

import (
	"fmt"
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

type FilterQuery struct {
	UserUUID    string `form:"user_id" binding:"omitempty,uuid"`
	ServiceName string `form:"service_name" binding:"omitempty,min=5,max=30"`
	StartDate   string `form:"start_date" binding:"omitempty,month_year"`
	EndDate     string `form:"end_date" binding:"omitempty,month_year"`
}

type Filter struct {
	UserUUID    *uuid.UUID
	ServiceName *string
	StartDate   *time.Time
	EndDate     *time.Time
}

func NewFilter(fq FilterQuery) (*Filter, error) {
	var userUUID *uuid.UUID
	if fq.UserUUID != "" {
		id, err := uuid.Parse(fq.UserUUID)
		if err != nil {
			return nil, ErrBadRequest(fmt.Errorf("invalid user_id: %w", err))
		}
		userUUID = &id
	}

	var serviceName *string
	if fq.ServiceName != "" {
		serviceName = &fq.ServiceName
	}

	var startDate *time.Time
	if fq.StartDate != "" {
		t, err := helpers.ParseTime(fq.StartDate)
		if err != nil {
			return nil, ErrBadRequest(fmt.Errorf("invalid start_date: %w", err))
		}
		startDate = &t
	}

	var endDate *time.Time
	if fq.EndDate != "" {
		t, err := helpers.ParseTime(fq.EndDate)
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
