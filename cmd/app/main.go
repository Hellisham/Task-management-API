package main

import (
	"log"
)

func main() {
	data = db.Connect()
	err := data.Automigrate(&models.Task{})
	if err != nil {
		log.Fatal("Error migrate database: %v", err)
	}
}
