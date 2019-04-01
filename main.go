package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yorodm/devgo/internal/engine"
	"github.com/yorodm/devgo/internal/web"
)

func main() {
	godotenv.Load()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	var (
		port      = environ("PORT", "3000")
		originStr = environ("SITE_URL", "http://localhost:"+port)
		dbURL     = environ("DATABASE_URL", "postgresql://root@127.0.0.1:26257/devgo?sslmode=disable")
		secret    = environ("SECRET_KEY", "1qazxsw23edcvfr45tgbgfr432we56789")
		blog      *engine.Engine
	)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("could not open db connection: %v", err)
	}
	blog = engine.New(db, secret)
	url, err := url.Parse(originStr)
	if err != nil {
		return fmt.Errorf("Invalid SITE_URL")
	}
	if err = web.StartEngine(blog, url); err != nil {
		return err
	}
	return nil
}

func environ(key, _default string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return _default

}
