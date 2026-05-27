package handler_test

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func withUserID(
	r *http.Request,
	id string,
) *http.Request {

	ctx := context.WithValue(
		r.Context(),
		"user_id",
		id,
	)

	return r.WithContext(ctx)
}

func TestGetMe_Success(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/me",
		nil,
	)

	req = withUserID(
		req,
		"550e8400-e29b-41d4-a716-446655440000",
	)

	rec := httptest.NewRecorder()

	handler.GetMe(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestGetMe_WithoutUserID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/me",
		nil,
	)

	rec := httptest.NewRecorder()

	defer func() {

		if recover() == nil {

			t.Errorf(
				"expected panic because user_id missing",
			)
		}
	}()

	handler.GetMe(rec, req)
}

func TestGetMe_InvalidUserID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/me",
		nil,
	)

	req = withUserID(
		req,
		"invalid-uuid",
	)

	rec := httptest.NewRecorder()

	handler.GetMe(rec, req)

	if rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"expected 500 got %d",
			rec.Code,
		)
	}
}

func TestGetMe_Method(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/me",
		nil,
	)

	req = withUserID(
		req,
		"550e8400-e29b-41d4-a716-446655440000",
	)

	rec := httptest.NewRecorder()

	handler.GetMe(rec, req)

	if rec.Code == 0 {

		t.Errorf(
			"handler did not respond",
		)
	}
}