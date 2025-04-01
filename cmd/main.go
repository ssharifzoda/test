package main

import (
	"log"
	"test/internal/database"
	"test/internal/handler"
	"test/internal/server"
	"test/internal/service"
)

func main() {
	connection := database.NewConnectPostgres()
	db := database.NewDatabase(connection)
	service1 := service.NewService(db)
	handlers := handler.NewHandler(service1)
	server1 := new(server.Server)
	err := server1.ServerRun(handlers.Init(), "5005")
	if err != nil {
		log.Println(err)
	}
}
