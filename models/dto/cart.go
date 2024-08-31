package dto

import "github.com/KhanbalaRashidov/Go-Ecommerce/models"

type CartCheckoutDto struct {
	Items []models.CartCheckoutItem `json:"items" validate:"required"`
}
