package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zekhoi/golang-todo/app/handlers"
	"github.com/zekhoi/golang-todo/app/repositories"
	"github.com/zekhoi/golang-todo/app/routes"
	"github.com/zekhoi/golang-todo/app/services"
	"github.com/zekhoi/golang-todo/config"
)

func main() {

	// Initialize database
	config.InitDB()

	// Get database connection
	db := config.GetDB()

	r := gin.Default()
	r.Use(gin.Logger())

	// Initialize repositories
	activityRespository := repositories.NewActivityRepository(db)
	todoRespository := repositories.NewToDoRepository(db)

	// Initialize services
	activityService := services.NewActivityService(activityRespository)
	todoService := services.NewToDoService(todoRespository)

	// Initialize handlers
	activityHandler := handlers.NewActivityHandler(activityService)
	todoHandler := handlers.NewToDoHandler(todoService)

	// Register routes
	routes.RegisterActivityRoute(r, activityHandler)
	routes.RegisterToDoRoute(r, todoHandler)

	// Start server
	log.Println("Server running on port localhost:3030")
	err := r.Run(":3030")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
