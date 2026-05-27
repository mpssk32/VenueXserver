package service_test

import (
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/service"
	"testing"
)


func TestCreateAuditLog_EmptyAction(t *testing.T) {

	config.ConnectDB()

	err := service.CreateAuditLog(
		models.AuditLog{},
	)

	if err != nil {

		t.Log(err)
	}
}