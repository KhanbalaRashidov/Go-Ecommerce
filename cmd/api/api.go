package api

import (
	"database/sql"
	"github.com/KhanbalaRashidov/Go-Ecommerce/handler"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/order"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/product"
	"github.com/KhanbalaRashidov/Go-Ecommerce/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

type ApiServer struct {
	address string
	db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *ApiServer {
	return &ApiServer{address: address, db: db}
}

func (s *ApiServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	productStore := product.NewStore(s.db)
	orderStore := order.NewStore(s.db)

	handlerRoute := handler.NewHandler(productStore, orderStore, userStore)
	handlerRoute.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return http.ListenAndServe(s.address, router)
}
