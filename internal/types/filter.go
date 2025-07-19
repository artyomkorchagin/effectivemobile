package types

type Filter struct {
	UserUUID    string `form:"user_id" json:"user_id"`
	ServiceName string `form:"service_name" json:"service_name"`
	StartDate   string `form:"start_date" json:"start_date"`
	EndDate     string `form:"end_date" json:"end_date"`
}
