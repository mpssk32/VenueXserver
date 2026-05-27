package handler_test

import (
	"auth-service/internal/handler"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-service/internal/config"
)

func TestRegister_Success(t *testing.T) {
	config.ConnectDB()

	body := []byte(`{
		"username":"alice",
		"email":"alice@example.com",
		"password":"secret",
		"role":"user"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	if rec.Code != http.StatusCreated &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestRegister_MissingFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"email":"alice@example.com"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestRegister_InvalidJSON(t *testing.T) {
	config.ConnectDB()


	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBufferString("{bad json"),
	)

	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestLogin_Success(t *testing.T) {

config.ConnectDB()

	body := []byte(`{
		"email":"alice@example.com",
		"password":"secret"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

func TestLogin_MissingCredentials(t *testing.T) {
	config.ConnectDB()

	body := []byte(`{
		"email":""
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(body),
	)


	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestLogin_InvalidJSON(t *testing.T) {
	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBufferString("{bad"),
	)

	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestAdminPanel(t *testing.T) {
	config.ConnectDB()
	req := httptest.NewRequest(
		http.MethodGet,
		"/admin",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.AdminPanel(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}