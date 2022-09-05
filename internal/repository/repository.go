package repository

import (
	"github.com/jmoiron/sqlx"
	"go-urlsaver"
	"go-urlsaver/internal/repository/postgres"
)

type Authorization interface {
	CreateUser(user go_url_saver.User) (int, error)
	GetUser(username, password string) (go_url_saver.User, error)
}

type URL interface {
	Create(userID int, url go_url_saver.Url) (int, error)
	GetAll(userID int) ([]go_url_saver.UrlResponse, error)
	GetByID(userID, urlID int) (go_url_saver.UrlResponse, error)
	Delete(userID, urlID int) error
	Update(userID, urlID int, input go_url_saver.UpdateUrl) error
}

type Repository struct {
	Authorization
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		URL:           postgres.NewURLPostgres(db),
	}
}
