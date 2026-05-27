package routes_test

import (
	"concert-service/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"
	"concert-service/internal/config"

	"github.com/gorilla/mux"
)

func TestRoutes(t *testing.T) {
	
	config.ConnectDB()

	r := mux.NewRouter()

	routes.RegisterRoutes(r)

	req := httptest.NewRequest(
		http.MethodGet,
		"/events",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code == 404 {

		t.Error("route not registered")
	}
}