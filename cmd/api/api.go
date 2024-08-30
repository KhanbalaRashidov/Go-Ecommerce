package api

import "net/http"

type ApiServer struct {
	address string
}

func NewApiServer(address string) *ApiServer {
	return &ApiServer{address: address}
}

func (s *ApiServer) Run() error {

	return http.ListenAndServe(s.address, nil)
}
