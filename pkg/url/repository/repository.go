package repository

import "database/sql"

type Url interface {
	Create(short, full string) (string, error)
	GetFull(short string) (full string, err error)
}

type UrlRepository struct {
	Url
}

func NewRepositoryUrl(hm_full, hm_short map[string]string) *UrlRepository {
	return &UrlRepository{
		Url: NewInMemoryUrlRepository(hm_full, hm_short),
	}
}

func NewRepositoryUrlDB(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		Url: NewPostgresUrlRepository(db),
	}
}
