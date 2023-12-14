package service

import (
	"task-level-0/internal/repository"

	"github.com/sirupsen/logrus"
)

type OrderService struct {
	repository *repository.Repository
}

func NewOrderService(repo *repository.Repository) *OrderService {
	return &OrderService{
		repository: repo,
	}
}

func (s *OrderService) GetOrderById(id string) ([]byte, error) {
	order, hit := s.repository.Cache.GetOrderById(id)
	if hit {
		logrus.Info("cache hit: order was retreived from cache")
		return order, nil
	}

	order, err := s.repository.Postgres.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	key, err := s.repository.Cache.AddOrder(id, order)
	if err != nil {
		logrus.WithError(err).Infof("adding to cache order with ID = %s was unsuccessful", key)
	}
	logrus.Infof("order with ID = %s was added in cache", key)

	return order, nil
}

func (s *OrderService) AddOrder(id string, order []byte) (string, error) {
	id, err := s.repository.Cache.AddOrder(id, order)
	if err != nil {
		return id, err
	}

	return s.repository.Postgres.AddOrder(id, order)
}
