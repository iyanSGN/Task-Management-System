package models

import (
	"time"

	"main.go/app/task"
)

type MTask struct {
	ID          uint   		`json:"id" gorm:"primary_key"`
	Title       string 		`json:"title"`
	Description string 		`json:"description"`
	Due         time.Time	`json:"due"`
	Status		string		`json:"status"`
}

func(mk *MTask) ToResponse() task.TaskResponseDTO{
	return task.TaskResponseDTO{
		ID: mk.ID,
		Title: mk.Title,
	}
}