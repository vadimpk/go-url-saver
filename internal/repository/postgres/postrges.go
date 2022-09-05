package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-urlsaver/internal/config"
)

const (
	usersTable = "users"
	urlsTable  = "urls"
)

func NewPostgresDB(cfg config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
