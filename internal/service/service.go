package service

import "go-urlsaver/internal/repository"

type Authorization interface {
}

type URLList interface {
}

type URL interface {
}

type Service struct {
	Authorization
	URLList
	URL
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
