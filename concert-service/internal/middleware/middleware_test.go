package middleware_test

import (
	"concert-service/internal/middleware"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	sharedjwt "venuex/shared/jwt"
)

func TestJWTMiddleware_NoToken(t *testing.T) {

	handler := sharedjwt.JWTMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}
func TestRoleMiddleware_Forbidden(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.RoleMiddleware(
		"admin",
	)(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/admin",
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

func TestRoleMiddleware_Success(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.RoleMiddleware(
		"admin",
	)(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/admin",
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

func TestRoleMiddleware_UserInsteadOfVenue(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware.RoleMiddleware(
		"venue",
	)(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/venue",
		nil,
	)

	ctx := context.WithValue(
		req.Context(),
		"role",
		"artist",
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

func TestJWTMiddleware_NoHeader(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := sharedjwt.JWTMiddleware(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}

func TestJWTMiddleware_InvalidToken(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := sharedjwt.JWTMiddleware(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer invalidtoken",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}

func TestJWTMiddleware_InvalidBearer(t *testing.T) {

	next := http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.WriteHeader(http.StatusOK)
	})

	handler := sharedjwt.JWTMiddleware(next)

	req := httptest.NewRequest(
		http.MethodGet,
		"/protected",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"InvalidFormat",
	)

	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}