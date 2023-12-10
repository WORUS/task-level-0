package repository

import (
	"context"
	"task-level-0/internal/domain/model"

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

func (p *OrderPostgres) GetOrder(id string) (model.Order, error) {
	var order model.Order
	err := p.pgx.QueryRow(context.Background(), "SELECT order_json FROM orders WHERE order_uid=$1", id).Scan(&order)
	if err != nil {
		return order, err
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
		logrus.Fatal(err)
	}
	// sd, err := p.pgx.Prepare(context.Background(), "insert", "INSERT INTO orders(order_uid, order) values(?, ?)")
	// if err != nil {
	// 	logrus.Fatal("Error occurred adding order into db")
	// }
	// p.pgx.Exec(ctx, query, args)

	return id, nil
}
