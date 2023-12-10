package service

import (
	"task-level-0/internal/domain/model"
	"task-level-0/internal/repository"
)

type Order interface {
	GetOrder(id int) model.Order
	AddOrder(order model.Order) int
}

type Service struct {
	Order
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repo),
	}
}
