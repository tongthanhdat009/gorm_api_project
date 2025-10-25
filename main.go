package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/gorm_api_project/database" // THÊM: Import database
	"github.com/tongthanhdat009/gorm_api_project/handlers"
	"github.com/tongthanhdat009/gorm_api_project/middlewares"
	// SỬA LẠI: Dùng 'repository' (mới)
	"github.com/tongthanhdat009/gorm_api_project/repository" 
	"github.com/tongthanhdat009/gorm_api_project/services"
)

func main() {
	// 1. THÊM: Kết nối CSDL GORM
	database.ConnectDatabase()

	// 2. SỬA: Khởi tạo GORM repository
	// Xóa: repo := repositories.NewInMemoryTaskRepository()
	repo := repository.NewTaskRepository(database.DB)

	// 3. Giữ nguyên service và handler
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	router := gin.Default()

	// Giữ nguyên middleware
	router.Use(middlewares.LoggingMiddleware())

	api := router.Group("/api")
	api.GET("/tasks", handler.ListTasks) // Route này giờ đã hỗ trợ phân trang
	api.POST("/tasks", middlewares.SimpleAuthMiddleware(), handler.CreateTask)
	api.GET("/tasks/:id", handler.GetTask)
	api.PUT("/tasks/:id", middlewares.SimpleAuthMiddleware(), handler.UpdateTask)
	api.DELETE("/tasks/:id", middlewares.SimpleAuthMiddleware(), handler.DeleteTask)

	if err := router.Run(":8080"); err != nil { // Thêm port :8080
		log.Fatalf("failed to run server: %v", err)
	}
}