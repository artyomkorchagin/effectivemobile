package types

type Subscription struct {
	ID          string
	ServiceName string
	Price       string
	UserUUID    string
	StartDate   string
	EndDate     string
}

type SubscriptionCreateRequest struct {
	ServiceName string
	Price       string
	UserUUID    string
	StartDate   string
	EndDate     string
}

func NewSubscriptionCreateRequest(serviceName, price, userUUID, startDate, endDate string) SubscriptionCreateRequest {
	return SubscriptionCreateRequest{
		ServiceName: serviceName,
		Price:       price,
		UserUUID:    userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}

func (s SubscriptionCreateRequest) JSON() string {
	// jsonify request entity for request here
	return ""
}
