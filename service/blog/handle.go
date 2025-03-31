package blog

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BlogHandler struct {
	blogRepository BlogRepositoryInterface
}

func NewBlogHandler(repository BlogRepositoryInterface) *BlogHandler {
	return &BlogHandler{blogRepository: repository}
}

func (b *BlogHandler) RegisterBlogHandler(router *mux.Router) {
	router.HandleFunc("/blog", b.handleBlogCreation).Methods("GET")
}

func (b *BlogHandler) handleBlogCreation(w http.ResponseWriter, r *http.Request) {
	// Handle Login
}
