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

func (s *OrderService) GetOrder(id string) ([]byte, error) {
	order, hit := s.repository.Cache.GetOrder(id)
	if hit {
		logrus.Info("Cache hit: order was retreived from cache")
		return order, nil
	}

	order, err := s.repository.Postgres.GetOrder(id)
	if err != nil {
		return nil, err
	}

	key, err := s.repository.Cache.AddOrder(id, order)
	if err != nil {
		logrus.WithError(err).Infof("Adding to cache order with ID = %s was unsuccessful", key)
	}
	logrus.Infof("Order with ID = %s was added in cache", key)

	return order, nil
}

func (s *OrderService) AddOrder(id string, order []byte) (string, error) {
	id, err := s.repository.Cache.AddOrder(id, order)
	if err != nil {
		logrus.Info("Error occurred while add order into cache")
	}

	return s.repository.Postgres.AddOrder(id, order)
}
