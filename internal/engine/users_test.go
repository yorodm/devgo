package engine

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	e  *Engine
	db *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/devgo_test?sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	e = New(db, "10293847")
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	ctx := context.TODO()
	err := e.CreateUser(ctx, "User", "user@user.com", "user", "123456")
	if err != nil {
		t.Error(err)
	}
}

func TestListUser(t *testing.T) {
	ctx := context.TODO()
	data, err := e.Users(ctx)
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
}
