package user

import (
	"fmt"
	"net/http"

	"github.com/gabin-ishimwe/complex-crud-go/service/auth"
	"github.com/gabin-ishimwe/complex-crud-go/types"
	"github.com/gabin-ishimwe/complex-crud-go/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store} // memory address to Handler
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Handle Login
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {
		// return error
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// Check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("User with email %s already exists", payload.Email))
		return
	}
	// Hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// if doesn't creater the new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, nil)

}
