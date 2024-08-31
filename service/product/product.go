package product

import (
	"database/sql"
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models"
	"github.com/KhanbalaRashidov/Go-Ecommerce/models/dto"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductByID(productID int) (*models.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", productID)
	if err != nil {
		return nil, err
	}

	p := new(models.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) GetProductsByID(productIDs []int) ([]models.Product, error) {
	placeholders := strings.Repeat(",?", len(productIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// Convert productIDs to []interface{}
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []models.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil

}

func (s *Store) GetProducts() ([]*models.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*models.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product dto.CreateProductDto) error {
	_, err := s.db.Exec("INSERT INTO products (name, price, image, description, quantity) VALUES (?, ?, ?, ?, ?)", product.Name, product.Price, product.Image, product.Description, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(product models.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, price = ?, image = ?, description = ?, quantity = ? WHERE id = ?", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*models.Product, error) {
	product := new(models.Product)

	err := rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
