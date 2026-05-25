package tests

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/service"
	"testing"
)

func TestRegister_InvalidRole(t *testing.T) {

	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "test",
		Email:    "test@mail.com",
		Password: "123456",
		Role:     "super_admin",
	}

	err := service.Register(req)

	if err == nil {

		t.Error(
			"expected invalid role error",
		)
	}
}

func TestRegister_ValidRole(t *testing.T) {

	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "test",
		Email:    "valid@mail.com",
		Password: "123456",
		Role:     "user",
	}

	err := service.Register(req)

	if err != nil {

		t.Errorf(
			"expected successful registration, got %v",
			err,
		)
	}
}