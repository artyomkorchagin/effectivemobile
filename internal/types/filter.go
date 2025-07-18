package types

type Filter struct {
	UserUUID    string `form:"user_id"`
	ServiceName string `form:"service_name"`
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
}
