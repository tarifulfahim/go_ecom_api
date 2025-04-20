package main

import (
	"go_ecom_api/cmd/api"
	"log"
)

func main() {
	server := api.NewServerAPI(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
