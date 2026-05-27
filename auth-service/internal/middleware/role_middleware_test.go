package middleware_test

import (
	"auth-service/internal/middleware"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoleMiddleware_Allowed(t *testing.T) {

	handler := middleware.RoleMiddleware(
		"admin",
	)(
		http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			w.WriteHeader(http.StatusOK)
		}),
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	ctx := context.WithValue(
		req.Context(),
		"role",
		"admin",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}

func TestRoleMiddleware_Forbidden(t *testing.T) {

	handler := middleware.RoleMiddleware(
		"admin",
	)(
		http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			w.WriteHeader(http.StatusOK)
		}),
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	ctx := context.WithValue(
		req.Context(),
		"role",
		"user",
	)

	req = req.WithContext(ctx)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {

		t.Errorf(
			"expected 403 got %d",
			rec.Code,
		)
	}
}