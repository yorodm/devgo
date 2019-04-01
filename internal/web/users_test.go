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
	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	request.Header.Add("Content-Type", "application/json")
	response := httptest.NewRecorder()
	s, mock := setupServiceAndMock(t)
	mock.ExpectQuery("select id, email, username from users").WillReturnRows(sqlmock.NewRows([]string{"1", "user@user.com", "user"}))
	s.listUsers(response, request)
	if r := response.Result(); r.StatusCode == 500 {
		t.Error(r)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
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
