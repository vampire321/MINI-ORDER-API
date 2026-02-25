package service

import(
	"mini-order-api/internal/model"
	"mini-order-api/internal/repository"
	"errors"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo : repo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	var ErrInvalidAmount = errors.New("order amount must be greater than zero")
	if order.Amount <= 0 {
		return ErrInvalidAmount
	}
	return s.repo.Create(order)
}

func (s *OrderService) GetOrders() ([]*model.Order,error){
	return s.repo.GetALL()
}