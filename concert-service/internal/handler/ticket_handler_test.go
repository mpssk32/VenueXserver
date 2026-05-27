package handler_test

import (
	"bytes"
	"concert-service/internal/config"
	"concert-service/internal/handler"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTicket_BadRequest(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/tickets",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateTicket(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestCreateTicket_EmptyFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"event_id":""
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tickets",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateTicket(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestGetUserTickets_EmptyUserID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/tickets",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetUserTickets(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestGetUserTickets(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/tickets?user_id=550e8400-e29b-41d4-a716-446655440000",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetUserTickets(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}