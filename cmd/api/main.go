package main

import (
	"example/Books_Go_Api/internal/database"
	"example/Books_Go_Api/internal/employee"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.CreateEmployeeTable(db); err != nil {
		log.Fatal(err)
	}

	repo := employee.NewPostgresRepository(db)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)

	router := gin.Default()

	router.GET("/employees", handler.GetAll)
	router.GET("/employees/:id", handler.GetById)
	router.POST("/employees", handler.RegisterEmployee)

	router.PATCH("/employees/:id/fire", handler.Fire)
	router.PATCH("/employees/:id/employ", handler.Employ)

	router.Run(":8080")
}