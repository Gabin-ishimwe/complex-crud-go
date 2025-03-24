package user

import (
	"net/http"

	"github.com/gabin-ishimwe/complex-crud-go/types"
	"github.com/gabin-ishimwe/complex-crud-go/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{} // memory address to Handler
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {
		// return error
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}
