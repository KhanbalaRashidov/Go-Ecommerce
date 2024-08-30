package main

import (
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/cmd/api"
	"github.com/KhanbalaRashidov/Go-Ecommerce/configs"
	"log"
)

func main() {

	apiServer := api.NewApiServer(fmt.Sprintf(":%s", configs.Envs.Port))
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
