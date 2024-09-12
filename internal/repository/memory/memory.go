package memory

import (
	"errors"
	"os"
	"strconv"
	"url-shortener/internal/entity"
	"url-shortener/pkg/utils"
)

type MemoryRepository struct {
	storage map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{storage: make(map[string]string)}
}

func (r *MemoryRepository) SaveURL(data entity.URL) (string, error) {
	length, err := strconv.Atoi(os.Getenv("URL_LENGTH"))
	if err != nil {
		return "", errors.New("invalid shorten url length value")
	}
	alias := utils.GetRandomString(length)
	r.storage[alias] = data.Url
	return alias, nil
}

func (r *MemoryRepository) GetURL(data entity.URL) (string, error) {
	url, exists := r.storage[data.Alias]
	if !exists {
		return "", errors.New("url not found")
	}
	return url, nil
}
