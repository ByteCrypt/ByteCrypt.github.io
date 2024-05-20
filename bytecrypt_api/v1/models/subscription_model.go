package models

type SubscriptionResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewSubscriptionResponse(ok bool, message string) SubscriptionResponse {
	return SubscriptionResponse{
		Ok:      ok,
		Message: message,
	}
}

type Subscription struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewSubscription(email string, name string) Subscription {
	return Subscription{
		Email: email,
		Name:  name,
	}
}
