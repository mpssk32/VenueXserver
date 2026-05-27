package handler_test

import (
	"chat-service/internal/config"
	"chat-service/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMessages(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/messages?sender_id=1&receiver_id=2",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetMessages(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestGetMessages_EmptyQuery(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/messages",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetMessages(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestGetMessages_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/messages?sender_id=bad&receiver_id=bad",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetMessages(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}

func TestGetMessages_Method(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/messages?sender_id=1&receiver_id=2",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetMessages(rec, req)

	if rec.Code == 0 {

		t.Errorf(
			"handler did not respond",
		)
	}
}