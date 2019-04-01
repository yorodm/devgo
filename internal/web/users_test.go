package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/yorodm/devgo/internal/engine"
)

func setupServiceAndMock(t *testing.T) (s *service, mock sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fail()
	}
	e := engine.New(db, "123456")
	s = &service{e}
	return
}

func TestListUsers(t *testing.T) {
	var m []map[string]interface{}
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	request.Header.Add("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s, mock := setupServiceAndMock(t)
	rows := sqlmock.NewRows([]string{"id", "email", "username"}).
		AddRow(1, "adrian@nowhere.com", "adrian").
		AddRow(2, "hernan@nowhere.com", "hernan")
	mock.ExpectQuery("select id, email, username from users").WillReturnRows(rows)
	s.listUsers(response, request)
	r := response.Result()
	defer r.Body.Close()
	if r.StatusCode == 500 {
		t.Error(r)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		t.Error(err)
	}
	if m[0]["email"] != "adrian@nowhere.com" {
		t.Errorf("Wrong response %v", m)
	}

}

func TestCreateUsers(t *testing.T) {
	c := createUserInput{
		Name:     "Named User",
		Email:    "user@user.com",
		Username: "user",
		Password: "1029384765",
	}
	body, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	response := httptest.NewRecorder()
	s, mock := setupServiceAndMock(t)
	mock.ExpectExec("insert into users*").
		WithArgs(c.Name, c.Email, c.Username, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.createUser(response, request)
	if r := response.Result(); r.StatusCode == 500 {
		t.Error(r)
	} else {
		t.Log(r)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
