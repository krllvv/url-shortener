package service

import (
	"fmt"
	"net/url"
	"url-shortener/internal/entity"
	"url-shortener/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveURL(data entity.URL) (string, error) {
	u, err := url.Parse(data.Url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return "", fmt.Errorf("incorrect url")
	}
	return s.repo.SaveURL(entity.URL{Url: fmt.Sprint(u)})
}

func (s *Service) GetURL(data entity.URL) (string, error) {
	return s.repo.GetURL(data)
}
