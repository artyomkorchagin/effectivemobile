package types

import "github.com/google/uuid"

type Filter struct {
	UserUUID    *uuid.UUID `form:"user_id" json:"user_id" validate:"uuid"`
	ServiceName *string    `form:"service_name" json:"service_name"`
	StartDate   *string    `form:"start_date" json:"start_date"`
	EndDate     *string    `form:"end_date" json:"end_date"`
}
