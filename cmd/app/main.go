package main

import (
	"log"

	"myToDoApp/internal/handler"
	"myToDoApp/service/server"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
