package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gabin-ishimwe/complex-crud-go/service/blog"
	"github.com/gabin-ishimwe/complex-crud-go/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	mainRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	// dependy injection
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(mainRouter)

	blogStore := blog.NewBlogRepository(s.db)
	blogHandler := blog.NewBlogHandler(blogStore)
	blogHandler.RegisterBlogHandler(mainRouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
