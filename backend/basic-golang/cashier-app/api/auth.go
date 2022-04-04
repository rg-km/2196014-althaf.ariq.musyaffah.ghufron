package api

import (
	"encoding/json"
	"net/http"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	LoginRequest := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&LoginRequest)
	if err != nil {
		encoder.Encode(AuthErrorResponse{Error: "Invalid request"})
		return
	}

	if LoginRequest.Username == "admin" && LoginRequest.Password == "admin" {
		encoder.Encode(LoginSuccessResponse{Username: LoginRequest.Username})
		return
	}

	encoder.Encode(AuthErrorResponse{Error: "Invalid credentials"})

}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this

}
