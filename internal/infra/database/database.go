package database

import (
	postgresDatabase "container-manager/internal/infra/database/postgres"
	"database/sql"
)

type Database interface {
	Connect() (*sql.DB, error)
	Disconnect(*sql.DB) error
}

func NewDatabase(dbType string) Database {
	switch dbType {
	case "postgres":
		return postgresDatabase.NewPostgresDatabase()
	default:
		return postgresDatabase.NewPostgresDatabase()
	}
}
