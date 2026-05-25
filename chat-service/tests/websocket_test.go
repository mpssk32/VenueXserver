package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"chat-service/internal/handler"

	"github.com/gorilla/mux"
)

func TestWebSocketEndpoint(t *testing.T) {

	r := mux.NewRouter()

	r.HandleFunc(
		"/ws",
		handler.ChatHandler,
	)

	req := httptest.NewRequest(
		"GET",
		"/ws",
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusSwitchingProtocols &&
		w.Code != http.StatusBadRequest {

		t.Errorf(
			"unexpected status code: %d",
			w.Code,
		)
	}
}


// go test ./tests/...