package service_test

import (
	"auth-service/internal/config"
	"auth-service/internal/service"
	"testing"
)


func TestDeleteUser_EmptyID(t *testing.T) {

	config.ConnectDB()

	err := service.DeleteUser("")

	if err != nil {

		t.Log(err)
	}
}