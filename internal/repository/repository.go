package repository

import (
	"task-level-0/internal/repository/cache"
	"task-level-0/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres interface {
	GetOrders(limit int) (map[string][]byte, error)
	GetOrderById(id string) ([]byte, error)
	AddOrder(id string, order []byte) (string, error)
}

type Cache interface {
	GetOrderById(id string) ([]byte, bool)
	AddOrder(id string, order []byte) (string, error)
	DeleteOrder(id string) error
	GetCapacity() int
}

type Repository struct {
	Cache
	Postgres
}

func NewRepository(pgx *pgxpool.Pool, cacheCapacity int, ordersMap map[string][]byte) *Repository {
	return &Repository{
		Cache:    cache.NewOrderCache(cacheCapacity),
		Postgres: postgres.NewOrderPostgres(pgx),
	}
}

func (r *Repository) RestoreCache() {
	value, err := r.Postgres.GetOrders(r.Cache.GetCapacity())
	if err != nil {
		logrus.WithError(err).Fatal("error occurred restore cache: get data from db failed")
		return
	}
	for k := range value {
		_, err := r.Cache.AddOrder(k, value[k])
		if err != nil {
			logrus.WithError(err).Fatal("error occurred restore cache: add db data into cache failed")
		}
	}
}
