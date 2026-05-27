package service_test

import (
	"auth-service/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(
	user models.User,
) error {

	args := m.Called(user)

	return args.Error(0)
}

func (m *MockRepository) GetUserByEmail(
	email string,
) (models.User, error) {

	args := m.Called(email)

	return args.Get(0).(models.User),
		args.Error(1)
}

func (m *MockRepository) GetUserByID(
	id string,
) (models.User, error) {

	args := m.Called(id)

	return args.Get(0).(models.User),
		args.Error(1)
}

func (m *MockRepository) GetAllUsers() (
	[]models.User,
	error,
) {

	args := m.Called()

	return args.Get(0).([]models.User),
		args.Error(1)
}

func (m *MockRepository) DeleteUser(
	userID string,
) error {

	args := m.Called(userID)

	return args.Error(0)
}

func (m *MockRepository) CreateAuditLog(
	log models.AuditLog,
) error {

	args := m.Called(log)

	return args.Error(0)
}