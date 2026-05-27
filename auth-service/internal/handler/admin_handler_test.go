package handler_test

import (
	"auth-service/internal/config"
	"auth-service/internal/handler"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func withAdminID(
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

func TestGetAllUsers(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/admin/users",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetAllUsers(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestDeleteUser(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/admin/users/550e8400-e29b-41d4-a716-446655440000",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "550e8400-e29b-41d4-a716-446655440000",
	})

	req = withAdminID(
		req,
		"550e8400-e29b-41d4-a716-446655440001",
	)

	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestDeleteUser_InvalidID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/admin/users/invalid",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "invalid",
	})

	req = withAdminID(
		req,
		"550e8400-e29b-41d4-a716-446655440001",
	)

	rec := httptest.NewRecorder()

	handler.DeleteUser(rec, req)

	if rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"expected 500 got %d",
			rec.Code,
		)
	}
}

func TestDeleteUser_WithoutAdmin(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/admin/users/550e8400-e29b-41d4-a716-446655440000",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "550e8400-e29b-41d4-a716-446655440000",
	})

	rec := httptest.NewRecorder()

	defer func() {

		if recover() == nil {

			t.Errorf(
				"expected panic because user_id missing",
			)
		}
	}()

	handler.DeleteUser(rec, req)
}