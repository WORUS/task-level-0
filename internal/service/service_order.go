package service

import (
	"task-level-0/internal/domain/model"
	"task-level-0/internal/repository"
)

type OrderService struct {
	repository *repository.Repository
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{
		repository: repo,
	}
}

func (s *OrderService) GetOrder(id int) model.Order {
	return s.repository.GetOrder(id)
}

func (s *OrderService) AddOrder(order []byte) (string, error) {
	return s.repository.AddOrder(order)

}
