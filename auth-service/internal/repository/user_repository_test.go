package repository_test

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"testing"
)

func TestCreateUser(t *testing.T) {

	config.ConnectDB()

	user := models.User{
		Username: "repo_user",
		Email:    "repo@test.com",
		Password: "123456",
		Role:     "user",
	}

	err := repository.CreateUser(user)

	if err != nil {

		t.Log(err)
	}
}

func TestCreateUser_Empty(t *testing.T) {

	config.ConnectDB()

	user := models.User{}

	err := repository.CreateUser(user)

	if err != nil {

		t.Log(err)
	}
}

func TestGetUserByEmail(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetUserByEmail(
		"repo@test.com",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetUserByEmail_Invalid(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetUserByEmail(
		"fake@test.com",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetUserByID(t *testing.T) {

	config.ConnectDB()

	_, err := repository.GetUserByID(
		"invalid",
	)

	if err != nil {

		t.Log(err)
	}
}

func TestGetAllUsers(t *testing.T) {

	config.ConnectDB()

	users, err := repository.GetAllUsers()

	if err != nil {

		t.Log(err)
	}

	if users == nil {

		t.Log(
			"users slice is nil",
		)
	}
}

func TestDeleteUser(t *testing.T) {

	config.ConnectDB()

	err := repository.DeleteUser(
		"invalid",
	)

	if err != nil {

		t.Log(err)
	}
}