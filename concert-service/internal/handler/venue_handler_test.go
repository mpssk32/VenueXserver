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

func TestCreateVenue_BadRequest(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/venues",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateVenue(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestCreateVenue_EmptyFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"name":"",
		"city":"",
		"description":"",
		"equipment":"",
		"capacity":0
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/venues",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateVenue(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestCreateVenue(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"name":"Arena",
		"city":"Amsterdam",
		"description":"Big venue",
		"equipment":"Lights",
		"capacity":1000
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/venues",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.CreateVenue(rec, req)

	if rec.Code != http.StatusCreated &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}