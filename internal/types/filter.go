package types

// @Description Filter represents the query parameters used to filter subscriptions.
// @Description Includes user ID, service name, and date range.
type Filter struct {
	// UserUUID is the unique identifier of the user.
	// Example: "a1b2c3d4-e5f6-7890-g1h2-i3j4k5l6m7n8"
	UserUUID string `form:"user_id" json:"user_id"`

	// ServiceName is the name of the service to filter by.
	// Example: "Yandex Plus"
	ServiceName string `form:"service_name" json:"service_name"`

	// StartDate is the start date of the filter range (in MM-YYY format).
	// Example: "01-2025"
	StartDate string `form:"start_date" json:"start_date"`

	// EndDate is the end date of the filter range (in MM-YYY format).
	// Example: "05-2025"
	EndDate string `form:"end_date" json:"end_date"`
}
