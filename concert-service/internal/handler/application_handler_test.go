package handler_test

import (
	"bytes"
	"concert-service/internal/config"
	"concert-service/internal/handler"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// cd concert-service && go test ././tests/...

// проверяет, что при создании заявки с пустыми полями возвращается статус 400
func TestCreateApplication_BadRequest(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/applications",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateApplication(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}


// проверяет, что при создании заявки с пустыми полями возвращается статус 400
func TestCreateApplication_EmptyFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"venue_id":"",
		"message":""
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/applications",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateApplication(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}


func TestGetVenueApplications(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/applications",
		nil,
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.GetVenueApplications(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}


func TestApproveApplication(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPut,
		"/applications/1/approve",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.ApproveApplication(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestRejectApplication(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPut,
		"/applications/1/reject",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.RejectApplication(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}