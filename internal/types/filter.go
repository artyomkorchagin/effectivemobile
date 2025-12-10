package types

import (
	"time"

	"github.com/google/uuid"
)

type Filter struct {
	UserUUID    *uuid.UUID `form:"user_id" json:"user_id" validate:"uuid"`
	ServiceName *string    `form:"service_name" json:"service_name"`
	StartDate   *time.Time `form:"start_date" json:"start_date"`
	EndDate     *time.Time `form:"end_date" json:"end_date"`
}
