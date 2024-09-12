package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"url-shortener/config"
	"url-shortener/internal/entity"
	"url-shortener/pkg/utils"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("error connecting to database")
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS urls(
	    alias TEXT NOT NULL PRIMARY KEY,
	    url TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating table %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *PostgresRepository) SaveURL(data entity.URL) (string, error) {
	length, err := strconv.Atoi(os.Getenv("URL_LENGTH"))
	if err != nil {
		return "", errors.New("invalid shorten url length value")
	}
	alias := utils.GetRandomString(length)
	_, err = r.db.Exec(`INSERT INTO urls(url, alias) VALUES($1, $2)`, data.Url, alias)
	if err != nil {
		return "", errors.New("error saving url")
	}
	return alias, nil
}

func (r *PostgresRepository) GetURL(data entity.URL) (string, error) {
	var resURL string
	err := r.db.QueryRow(`SELECT url FROM urls WHERE alias = $1`, data.Alias).Scan(&resURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("url not found")
		}
		return "", errors.New("error getting url")
	}

	return resURL, nil
}
