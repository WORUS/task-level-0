package repository

import (
	"task-level-0/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Order interface {
	GetOrder(id string) (model.Order, error)
	AddOrder(id string, order []byte) (string, error)
}

type Repository struct {
	Order
}

func NewRepository(pgx *pgxpool.Pool) *Repository {
	return &Repository{
		Order: NewOrderPostgres(pgx),
	}
}
