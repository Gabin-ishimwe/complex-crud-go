package main

import (
	"log"

	"github.com/gabin-ishimwe/complex-crud-go/cmd/api"
	"github.com/gabin-ishimwe/complex-crud-go/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
