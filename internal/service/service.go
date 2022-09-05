package service

import (
	"go-urlsaver"
	"go-urlsaver/internal/repository"
)

type Authorization interface {
	CreateUser(user go_url_saver.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type URL interface {
	CreateURL(userID int, url go_url_saver.Url) (int, error)
	GetAll(userID int) ([]go_url_saver.UrlResponse, error)
	GetByID(userID, urlID int) (go_url_saver.UrlResponse, error)
	DeleteURL(userID, urlID int) error
	UpdateURL(userID, urlID int, input go_url_saver.UpdateUrl) error
}

type Service struct {
	Authorization
	URL
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		URL:           NewUrlService(repo.URL),
	}
}
