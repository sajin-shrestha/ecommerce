package user

import (
	"encoding/json"
	"net/http"

	"github.com/sajin-shrestha/ecommerce/types"
	"github.com/sajin-shrestha/ecommerce/utils"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPayLoad
	if err := utils.ParseJSON(r.Body, payload); err != nil {
		
	}

	// check if the user exists

	// if it doesnot we create the new user

}