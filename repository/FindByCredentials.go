package repository

import (
	"basic-jwt-auth/models"
	"errors"
)

func FindByCredentials(email, password string) (*models.User, error) {
	if email == "test@mail.com" && password == "test12345" {
		return &models.User{
			ID:             1,
			Email:          "test@mail.com",
			Password:       "test12345",
			FavoritePhrase: "Hello World",
		}, nil
	}
	return nil, errors.New("user not found")
}
