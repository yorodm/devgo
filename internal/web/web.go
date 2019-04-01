package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/yorodm/devgo/internal/engine"
)

type service struct {
	*engine.Engine
}

// NewHandler creates a new http handler for the blog engine
func NewHandler(e *engine.Engine, url *url.URL) (http.Handler, error) {
	r := chi.NewRouter()
	web := &service{e}
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{url.String()},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Routes
	r.Get("/user", web.listUsers)
	r.Post("/user", web.createUser)
	return r, nil
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func jsonResponse(w http.ResponseWriter, v interface{}, status int) {
	data, err := json.Marshal(v)
	if err != nil {
		serverError(w, fmt.Errorf("error marshalling response %v", err))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
