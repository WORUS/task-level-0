package service

import (
	"task-level-0/internal/repository"
)

type Order interface {
	GetOrderById(id string) ([]byte, error)
	AddOrder(id string, order []byte) (string, error)
}

type Service struct {
	Order
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repo),
	}
}
