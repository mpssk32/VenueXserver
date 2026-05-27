package service_test

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {

	mockRepo := new(MockRepository)

	svc := service.NewAuthService(
		mockRepo,
	)

	mockRepo.
		On(
			"GetUserByEmail",
			"test@test.com",
		).
		Return(models.User{}, errors.New("not found"))

	mockRepo.
		On(
			"CreateUser",
			mock.Anything,
		).
		Return(nil)

	err := svc.Register(
		models.RegisterRequest{
			Username: "test",
			Email:    "test@test.com",
			Password: "123456",
			Role:     "user",
		},
	)

	if err != nil {

		t.Errorf(
			"unexpected error %v",
			err,
		)
	}

	mockRepo.AssertExpectations(t)
}