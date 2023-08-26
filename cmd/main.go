package main

import (
	"log"

	"github.com/Cgigatu/service"
	"github.com/Cgigatu/service/pkg/handler"
)

func main() {
	handlers := handler.InitRoutes()
	srv := new(service.Server)
	if err := srv.Run("8080", handlers); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
