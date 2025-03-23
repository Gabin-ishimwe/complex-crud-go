package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gabin-ishimwe/complex-crud-go/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Envs.PgHost, 5000, config.Envs.PgUser, config.Envs.PgPassword, config.Envs.PgDbName)
	db, error := sql.Open(config.Envs.PgDriverName, dataSourceName)
	if error != nil {
		log.Fatal("Error connecting to the database: ", error)
	}
	defer db.Close()

	// Check if database connection was successful
	if error := db.Ping(); error != nil {
		log.Fatal(error)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db, nil
}
