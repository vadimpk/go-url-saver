package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-urlsaver"
)

type URLPostgres struct {
	db *sqlx.DB
}

func NewURLPostgres(db *sqlx.DB) *URLPostgres {
	return &URLPostgres{db: db}
}
func (r *URLPostgres) Create(userID int, url go_url_saver.Url) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (url, description, user_id) VALUES($1, $2, $3) RETURNING id", urlsTable)
	row := r.db.QueryRowx(query, url.Url, url.Description, userID)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *URLPostgres) GetAll(userID int) ([]go_url_saver.UrlResponse, error) {
	var urls []go_url_saver.UrlResponse
	query := fmt.Sprintf("SELECT url, description FROM %s WHERE user_id = $1", urlsTable)

	err := r.db.Select(&urls, query, userID)

	return urls, err
}

func (r *URLPostgres) GetByID(userID, urlID int) (go_url_saver.UrlResponse, error) {
	var url go_url_saver.UrlResponse
	query := fmt.Sprintf("SELECT url, description FROM %s WHERE user_id = $1 AND id = $2", urlsTable)

	err := r.db.Get(&url, query, userID, urlID)

	return url, err
}

func (r *URLPostgres) Delete(userID, urlID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", urlsTable)

	_, err := r.db.Exec(query, userID, urlID)

	return err
}

func (r *URLPostgres) Update(userID, urlID int, input go_url_saver.UpdateUrl) error {
	query := fmt.Sprintf("UPDATE %s SET url = $1, description = $2 WHERE user_id = $3 AND id = $4", urlsTable)

	_, err := r.db.Exec(query, input.Url, input.Description, userID, urlID)

	return err
}
