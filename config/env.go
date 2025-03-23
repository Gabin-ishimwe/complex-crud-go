package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	PgHost       string
	PgPort       string
	PgUser       string
	PgPassword   string
	PgDbName     string
	PgDriverName string
}

const (
	host       = "localhost"
	port       = "5432"
	user       = "root"
	password   = "admin"
	dbname     = "crud-db"
	driverName = "postgres"
)

var Envs = initConfig()

func initConfig() PostgresConfig {
	godotenv.Load()
	return PostgresConfig{
		PgHost:       getEnv("PG_HOST", host),
		PgPort:       getEnv("PG_PORT", port),
		PgUser:       getEnv("PG_USER", user),
		PgPassword:   getEnv("PG_PASSWORD", password),
		PgDbName:     getEnv("PG_DB_NAME", dbname),
		PgDriverName: getEnv("PG_DRIVER_NAME", driverName),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("value", value)
		return value
	}
	return fallback
}
