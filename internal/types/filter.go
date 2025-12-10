package types

import (
	"time"

	"github.com/artyomkorchagin/effectivemobile/pkg/helpers"
	"github.com/google/uuid"
)

type FilterJSON struct {
	UserUUID    string `json:"user_id" binding:"-"`
	ServiceName string `json:"service_name" binding:"-"`
	StartDate   string `json:"start_date" binding:"-"`
	EndDate     string `json:"end_date" binding:"-"`
}

type Filter struct {
	UserUUID    *uuid.UUID `validate:"uuid"`
	ServiceName *string    `validate:"min=5,max=30"`
	StartDate   *time.Time `validate:"month_year"`
	EndDate     *time.Time `validate:"month_year"`
}

func NewFilter(fj FilterJSON) (*Filter, error) {

	userid, err := uuid.Parse(fj.UserUUID)
	if err != nil {
		return nil, err
	}

	start, err := helpers.ParseTime(fj.StartDate)
	if err != nil {
		return nil, err
	}
	end, err := helpers.ParseTime(fj.StartDate)
	if err != nil {
		return nil, err
	}

	return &Filter{
		UserUUID:    &userid,
		ServiceName: &fj.ServiceName,
		StartDate:   &start,
		EndDate:     &end,
	}, nil
}
