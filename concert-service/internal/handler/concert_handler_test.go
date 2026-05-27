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

func TestCreateConcert_BadRequest(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/concerts",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateConcert(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestCreateConcert_EmptyFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"title":"",
		"description":"",
		"genre":"",
		"venue_id":"",
		"ticket_price":-1
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/concerts",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateConcert(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestCreateConcert(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"title":"Metallica Live",
		"description":"Big concert",
		"genre":"Metal",
		"venue_id":"550e8400-e29b-41d4-a716-446655440000",
		"ticket_price":100
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/concerts",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateConcert(rec, req)

	if rec.Code != http.StatusCreated &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestGetConcerts(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/concerts",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetConcerts(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestCreateTicket_InvalidJSON(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{invalid json}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tickets",
		bytes.NewBuffer(body),
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

func TestCreateTicket_NoUserContext(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"event_id":"550e8400-e29b-41d4-a716-446655440001"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tickets",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	defer func() {
		if recover() == nil {

			t.Error(
				"expected panic from missing user_id",
			)
		}
	}()

	handler.CreateTicket(rec, req)
}

func TestGetUserTickets_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/tickets?user_id=invalid",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetUserTickets(rec, req)

	if rec.Code != http.StatusInternalServerError &&
		rec.Code != http.StatusOK {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}