package services

import (
	"errors"
	"strings"

	"github.com/tongthanhdat009/gorm_api_project/models"
	"github.com/tongthanhdat009/gorm_api_project/repository" 
)

type TaskService interface {
	ListTasks(page, limit int) ([]models.Task, error)
	GetTask(id uint) (models.Task, error)
	CreateTask(input *models.Task) error // Sửa: nhận con trỏ, trả về lỗi
	UpdateTask(id uint, input models.Task) (models.Task, error) // Sửa: ID là uint
	DeleteTask(id uint) error // Sửa: ID là uint
}

type taskService struct {
	// SỬA LẠI: Dùng interface mới
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

// Sửa: Thêm page, limit
func (s *taskService) ListTasks(page, limit int) ([]models.Task, error) {
	return s.repo.FindAll(page, limit)
}

// Sửa: ID là uint
func (s *taskService) GetTask(id uint) (models.Task, error) {
	return s.repo.FindByID(id)
}

// Sửa: Logic cho GORM
func (s *taskService) CreateTask(input *models.Task) error {
	if strings.TrimSpace(input.Title) == "" {
		return errors.New("title is required")
	}
	// Gán giá trị mặc định nếu cần
	if input.Status == "" {
		input.Status = "pending"
	}
	// Repo GORM sẽ cập nhật ID vào con trỏ 'input'
	return s.repo.Create(input)
}

// Sửa: Logic cho GORM và ID là uint
func (s *taskService) UpdateTask(id uint, input models.Task) (models.Task, error) {
	if strings.TrimSpace(input.Title) == "" {
		return models.Task{}, errors.New("title is required")
	}

	// 1. Tìm task hiện có
	task, err := s.repo.FindByID(id)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}

	// 2. Cập nhật các trường
	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status

	// 3. Lưu lại (dùng con trỏ)
	if err := s.repo.Update(&task); err != nil {
		return models.Task{}, err
	}
	
	return task, nil
}

// Sửa: ID là uint
func (s *taskService) DeleteTask(id uint) error {
	// Kiểm tra xem task có tồn tại không (tùy chọn)
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("task not found")
	}
	return s.repo.Delete(id)
}