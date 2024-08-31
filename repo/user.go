package repo

import "github.com/KhanbalaRashidov/Go-Ecommerce/models"

type UserStore interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(models.User) error
}
