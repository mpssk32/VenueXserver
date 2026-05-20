package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
)

func CreateAuditLog(log models.AuditLog) error {
	return repository.CreateAuditLog(log)
}