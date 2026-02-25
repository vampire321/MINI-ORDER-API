package service

import(
	"mini-order-api/internal/model"
	"mini-order-api/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo : repo}
}

func (s *OrderS)