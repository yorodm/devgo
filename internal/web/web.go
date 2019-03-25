package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/csrf"
	"github.com/yorodm/devgo/internal/engine"
)

type service struct {
	*engine.Engine
}

// StartEngine starts the blog engine
func StartEngine(e *engine.Engine, url *url.URL) error {
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
	r.Use(csrf.Protect([]byte("32-byte-long-auth-key")))
	// Routes
	r.Get("/users", web.users)
	r.Post("/user", web.createUser)
	// Start server
	log.Println(url)
	server := http.Server{
		Addr:              ":" + url.Port(),
		Handler:           r,
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout:       time.Second * 15,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
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
