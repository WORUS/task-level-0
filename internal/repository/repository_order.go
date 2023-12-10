package repository

import (
	"task-level-0/internal/domain/model"

	"github.com/jackc/pgx/v5"
)

type OrderPostgres struct {
	pgx *pgx.Conn
}

func NewOrderPostgres(pgx *pgx.Conn) *OrderPostgres {
	return &OrderPostgres{pgx: pgx}
}

func (p *OrderPostgres) GetOrder(id int) model.Order {
	var order model.Order
	return order
}

func (p *OrderPostgres) AddOrder(order model.Order) int {
	return 1
}
