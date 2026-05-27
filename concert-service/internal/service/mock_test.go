package service_test

import (
	"concert-service/internal/models"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockConcertRepository struct {
	mock.Mock
}

func (m *MockConcertRepository) CreateConcert(
	concert models.Concert,
) error {

	args := m.Called(concert)

	return args.Error(0)
}

func TestMockCreateConcert(t *testing.T) {

	mockRepo := new(MockConcertRepository)

	concert := models.Concert{
		Title: "Mock Concert",
	}

	mockRepo.On(
		"CreateConcert",
		concert,
	).Return(nil)

	err := mockRepo.CreateConcert(concert)

	if err != nil {

		t.Error(err)
	}

	mockRepo.AssertExpectations(t)
}