package main

import (
	"github.com/joho/godotenv"
	"log"
	"test/internal/database"
	"test/internal/handler"
	"test/internal/server"
	"test/internal/service"
	"test/pkg"
)

func main() {
	if err := pkg.InitConfig(); err != nil {
		log.Fatal()
	}
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
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
