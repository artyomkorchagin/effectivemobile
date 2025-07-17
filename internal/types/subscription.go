package types

type Subscription struct {
	ID          string
	ServiceName string
	UserUUID    string
	StartDate   string
	EndDate     string
}

type SubscriptionCreateRequest struct {
	ServiceName string
	UserUUID    string
	StartDate   string
	EndDate     string
}

func NewSubscriptionCreateRequest(serviceName, userUUID, startDate, endDate string) SubscriptionCreateRequest {
	return SubscriptionCreateRequest{
		ServiceName: serviceName,
		UserUUID:    userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}

func (s SubscriptionCreateRequest) JSON() string {
	// jsonify request entity for request here
	return ""
}
