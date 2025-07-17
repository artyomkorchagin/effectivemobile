package types

type Subscription struct {
	ID           string
	Service_name string
	User_uuid    string
	Start_date   string
	End_date     string
}

type SubscriptionCreateRequest struct {
	Service_name string
	User_uuid    string
	Start_date   string
	End_date     string
}

func NewSubscription(s SubscriptionCreateRequest) Subscription {
	// generate uuid for subscription here
	uuid := "some uuid"
	return Subscription{
		ID:           uuid,
		Service_name: s.Service_name,
		User_uuid:    s.User_uuid,
		Start_date:   s.Start_date,
		End_date:     s.End_date,
	}
}

func (s Subscription) JSON() string {

}
