package services

import (
	task "github.com/Hellisham/task_manager/internal/models"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{
		DB: db,
	}
}

func (service *TaskService) GetAllTasksService() ([]task.Task, error) {
	var tasks []task.Task
	service.DB.Find(&tasks)
	if service.DB.Error != nil {
		return nil, service.DB.Error
	}
	return tasks, nil
}
