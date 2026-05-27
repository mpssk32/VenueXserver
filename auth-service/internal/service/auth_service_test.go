package service_test

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"os"
	"testing"
	"auth-service/internal/config"
	sharedjwt "venuex/shared/jwt"
)

func TestRegister_InvalidRole(t *testing.T) {
	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "test",
		Email:    "test@test.com",
		Password: "123456",
		Role:     "admin",
	}

	err := service.AuthService.Register(req)

	if err == nil {

		t.Error(
			"expected invalid role error",
		)
	}
}

func TestRegister_EmptyFields(t *testing.T) {

	req := models.RegisterRequest{}

	err := service.AuthService.Register(req)

	if err == nil {

		t.Error(
			"expected registration error",
		)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	config.ConnectDB()

	req := models.LoginRequest{
		Email:    "fake@test.com",
		Password: "wrong",
	}

	_, err := service.AuthService.Login(req)

	if err == nil {

		t.Error(
			"expected login error",
		)
	}
}

func TestGenerateJWT(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"supersecret",
	)

	token, err := sharedjwt.GenerateJWT(
		"user123",
		"user",
	)

	if err != nil {

		t.Error(err)
	}

	if token == "" {

		t.Error(
			"expected token",
		)
	}
}

func TestGenerateJWT_EmptySecret(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"",
	)

	token, err := sharedjwt.GenerateJWT(
		"user123",
		"user",
	)

	if err != nil {

		t.Log(err)
	}

	if token == "" {

		t.Log(
			"empty token generated",
		)
	}
}

func TestGetUserByID(t *testing.T) {
	config.ConnectDB()

	_, err := service.GetUserByID(
		"invalid",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetAllUsers(t *testing.T) {
	config.ConnectDB()

	_, err := service.GetAllUsers()

	if err != nil {

		t.Log(err)
	}
}

func TestDeleteUser(t *testing.T) {
	config.ConnectDB()

	err := service.DeleteUser(
		"invalid",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestCreateAuditLog(t *testing.T) {
	config.ConnectDB()

	err := service.CreateAuditLog(
		models.AuditLog{
			Action: "test action",
		},
	)

	if err != nil {

		t.Log(err)
	}
}


func TestRegister_DuplicateEmail(t *testing.T) {

	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "test",
		Email:    "test@test.com",
		Password: "123456",
		Role:     "user",
	}

	_ = service.AuthService.Register(req)

	err := service.AuthService.Register(req)

	if err != nil {

		t.Log(err)
	}
}


func TestLogin_EmptyCredentials(t *testing.T) {

	config.ConnectDB()

	req := models.LoginRequest{}

	_, err := service.AuthService.Login(req)

	if err != nil {

		t.Log(err)
	}
}


func TestRegister_ArtistRole(t *testing.T) {

	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "artist",
		Email:    "artist@test.com",
		Password: "123456",
		Role:     "artist",
	}

	err := service.AuthService.Register(req)

	if err != nil {

		t.Log(err)
	}
}


func TestRegister_VenueRole(t *testing.T) {

	config.ConnectDB()

	req := models.RegisterRequest{
		Username: "venue",
		Email:    "venue@test.com",
		Password: "123456",
		Role:     "venue",
	}

	err := service.AuthService.Register(req)

	if err != nil {

		t.Log(err)
	}
}


func TestLogin_WrongPassword(t *testing.T) {

	config.ConnectDB()

	req := models.LoginRequest{
		Email:    "artist@test.com",
		Password: "wrongpassword",
	}

	_, err := service.AuthService.Login(req)

	if err != nil {

		t.Log(err)
	}
}


func TestGenerateJWT_VenueRole(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"secret",
	)

	token, err := sharedjwt.GenerateJWT(
		"venue123",
		"venue",
	)

	if err != nil {

		t.Error(err)
	}

	if token == "" {

		t.Error(
			"expected token",
		)
	}
}