package web

import (
	"encoding/json"
	"net/http"
)

type createUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *service) createUser(w http.ResponseWriter, r *http.Request) {
	var dto createUserInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		serverError(w, err)
		return
	}
	err := s.CreateUser(r.Context(), dto.Name, dto.Email, dto.Username, dto.Password)
	if err != nil {
		serverError(w, err)
		return
	}
}

func (s *service) listUsers(w http.ResponseWriter, r *http.Request) {
	data, err := s.ListUsers(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}
	jsonResponse(w, data, http.StatusOK)
}
