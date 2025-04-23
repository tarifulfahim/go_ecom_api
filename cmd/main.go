package main

import (
	"go_ecom_api/cmd/api"
	"go_ecom_api/config"
	"log"
)

func main() {
	server := api.NewServerAPI(":8080", config.ENVS.DatabaseURL)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
