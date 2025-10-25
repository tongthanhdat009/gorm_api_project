package database

import (
	"errors" // Thêm import 'errors'
	"fmt"
	"log"

	"github.com/tongthanhdat009/gorm_api_project/models" // Đã có
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:12345@tcp(127.0.0.1:3306)/chuong10?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connection successfully established")

	// Tự động migrate schema
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migrated")

	// GỌI HÀM SEED DỮ LIỆU
	seedDatabase()
}

// seedDatabase thêm dữ liệu mẫu nếu bảng 'tasks' trống
func seedDatabase() {
	// 1. Kiểm tra xem đã có task nào chưa
	var task models.Task
	err := DB.First(&task).Error

	// 2. Nếu không tìm thấy (lỗi là 'record not found'), thì chúng ta mới seed
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("No tasks found, seeding database...")

		// Tạo một vài task mẫu
		tasks := []models.Task{
			{
				Title:       "Hoàn thành Đồ án nhỏ 2",
				Description: "Refactor API dùng GORM và MySQL",
				Status:      "in_progress",
			},
			{
				Title:       "Học về Pagination",
				Description: "Implement Limit() và Offset()",
				Status:      "pending",
			},
			{
				Title:       "Test API với Postman",
				Description: "Kiểm tra GET, POST, PUT, DELETE",
				Status:      "pending",
			},
			{
				Title:       "Uống cà phê sữa",
				Description: "Task đã hoàn thành",
				Status:      "completed",
			},
			{
				Title:       "Uống cà phê muối",
				Description: "Task đã hoàn thành",
				Status:      "completed",
			},
			{
				Title:       "Uống cà phê ê",
				Description: "Task đã hoàn thành",
				Status:      "completed",
			},
		}

		// 3. Thêm các task này vào CSDL
		if err := DB.Create(&tasks).Error; err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}

		fmt.Printf("Seeded %d tasks successfully.\n", len(tasks))
	} else if err != nil {
		// Lỗi khác khi truy vấn
		log.Fatalf("Failed to check for tasks: %v", err)
	} else {
		// Nếu không lỗi và err == nil, nghĩa là đã có data
		fmt.Println("Database already seeded.")
	}
}