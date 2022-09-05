package service

import (
	"go-urlsaver"
	"go-urlsaver/internal/repository"
)

type UrlService struct {
	repo repository.URL
}

func NewUrlService(repo repository.URL) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) CreateURL(userID int, url go_url_saver.Url) (int, error) {
	return s.repo.Create(userID, url)
}

func (s *UrlService) GetAll(userID int) ([]go_url_saver.UrlResponse, error) {
	return s.repo.GetAll(userID)
}

func (s *UrlService) GetByID(userID, urlID int) (go_url_saver.UrlResponse, error) {
	return s.repo.GetByID(userID, urlID)
}

func (s *UrlService) DeleteURL(userID, urlID int) error {
	return s.repo.Delete(userID, urlID)
}

func (s *UrlService) UpdateURL(userID, urlID int, input go_url_saver.UpdateUrl) error {

	url, err := s.GetByID(userID, urlID)
	if err != nil {
		return err
	}

	if input.Url == nil {
		input.Url = &url.Url
	}

	if input.Description == nil {
		input.Description = &url.Description
	}

	return s.repo.Update(userID, urlID, input)
}
