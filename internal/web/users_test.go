package web

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/yorodm/devgo/internal/engine"
)

var (
	e    *engine.Engine
	mock sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
}

func TestUsers(t *testing.T) {

	db, mock, err := sqlmock.New()
	defer func() {
		if err != nil {
			t.Error(err)
		}
	}()
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectCommit()
	e = engine.New(db, "123456")
}
