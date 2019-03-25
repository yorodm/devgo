package web

import (
	"encoding/json"
	"net/http"
)

type createUserInput struct {
	email    string `json:"email"`
	username string `json:"name"`
	password string `json:"password"`
}

func (s *service) createUser(w http.ResponseWriter, r *http.Request) {
	var dto createUserInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err := s.CreateUser(r.Context(), dto.email, dto.username, dto.password)
	if err != nil {
		serverError(w, err)
	}
}

func (s *service) users(w http.ResponseWriter, r *http.Request) {
	data, err := s.Users(r.Context())
	if err != nil {
		serverError(w, err)
	}
	jsonResponse(w, data, http.StatusOK)
}
