package handler_test

import (
	"concert-service/internal/config"
	"concert-service/internal/handler"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"context"

	"github.com/gorilla/mux"
)
// cd concert-service && go test ./tests/...

// проверяет, что при запросе всех событий возвращается статус 200
func TestGetAllEventsHandler(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.GetAllEvents(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}
// проверяет, что при поиске событий по названию возвращается статус 200
func TestSearchEventsHandler(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/search?title=Metallica",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.SearchEvents(rec, req)

	if rec.Code != http.StatusOK {

		t.Errorf(
			"expected 200 got %d",
			rec.Code,
		)
	}
}
// проверяет, что при поиске событий по названию возвращается статус 200
func TestGetEventByIDHandler(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/550e8400-e29b-41d4-a716-446655440002",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "550e8400-e29b-41d4-a716-446655440002",
	})

	rec := httptest.NewRecorder()

	handler.GetEventByID(rec, req)

	if rec.Code != http.StatusOK &&
		rec.Code != http.StatusNotFound {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

// проверяет, что при попытке создать событие с невалидными данными возвращается статус 400
func TestCreateEventHandler_BadRequest(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateEvent(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}
// проверяет, что при поиске событий с пустым названием возвращается статус 400
func TestSearchEventsHandler_EmptyQuery(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/search",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.SearchEvents(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

// проверяет, что при поиске событий с невалидным ID возвращается статус 404 или 500
func TestGetEventByIDHandler_InvalidID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/invalid-id",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "invalid-id",
	})

	rec := httptest.NewRecorder()

	handler.GetEventByID(rec, req)

	if rec.Code != http.StatusNotFound &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}

// проверяет, что при попытке обновить событие другим пользователем возвращается статус 403 или 404
func TestUpdateEvent_Forbidden(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"title":"Updated Concert"
	}`)

	req := httptest.NewRequest(
		http.MethodPut,
		"/events/test",
		bytes.NewBuffer(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		"another-user",
	)

	req = req.WithContext(ctx)

	req = mux.SetURLVars(req, map[string]string{
		"id": "550e8400-e29b-41d4-a716-446655440002",
	})

	rec := httptest.NewRecorder()

	handler.UpdateEvent(rec, req)

	if rec.Code != http.StatusForbidden &&
		rec.Code != http.StatusNotFound {

		t.Errorf(
			"unexpected status code %d",
			rec.Code,
		)
	}
}


// проверяет, что при попытке создать событие с невалидными данными возвращается статус 400
func TestCreateEventHandler_InvalidData(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.CreateEvent(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

// проверяет, что при поиске событий с пустым названием возвращается статус 400
func TestSearchEventsHandler_EmptyTitle(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/search",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.SearchEvents(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}


// проверяет, что при поиске событий с невалидным ID возвращается статус 404 или 500
func TestGetEventByIDHandler_InvalidUUID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodGet,
		"/events/invalid",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "invalid",
	})

	rec := httptest.NewRecorder()

	handler.GetEventByID(rec, req)

	if rec.Code != http.StatusNotFound &&
		rec.Code != http.StatusInternalServerError {

		t.Errorf(
			"unexpected status %d",
			rec.Code,
		)
	}
}

//проверяет на пустые поля при создании события, что возвращается статус 400
func TestCreateEventHandler_EmptyFields(t *testing.T) {

	config.ConnectDB()

	body := []byte(`{
		"title":"",
		"venue_id":"",
		"artist_id":""
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/events",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	handler.CreateEvent(rec, req)

	if rec.Code != http.StatusBadRequest {

		t.Errorf(
			"expected 400 got %d",
			rec.Code,
		)
	}
}

// проверяет, что при попытке обновить событие с невалидными данными возвращается статус 400 или 404
func TestUpdateEvent_InvalidBody(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodPut,
		"/events/1",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	rec := httptest.NewRecorder()

	handler.UpdateEvent(rec, req)

	if rec.Code != http.StatusNotFound &&
		rec.Code != http.StatusBadRequest {

		t.Errorf(
			"unexpected status %d",
			rec.Code,
		)
	}
}


// проверяет, что при попытке обновить событие с невалидными данными возвращается статус 400 или 404
func TestDeleteEvent_InvalidID(t *testing.T) {

	config.ConnectDB()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/events/invalid",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{
		"id": "invalid",
	})

	rec := httptest.NewRecorder()

	handler.DeleteEvent(rec, req)

	if rec.Code != http.StatusNotFound &&
		rec.Code != http.StatusForbidden {

		t.Errorf(
			"unexpected status %d",
			rec.Code,
		)
	}
}


