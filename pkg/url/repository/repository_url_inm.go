package repository

import (
	"fmt"
	"sync"
)

type InMemoryUrlRepository struct {
	full_urls  map[string]string
	short_urls map[string]string
	mu         sync.RWMutex
}

func NewInMemoryUrlRepository(hm_full, hm_short map[string]string) *InMemoryUrlRepository {
	return &InMemoryUrlRepository{
		full_urls:  hm_full,
		short_urls: hm_short,
	}
}

func (r *InMemoryUrlRepository) Create(short, full string) (string, error) {

	if value, ok := r.full_urls[full]; !ok {
		r.mu.RLock()
		defer r.mu.RUnlock()
		r.full_urls[full] = short
		r.short_urls[short] = full
		return short, nil
	} else {
		return value, nil
	}
}

func (r *InMemoryUrlRepository) GetFull(short string) (full string, err error) {

	if value, ok := r.short_urls[short]; !ok {
		return "3", fmt.Errorf("No short URL.")
	} else {
		return value, nil
	}
}
