package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type URLList interface {
}

type URL interface {
}

type Repository struct {
	Authorization
	URLList
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
