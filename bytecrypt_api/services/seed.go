package services

import (
	"bytecrypt_api/v1/models"
	"fmt"
)

var seedList = []models.Subscription{
	{Email: "a@yahoo.com", Name: "a"},
	{Email: "b@yahoo.com", Name: "b"},
	{Email: "c@yahoo.com", Name: "c"},
	{Email: "d@yahoo.com", Name: "t"},
	{Email: "e@yahoo.com", Name: "e"},
	{Email: "f@yahoo.com", Name: "f"},
	{Email: "g@yahoo.com", Name: "g"},
	{Email: "h@yahoo.com", Name: "h"},
	{Email: "i@yahoo.com", Name: "i"},
	{Email: "j@yahoo.com", Name: "j"},
	{Email: "k@yahoo.com", Name: "k"},
	{Email: "l@yahoo.com", Name: "l"},
	{Email: "m@yahoo.com", Name: "m"},
	{Email: "n@yahoo.com", Name: "n"},
	{Email: "o@yahoo.com", Name: "o"},
	{Email: "p@yahoo.com", Name: "p"},
	{Email: "q@yahoo.com", Name: "q"},
	{Email: "r@yahoo.com", Name: "r"},
	{Email: "s@yahoo.com", Name: "s"},
	{Email: "t@yahoo.com", Name: "t"},
	{Email: "u@yahoo.com", Name: "u"},
	{Email: "v@yahoo.com", Name: "v"},
	{Email: "w@yahoo.com", Name: "w"},
	{Email: "x@yahoo.com", Name: "x"},
	{Email: "y@yahoo.com", Name: "y"},
	{Email: "z@yahoo.com", Name: "z"},
	{Email: "n@protonmail.com", Name: "n"},
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
