package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tongthanhdat009/gorm_api_project/models"
	"github.com/tongthanhdat009/gorm_api_project/services"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// ListTasks (GET /tasks) - Cập nhật để hỗ trợ phân trang
func (h *TaskHandler) ListTasks(c *gin.Context) {
	// Lấy query params 'page' và 'limit'
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "3")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 10 {
		limit = 3
	}

	tasks, err := h.service.ListTasks(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GetTask (GET /tasks/:id) - Sửa Atoi thành ParseUint
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32) // Sửa: dùng ParseUint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	task, err := h.service.GetTask(uint(id)) // Sửa: ép kiểu uint
	if err != nil {
		// Lỗi "record not found" của GORM nên trả về 404
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// CreateTask (POST /tasks) - Sửa logic gọi service
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Service GORM mới nhận con trỏ và cập nhật ID
	if err := h.service.CreateTask(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Trả về 'input' đã được GORM cập nhật ID
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// UpdateTask (PUT /tasks/:id) - Sửa Atoi thành ParseUint
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32) // Sửa: dùng ParseUint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	task, err := h.service.UpdateTask(uint(id), input) // Sửa: ép kiểu uint
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "task not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DeleteTask (DELETE /tasks/:id) - Sửa Atoi thành ParseUint
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32) // Sửa: dùng ParseUint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteTask(uint(id)); err != nil { // Sửa: ép kiểu uint
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.Status(http.StatusNoContent)
}