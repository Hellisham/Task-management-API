package handler

import (
	"encoding/json"
	"github.com/Hellisham/task_manager/internal/services"
	"net/http"
	"time"
)

type TaskHandler struct {
	Service *services.TaskService
}

type TaskResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created_At  time.Time `json:"created_at"`
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{
		Service: taskService,
	}
}

func (handler *TaskHandler) GetAllTasksHanlder(w http.ResponseWriter, r *http.Request) {
	task, err := handler.Service.GetAllTasksService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var taskresponse = []TaskResponse{}
	for _, v := range task {
		taskresponse = append(taskresponse, TaskResponse{
			Title:       v.Title,
			Description: v.Discretion,
			Created_At:  v.CreatedAt,
		})
		json.NewEncoder(w).Encode(taskresponse)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
