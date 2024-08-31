package handler

import "github.com/KhanbalaRashidov/Go-Ecommerce/repo"

type Handler struct {
	productStore repo.ProductStore
	orderStore   repo.OrderStore
	userStore    repo.UserStore
}

func NewHandler(productStore repo.ProductStore, orderStore repo.OrderStore, userStore repo.UserStore) *Handler {
	return &Handler{
		productStore: productStore,
		orderStore:   orderStore,
		userStore:    userStore,
	}
}
