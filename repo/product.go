package repo

import (
	"github.com/KhanbalaRashidov/Go-Ecommerce/models"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models/dto"
)

type ProductStore interface {
	GetProductByID(id int) (*models.Product, error)
	GetProductsByID(ids []int) ([]models.Product, error)
	GetProducts() ([]*models.Product, error)
	CreateProduct(dto dto.CreateProductDto) error
	UpdateProduct(models.Product) error
}
