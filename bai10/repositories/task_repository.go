package repository

import (
	// SỬA LẠI: Dùng module path từ go.mod
	"github.com/tongthanhdat009/gorm_api_project/models"
	"gorm.io/gorm"
)

// TaskRepository định nghĩa các hàm tương tác với CSDL
type TaskRepository interface {
	Create(task *models.Task) error
	FindAll(page, limit int) ([]models.Task, error)
	FindByID(id uint) (models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Create thêm một công việc mới vào CSDL
func (r *taskRepository) Create(task *models.Task) error {
	// GORM sẽ tự động cập nhật con trỏ 'task' với ID và timestamps
	return r.db.Create(task).Error
}

// FindAll lấy danh sách công việc với phân trang
func (r *taskRepository) FindAll(page, limit int) ([]models.Task, error) {
	var tasks []models.Task
	
	// Tính toán offset
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	err := r.db.Offset(offset).Limit(limit).Find(&tasks).Error
	return tasks, err
}

// FindByID tìm một công việc bằng ID
func (r *taskRepository) FindByID(id uint) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error // First tìm theo primary key
	return task, err
}

// Update cập nhật thông tin một công việc
func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete xóa một công việc (sẽ là soft delete vì dùng gorm.Model)
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}