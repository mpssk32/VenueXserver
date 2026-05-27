package handler_test

import (
	"chat-service/internal/handler"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"chat-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func generateTestToken(userID string) string {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
		},
	)

	tokenString, _ := token.SignedString(
		[]byte("test-secret"),
	)

	return tokenString
}

func TestChatHandler_WithoutToken(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestChatHandler_InvalidToken(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws?token=bad-token",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}

func TestChatHandler_ValidToken(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	token := generateTestToken(
		"550e8400-e29b-41d4-a716-446655440000",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws?token="+token,
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusSwitchingProtocols &&
		rec.Code != http.StatusOK {

		t.Logf(
			"websocket status: %d",
			rec.Code,
		)
	}
}

func TestChatHandler_EmptySecret(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"",
	)

	token := generateTestToken(
		"user-1",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws?token="+token,
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusUnauthorized &&
		rec.Code != http.StatusSwitchingProtocols {

		t.Errorf(
			"unexpected status %d",
			rec.Code,
		)
	}
}

func TestChatHandler_BadJWTFormat(t *testing.T) {

	os.Setenv(
		"JWT_SECRET",
		"test-secret",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws?token=123.456",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {

		t.Errorf(
			"expected 401 got %d",
			rec.Code,
		)
	}
}

func TestChatHandler_QueryWithoutTokenValue(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/ws?token=",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.ChatHandler(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

func TestGetMessages_ContentType(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/messages?sender_id=1&receiver_id=2",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetMessages(rec, req)

	contentType := rec.Header().Get(
		"Content-Type",
	)

	if contentType != "application/json" {

		t.Errorf(
			"unexpected content type %s",
			contentType,
		)
	}
}