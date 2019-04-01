package web

import (
	"encoding/json"
	"net/http"
)

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *service) login(w http.ResponseWriter, r *http.Request) {
	var dto loginInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		serverError(w, err)
		return
	}
	err := s.Login(r.Context(), dto.Username, dto.Password)
	if err != nil {
		serverError(w, err)
		return
	}
}
