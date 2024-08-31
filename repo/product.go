package repo

import "github.com/KhanbalaRashidov/Go-Ecommerce/models"

type ProductStore interface {
	GetProductByID(id int) (*models.Product, error)
	GetProductsByID(ids []int) ([]models.Product, error)
	GetProducts() ([]*models.Product, error)
	CreateProduct(models.Product) error
	UpdateProduct(models.Product) error
}
