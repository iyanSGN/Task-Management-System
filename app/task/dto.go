package task

import (
	"time"
)

type TaskRequestDTO struct {
	ID          	uint   			`json:"id"`
	Title       	string 			`json:"title" validate:"required"`
	Description 	string 			`json:"description"`
	Due         	time.Time		`json:"due"`
}

type TaskResponseDTO struct {
	ID 			uint		`json:"id"`
	Title		string		`json:"title"`
}