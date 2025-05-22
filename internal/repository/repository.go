package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository инициализируем репозиторий
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}
