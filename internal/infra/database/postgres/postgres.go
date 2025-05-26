package postgresDatabase

import (
	"container-manager/internal/infra/config"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDataBase struct{}

func (p *PostgresDataBase) Connect() (*sql.DB, error) {
	setting := config.GetConfig().Pg
	userName := setting.User
	password := setting.Password
	host := setting.Host
	databaseName := setting.Name
	connectString := "postgres://" + userName + ":" + password +
		"@" + host + "?" + "database=" + databaseName + "&sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Default().Println("Postgres connected")

	return db, err
}

func (p *PostgresDataBase) Disconnect(db *sql.DB) error {
	if err := db.Close(); err != nil {
		log.Fatal(err)
		return err
	}
	log.Default().Println("Postgres disconnected")
	return nil
}

func NewPostgresDatabase() *PostgresDataBase {
	return &PostgresDataBase{}
}
