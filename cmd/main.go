package main

import (
	"github.com/chigatu/service/pkg/handler"
	"github.com/chigatu/service/pkg/repository"
	"github.com/chigatu/service/pkg/service"
	"github.com/chigatu/service/server"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:    viper.GetString("db.host"),
		Port:    viper.GetString("db.port"),
		DBName:  viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("falied to initialize db: %v", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString(`port`), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %v", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
