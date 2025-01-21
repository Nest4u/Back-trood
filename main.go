package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	db "github.com/troodinc/trood-front-hackathon/database"
	_ "github.com/troodinc/trood-front-hackathon/docs"
	"github.com/troodinc/trood-front-hackathon/handlers"
)

// @title Trood Front Hackathon API
// @version 1.0
// @description This is the API documentation for the Trood Front Hackathon. Welcome to hell.
// @host localhost:8080
// @BasePath /

func main() {
	// Инициализация базы данных и обработчиков
	db.InitDatabase()
	handlers.InitProjects()

	// Создание нового роутера
	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://front-trood.vercel.app"}, // Укажите домен фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршруты API
	r.GET("/projects", handlers.GetProjects)
	r.GET("/projects/:id", handlers.GetProjectByID)
	r.POST("/projects", handlers.CreateProject)
	r.PUT("/projects/:id", handlers.EditProject)
	r.DELETE("/projects/:id", handlers.DeleteProject)

	r.GET("/projects/:id/vacancies", handlers.GetVacancies)
	r.POST("/projects/:id/vacancies", handlers.CreateVacancy)
	r.PUT("/vacancies/:id", handlers.EditVacancy)
	r.DELETE("/vacancies/:id", handlers.DeleteVacancy)

	// Запуск сервера
	port := "8080"
	log.Println("Server running on http://localhost:" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
