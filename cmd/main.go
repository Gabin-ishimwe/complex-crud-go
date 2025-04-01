package main

import (
	"log"

	"github.com/gabin-ishimwe/complex-crud-go/cmd/api"
	"github.com/gabin-ishimwe/complex-crud-go/db"
)

func main() {
	// Connect to Postgres
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	// Start Server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
