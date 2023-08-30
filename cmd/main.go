package main

import (
	"log"

	"github.com/chigatu/service/pkg/handler"
	"github.com/chigatu/service/pkg/repository"
	"github.com/chigatu/service/pkg/service"
	"github.com/chigatu/service/server"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %v", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString(`port`), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
