package repository

import (
	"task-level-0/internal/repository/cache"
	"task-level-0/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres interface {
	GetOrder(id string) ([]byte, error)
	AddOrder(id string, order []byte) (string, error)
}

type Cache interface {
	GetOrder(id string) ([]byte, bool)
	AddOrder(id string, order []byte) (string, error)
	DeleteOrder(id string) error
}

type Repository struct {
	Cache
	Postgres
}

func NewRepository(pgx *pgxpool.Pool, cacheCapacity uint, ordersMap map[string][]byte) *Repository {
	return &Repository{
		Cache:    cache.NewOrderCache(cacheCapacity, ordersMap),
		Postgres: postgres.NewOrderPostgres(pgx),
	}
}
