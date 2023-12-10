package repository

import (
	"task-level-0/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

type Order interface {
	GetOrder(id int) model.Order
	AddOrder(order model.Order) int
}

type Repository struct {
	Order
}

func NewRepository(pgx *pgx.Conn) *Repository {
	return &Repository{
		Order: NewOrderPostgres(pgx),
	}
}
