package api

import (
	"database/sql"
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

	return http.ListenAndServe(s.address, nil)
}
