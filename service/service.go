package service

import (
	"fmt"
	"example/solid/model"
)

// OrderService — high-level модуль, зависит только от абстракций model.
type OrderService struct {
	repo     model.RepositoryWriter
	notifier model.Notifier
}

func NewOrderService(repo model.RepositoryWriter, notifier model.Notifier) *OrderService {
	return &OrderService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	err := s.repo.CreateOrder(customer, products, total)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(model.OrderCreatedMessage, total)
	return s.notifier.Notify(customer, message)
}
