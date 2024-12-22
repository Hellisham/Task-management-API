package main

import (
	"github.com/Hellisham/task_manager/internal/database"
	task "github.com/Hellisham/task_manager/internal/models"
)

func main() {
	database := db.Connect()
	err := database.AutoMigrate(&task.Task{})
	if err != nil {
		panic(err)
	}
}
