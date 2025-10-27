package main

import (
	"accounts-service/internal/database"
	"accounts-service/internal/handlers"
	"accounts-service/internal/repository"
	"accounts-service/internal/routes"
	"accounts-service/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	repo := repository.NewAccountRepository(db)
	service := services.NewAccountService(repo)
	handler := handlers.NewAccountHandler(service)

	router := gin.Default()
	routes.RegisterAccountRoutes(router, handler)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
