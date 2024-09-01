package handler

import (
	"github.com/KhanbalaRashidov/Go-Ecommerce/repo"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/auth"
	"github.com/gorilla/mux"
	"net/http"
)

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

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.LoginHandle).Methods("POST")
	router.HandleFunc("/register", h.RegisterHandle).Methods("POST")

	// admin routes
	router.HandleFunc("/users/{userID}", auth.WithJWTAuth(h.GetUserHandle, h.userStore)).Methods(http.MethodGet)

	//cart
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.CheckoutHandle, h.userStore)).Methods(http.MethodPost)

	//product
	router.HandleFunc("/products", h.GetProductsHandle).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.GetProductHandle).Methods(http.MethodGet)

	// admin routes
	router.HandleFunc("/products", auth.WithJWTAuth(h.CreateProductHandle, h.userStore)).Methods(http.MethodPost)
}
