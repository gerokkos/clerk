package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gerokkos/clerk/api/models"
)

func (s *Server) ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.fail(w, "oops something evil has happened", 500)
		return
	}
	w.Write(resp)
}

func (s *Server) fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(resp)
}

func (s *Server) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestShouldGetPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	app := &Server{}
	req, err := http.NewRequest("GET", "/clerks?limit=1", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("^SELECT * FROM clerk$").WillReturnRows(rows)

	// now we execute our request
	app.Clerks(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := struct {
		User []*models.User
	}{
		[]*models.User{
			{ID: 11454, Name: models.Name{First: "Kara", Last: "Bootsman"}, Email: "kara.bootsman@example.com", Cell: "(179)-338-5378", Picture: models.Picture{Medium: "https://randomuser.me/api/portraits/med/women/65.jpg"}, Registered: models.Registered{Date: time.Date(2019, time.September, 25, 21, 34, 11, 424000000, time.UTC)}},
		}}
	app.assertJSON(w.Body.Bytes(), data.User, t)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
