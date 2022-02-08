package main

import (
	"log"

	"github.com/spf13/viper"

	"myToDoApp/internal/handler"
	"myToDoApp/internal/repository"
	"myToDoApp/internal/service"
	"myToDoApp/internal/service/server"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs : %s", err.Error())
	}

	db, err := repository.PostgresNewDB(repository.Config{
		Host:     "192.168.1.138",
		Port:     "1234",
		UserName: "rustam",
		Password: "1234",
		DBName:   "my_world",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
