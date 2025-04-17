package main

import (
	"fmt"
	"go_ecom_api/cmd/api"
	"log"
)

func main() {
	server := api.NewServerAPI(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("server is running on port :8080")
}
