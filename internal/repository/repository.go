package repository

import (
	"database/sql"
	"errors"
	"url-shortener/internal/entity"
	"url-shortener/internal/repository/memory"
	"url-shortener/internal/repository/postgres"
)

type Repository interface {
	SaveURL(data entity.URL) (string, error)
	GetURL(data entity.URL) (string, error)
}

func NewRepository(useDB bool, db *sql.DB) (Repository, error) {
	if useDB {
		if db == nil {
			return nil, errors.New("postgres is not connected")
		}
		return postgres.NewPostgresRepository(db), nil
	}
	return memory.NewMemoryRepository(), nil
}
