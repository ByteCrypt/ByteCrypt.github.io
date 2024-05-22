package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/v1/models"
	"context"
	"fmt"
	"strings"
)

func (provider *Provider) AddSubscription(sub models.Subscription) (models.Subscription, error) {
	emailValidation := provider.ValidateEmail(sub.Email)
	if emailValidation != nil {
		return models.NewSubscription("", ""), emailValidation
	}

	_, err := provider.GetSubscriptionByEmail(sub.Email)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return models.NewSubscription("", ""), fmt.Errorf("user already exists: %s", sub.Email)
	}

	addSub := database.AddSubscriptionParams{Email: sub.Email, Name: sub.Name}
	_, err = provider.Queries.AddSubscription(context.Background(), addSub)
	if err != nil {
		return models.NewSubscription("", ""), err
	}

	return sub, nil
}

func (provider *Provider) GetSubscriptionById(id int64) (models.Subscription, error) {
	sub, err := provider.Queries.GetSubscriptionById(context.Background(), id)
	if err != nil {
		return models.Subscription{}, err
	}

	return models.NewSubscription(sub.Email, sub.Name), nil
}

func (provider *Provider) GetSubscriptionByEmail(email string) (models.Subscription, error) {
	sub, err := provider.Queries.GetSubscriptionByEmail(context.Background(), email)
	if err != nil {
		return models.NewSubscription("", ""), err
	}

	return models.NewSubscription(sub.Email, sub.Name), nil
}

func (provider *Provider) RemoveSubscription(email string) error {
	emailValidation := provider.ValidateEmail(email)
	if emailValidation != nil {
		return emailValidation
	}

	err := provider.Queries.DeleteSubscriptionEmail(context.Background(), email)
	if err != nil {
		return err
	}

	return nil
}
