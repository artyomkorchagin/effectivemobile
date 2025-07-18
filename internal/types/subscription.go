package types

type Subscription struct {
	ID          uint64 `json:"id" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	UserUUID    string `json:"user_id" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"-"`
}

type SubscriptionCreateRequest struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	UserUUID    string `json:"user_id" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"-"`
}

func NewSubscriptionCreateRequest(serviceName, userUUID, startDate, endDate string, price uint) SubscriptionCreateRequest {
	return SubscriptionCreateRequest{
		ServiceName: serviceName,
		Price:       price,
		UserUUID:    userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
