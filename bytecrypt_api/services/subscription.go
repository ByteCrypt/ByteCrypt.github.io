package services

import (
	"bytecrypt_api/database"
	"bytecrypt_api/v1/models"
	"context"
	"fmt"
	"regexp"
	"strings"
)

const EMAIL_EXPRESSION = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

func (provider *Provider) ValidateEmail(email string) error {
	emailExp, err := regexp.Compile(EMAIL_EXPRESSION)
	if err != nil {
		return err
	}

	if !emailExp.MatchString(email) {
		return fmt.Errorf("email is invalid: %s", email)
	}

	return nil
}

func (provider *Provider) AddSubscription(sub models.Subscription) (models.Subscription, error) {
	emailValidation := provider.ValidateEmail(sub.Email)
	if emailValidation != nil {
		return models.Subscription{}, emailValidation
	}

	_, err := provider.GetSubscriptionByEmail(sub.Email)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return models.Subscription{}, fmt.Errorf("user already exists: %s", sub.Email)
	}

	addSub := database.AddSubscriptionParams{Email: sub.Email, Name: sub.Name}
	_, err = provider.Queries.AddSubscription(context.Background(), addSub)
	if err != nil {
		return models.Subscription{}, err
	}

	return sub, nil
}

func (provider *Provider) GetSubscriptionByEmail(email string) (models.Subscription, error) {
	sub, err := provider.Queries.GetSubscriptionEmail(context.Background(), email)
	if err != nil {
		return models.Subscription{}, err
	}

	return models.Subscription{Email: sub.Email, Name: sub.Name}, nil
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
