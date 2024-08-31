package repo

import "github.com/KhanbalaRashidov/Go-Ecommerce/models"

type OrderStore interface {
	CreateOrder(models.Order) (int, error)
	CreateOrderItem(models.OrderItem) error
}
