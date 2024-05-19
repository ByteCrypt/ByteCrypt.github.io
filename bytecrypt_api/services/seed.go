package services

import (
	"bytecrypt_api/v1/models"
	"fmt"
)

var seedList = []models.Subscription{
	{Email: "a@example.mail", Name: "a"},
	{Email: "b@example.mail", Name: "b"},
	{Email: "c@example.mail", Name: "c"},
	{Email: "d@example.mail", Name: "t"},
	{Email: "e@example.mail", Name: "e"},
	{Email: "f@example.mail", Name: "f"},
	{Email: "g@example.mail", Name: "g"},
	{Email: "h@example.mail", Name: "h"},
	{Email: "i@example.mail", Name: "i"},
	{Email: "j@example.mail", Name: "j"},
	{Email: "k@example.mail", Name: "k"},
	{Email: "l@example.mail", Name: "l"},
	{Email: "m@example.mail", Name: "m"},
	{Email: "n@example.mail", Name: "n"},
	{Email: "o@example.mail", Name: "o"},
	{Email: "p@example.mail", Name: "p"},
	{Email: "q@example.mail", Name: "q"},
	{Email: "r@example.mail", Name: "r"},
	{Email: "s@example.mail", Name: "s"},
	{Email: "t@example.mail", Name: "t"},
	{Email: "u@example.mail", Name: "u"},
	{Email: "v@example.mail", Name: "v"},
	{Email: "w@example.mail", Name: "w"},
	{Email: "x@example.mail", Name: "x"},
	{Email: "y@example.mail", Name: "y"},
	{Email: "z@example.mail", Name: "z"},
}

func (provider *Provider) SeedDatabase() error {
	var errs []error
	for _, sub := range seedList {
		_, err := provider.AddSubscription(sub)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		var e string
		for _, err := range errs {
			e += err.Error() + "\n"
		}
		return fmt.Errorf(e)
	}

	return nil
}

func (provider *Provider) DeleteDatabaseSeed() error {
	var errs []error
	for _, sub := range seedList {
		err := provider.RemoveSubscription(sub.Email)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		var e string
		for _, err := range errs {
			e += err.Error() + "\n"
		}
		return fmt.Errorf(e)
	}

	return nil
}

func (provider *Provider) GetAllSeeded() ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	var errs []error
	for _, sub := range seedList {
		s, err := provider.GetSubscriptionByEmail(sub.Email)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		subscriptions = append(subscriptions, s)
	}
	if len(errs) != 0 {
		var e string
		for _, err := range errs {
			e += err.Error() + "\n"
		}
		return nil, fmt.Errorf(e)
	}

	return subscriptions, nil
}
