package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type OrderPostgres struct {
	pgx *pgxpool.Pool
}

func NewOrderPostgres(pgx *pgxpool.Pool) *OrderPostgres {
	return &OrderPostgres{pgx: pgx}
}

func (p *OrderPostgres) GetOrderById(id string) ([]byte, error) {
	var order []byte

	err := p.pgx.QueryRow(context.Background(), "SELECT order_json FROM orders WHERE order_uid=$1", id).Scan(&order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (p *OrderPostgres) AddOrder(id string, content []byte) (string, error) {
	query := `INSERT INTO orders (order_uid, order_json) VALUES (@order_uid, @order_json)`
	args := pgx.NamedArgs{
		"order_uid":  id,
		"order_json": content,
	}

	_, err := p.pgx.Exec(context.Background(), query, args)
	if err != nil {
		logrus.WithError(err).Info("repository.postgres.AddOrder: order already exists")
	}

	return id, nil
}

func (p *OrderPostgres) GetOrders(limit int) (map[string][]byte, error) {
	orders := make(map[string][]byte)
	var key string
	var value []byte

	rows, err := p.pgx.Query(context.Background(), "SELECT order_uid, order_json FROM orders LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}
		orders[key] = value
	}
	return orders, nil
}
