package main

import (
	"fmt"
	"github.com/KhanbalaRashidov/Go-Ecommerce/cmd/api"
	"github.com/KhanbalaRashidov/Go-Ecommerce/configs"
	"github.com/KhanbalaRashidov/Go-Ecommerce/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySqlStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewApiServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
