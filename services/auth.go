package services

import (
	"errors"
	"qdjr-api/helpers"
	"qdjr-api/initializers"
	"qdjr-api/models"
)

type AuthService struct{}

var authHelper = new(helpers.AuthHelper)

func (service AuthService) Register(username string, email string, password string) (*models.User, error) {
	passwordHash, err := authHelper.HashPassword(password)
	if err != nil {
		panic(err)
	}
	user := &models.User{Username: username, Email: email, Password: passwordHash}
	result := initializers.DB.Create(user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (service AuthService) Login(username string, password string) (user *models.User, err error) {
	user = UserService{}.DetailByParam(username)
	if user.Id == 0 {
		err = errors.New("user not found")
		return user, err
	}
	err = authHelper.VerifyPassword(user.Password, password)
	if err != nil {
		return user, err
	}
	return user, nil
}
