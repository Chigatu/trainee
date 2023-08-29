package main

import (
	"log"

	"github.com/chigatu/service/pkg/handler"
	"github.com/chigatu/service/pkg/repository"
	"github.com/chigatu/service/pkg/service"
	"guthub.com/chigatu/service/server"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server", err.Error())
	}
}
