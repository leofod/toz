package repository

import (
	"database/sql"
)

type PostgresUrlRepository struct {
	db *sql.DB
}

func NewPostgresUrlRepository(db *sql.DB) *PostgresUrlRepository {
	return &PostgresUrlRepository{db: db}
}

func (r *PostgresUrlRepository) Create(short, full string) (string, error) {
	var short_url string
	row := r.db.QueryRow("SELECT short_url FROM urls WHERE full_url = $1", full)

	if err := row.Scan(&short_url); err != nil {
		q, err := r.db.Prepare("INSERT INTO urls (short_url, full_url) VALUES ($1, $2)")

		if err != nil {
			return "", err
		}

		_, err = q.Exec(short, full)

		if err != nil {
			return "", err
		}
		return short, nil
	}

	return short_url, nil
}

func (r *PostgresUrlRepository) GetFull(short string) (full string, err error) {
	row := r.db.QueryRow("SELECT full_url FROM urls WHERE short_url = $1", short)

	if err := row.Scan(&full); err != nil {
		return "", err
	}

	return
}
